// Task 6: Test RDS connectivity from Go

package awsutil

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type RDSTestResult struct {
	Connected bool   `json:"connected"`
	Host      string `json:"host"`
	Latency   string `json:"latency"`
	Version   string `json:"version,omitempty"`
	Error     string `json:"error,omitempty"`
}

// Task 6: Test connectivity to PostgreSQL (RDS or local)
func TestRDSConnection(databaseURL string) RDSTestResult {
	if databaseURL == "" {
		return RDSTestResult{
			Connected: false,
			Host:      "none",
			Error:     "DATABASE_URL not set",
		}
	}

	start := time.Now()
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return RDSTestResult{Connected: false, Error: err.Error()}
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return RDSTestResult{
			Connected: false,
			Latency:   time.Since(start).String(),
			Error:     err.Error(),
		}
	}

	var version string
	db.QueryRow("SELECT version()").Scan(&version)

	return RDSTestResult{
		Connected: true,
		Host:      databaseURL[:min(50, len(databaseURL))] + "...",
		Latency:   time.Since(start).Round(time.Millisecond).String(),
		Version:   version,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SimulatedRDSResponse() map[string]interface{} {
	return map[string]interface{}{
		"simulated": true,
		"message":   "RDS test simulated (no DATABASE_URL set)",
		"cli_cmd":   "aws rds describe-db-instances --db-instance-identifier day21-postgres",
		"sample_url": fmt.Sprintf("postgres://user:pass@day21-postgres.xxxxx.ap-south-1.rds.amazonaws.com:5432/day21db"),
	}
}
