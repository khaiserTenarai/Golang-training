// Tasks 9, 10: S3 bucket operations — upload employee report.
// Uses AWS SDK. When running locally without AWS credentials, returns a demo response.

package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"employee-api/config"
	"employee-api/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Handler struct {
	cfg config.Config
}

func NewS3Handler(cfg config.Config) *S3Handler {
	return &S3Handler{cfg: cfg}
}

// POST /api/s3/upload-report — uploads employee CSV report to S3
func (h *S3Handler) UploadReport(w http.ResponseWriter, r *http.Request) {
	// Generate CSV content
	csvContent := "ID,Name,Email,Department,Salary\n1,Piyush,piyush@test.com,Engineering,60000\n"
	fileName := fmt.Sprintf("reports/employee_report_%s.csv", time.Now().Format("2006-01-02_150405"))

	// Try AWS upload
	sess, err := session.NewSession(&aws.Config{Region: aws.String(h.cfg.AWSRegion)})
	if err != nil {
		// No AWS credentials — return demo response
		respond(w, 200, models.APIResponse{
			Success: true,
			Message: "S3 upload simulated (no AWS credentials configured)",
			Data: map[string]string{
				"bucket":   h.cfg.S3Bucket,
				"key":      fileName,
				"region":   h.cfg.AWSRegion,
				"size":     fmt.Sprintf("%d bytes", len(csvContent)),
				"command":  fmt.Sprintf("aws s3 cp employee_report.csv s3://%s/%s", h.cfg.S3Bucket, fileName),
			},
		})
		return
	}

	svc := s3.New(sess)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(h.cfg.S3Bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader([]byte(csvContent)),
		ContentType: aws.String("text/csv"),
	})
	if err != nil {
		respond(w, 200, models.APIResponse{
			Success: false,
			Message: "S3 upload failed (check AWS credentials): " + err.Error(),
			Data: map[string]string{
				"bucket": h.cfg.S3Bucket,
				"key":    fileName,
			},
		})
		return
	}

	respond(w, 200, models.APIResponse{
		Success: true,
		Message: "Report uploaded to S3",
		Data: map[string]string{
			"bucket": h.cfg.S3Bucket,
			"key":    fileName,
			"url":    fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", h.cfg.S3Bucket, h.cfg.AWSRegion, fileName),
		},
	})
}
