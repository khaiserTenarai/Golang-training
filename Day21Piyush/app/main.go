// Day 21: AWS + Go Integration — Mini Project (Task 15)
// Go Employee API with S3, Secrets Manager, CloudWatch, RDS via AWS SDK.
// Runs locally (in-memory, simulated AWS) or on AWS (real services).
//
// Run: cd app && go mod tidy && go run main.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"employee-api/awsutil"
	"employee-api/config"
	"employee-api/handler"
	"employee-api/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	var db *sql.DB
	if cfg.UseDB {
		var err error
		db, err = sql.Open("postgres", cfg.DatabaseURL)
		if err != nil {
			log.Fatal("DB error:", err)
		}
		db.Exec(`CREATE TABLE IF NOT EXISTS employees (
			id SERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL, age INT,
			department VARCHAR(50), salary DECIMAL(10,2),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`)
		fmt.Println("[DB] Connected to PostgreSQL")
		defer db.Close()
	} else {
		fmt.Println("[DB] In-memory mode (set DATABASE_URL for RDS)")
	}

	empHandler := handler.NewEmployeeHandler(db)
	r := mux.NewRouter()

	// Health
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		dbStatus := "in-memory"
		if db != nil {
			if db.Ping() == nil {
				dbStatus = "connected"
			} else {
				dbStatus = "disconnected"
			}
		}
		handler.Respond(w, 200, map[string]interface{}{
			"service": "employee-api", "status": "healthy",
			"database": dbStatus, "environment": cfg.Environment,
		})
	}).Methods("GET")

	// Employee CRUD
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/employees", empHandler.GetAll).Methods("GET")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.GetByID).Methods("GET")
	api.HandleFunc("/employees", empHandler.Create).Methods("POST")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.Delete).Methods("DELETE")

	// Task 5: S3 upload employee report
	api.HandleFunc("/aws/s3/upload", func(w http.ResponseWriter, r *http.Request) {
		key := fmt.Sprintf("reports/employees_%s.csv", time.Now().Format("20060102_150405"))
		csv := empHandler.GenerateCSV()
		s3Client, err := awsutil.NewS3Client(cfg.AWSRegion, cfg.S3Bucket)
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{
				Success: true, Message: "S3 simulated", Data: awsutil.SimulatedS3Response(cfg.S3Bucket, key, cfg.AWSRegion),
			})
			return
		}
		url, err := s3Client.Upload(key, []byte(csv), "text/csv")
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{Success: false, Message: err.Error(),
				Data: awsutil.SimulatedS3Response(cfg.S3Bucket, key, cfg.AWSRegion)})
			return
		}
		handler.Respond(w, 200, models.APIResponse{Success: true, Message: "Uploaded to S3", Data: map[string]string{"url": url, "key": key}})
	}).Methods("POST")

	// Task 5: S3 list objects
	api.HandleFunc("/aws/s3/list", func(w http.ResponseWriter, r *http.Request) {
		s3Client, err := awsutil.NewS3Client(cfg.AWSRegion, cfg.S3Bucket)
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{Success: true, Data: awsutil.SimulatedS3Response(cfg.S3Bucket, "reports/", cfg.AWSRegion)})
			return
		}
		objects, err := s3Client.ListObjects("reports/")
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{Success: false, Message: err.Error()})
			return
		}
		handler.Respond(w, 200, models.APIResponse{Success: true, Data: objects})
	}).Methods("GET")

	// Task 6: Test RDS connectivity
	api.HandleFunc("/aws/rds/test", func(w http.ResponseWriter, r *http.Request) {
		if cfg.DatabaseURL == "" {
			handler.Respond(w, 200, models.APIResponse{Success: true, Data: awsutil.SimulatedRDSResponse()})
			return
		}
		result := awsutil.TestRDSConnection(cfg.DatabaseURL)
		handler.Respond(w, 200, models.APIResponse{Success: result.Connected, Data: result})
	}).Methods("GET")

	// Task 8: Retrieve secret from Secrets Manager
	api.HandleFunc("/aws/secrets", func(w http.ResponseWriter, r *http.Request) {
		client, err := awsutil.NewSecretsClient(cfg.AWSRegion)
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{Success: true, Data: awsutil.SimulatedSecretResponse(cfg.SecretName)})
			return
		}
		secrets, err := client.GetSecret(cfg.SecretName)
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{Success: false, Message: err.Error(),
				Data: awsutil.SimulatedSecretResponse(cfg.SecretName)})
			return
		}
		handler.Respond(w, 200, models.APIResponse{Success: true, Message: "Secret retrieved", Data: map[string]interface{}{"keys": len(secrets)}})
	}).Methods("GET")

	// Task 11: CloudWatch log operations
	api.HandleFunc("/aws/cloudwatch/log", func(w http.ResponseWriter, r *http.Request) {
		client, err := awsutil.NewCloudWatchClient(cfg.AWSRegion, cfg.CloudWatchGroup)
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{Success: true, Data: awsutil.SimulatedCloudWatchResponse(cfg.CloudWatchGroup)})
			return
		}
		msg := fmt.Sprintf("[%s] API health check at %s", cfg.Environment, time.Now().Format(time.RFC3339))
		err = client.PutLogEvent("app-logs", msg)
		if err != nil {
			handler.Respond(w, 200, models.APIResponse{Success: false, Message: err.Error(),
				Data: awsutil.SimulatedCloudWatchResponse(cfg.CloudWatchGroup)})
			return
		}
		handler.Respond(w, 200, models.APIResponse{Success: true, Message: "Log sent to CloudWatch"})
	}).Methods("POST")

	fmt.Printf("\n[Server] Employee API on http://0.0.0.0:%s (%s)\n", cfg.Port, cfg.Environment)
	fmt.Println("  CRUD:   /api/employees")
	fmt.Println("  AWS S3: POST /api/aws/s3/upload  |  GET /api/aws/s3/list")
	fmt.Println("  AWS RDS: GET /api/aws/rds/test")
	fmt.Println("  AWS Secrets: GET /api/aws/secrets")
	fmt.Println("  AWS CW: POST /api/aws/cloudwatch/log")
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
