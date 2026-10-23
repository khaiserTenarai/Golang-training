#!/bin/bash
# Day 21: AWS CLI Commands — Tasks 1, 2, 4, 9, 11
# Run each section independently. Requires: aws configure

REGION="ap-south-1"
BUCKET="day21-piyush-employee-data"

# ============================================================
# Task 1: Configure AWS CLI
# ============================================================
echo "=== Task 1: Configure AWS CLI ==="
echo "Run: aws configure"
echo "  Enter: Access Key ID, Secret Access Key, Region (ap-south-1), Output (json)"
echo "  Config stored in: ~/.aws/credentials and ~/.aws/config"
echo ""

# ============================================================
# Task 2: Verify AWS Identity
# ============================================================
echo "=== Task 2: Verify AWS Identity ==="
aws sts get-caller-identity
echo ""

# ============================================================
# Task 4: S3 Operations Using CLI
# ============================================================
echo "=== Task 4: S3 CLI Operations ==="

# Create bucket
aws s3 mb s3://$BUCKET --region $REGION

# Upload a file
echo "ID,Name,Salary" > /tmp/test_report.csv
echo "1,Piyush,60000" >> /tmp/test_report.csv
aws s3 cp /tmp/test_report.csv s3://$BUCKET/reports/test_report.csv

# List objects
aws s3 ls s3://$BUCKET/reports/

# Download file
aws s3 cp s3://$BUCKET/reports/test_report.csv /tmp/downloaded_report.csv

# View bucket contents
aws s3 ls s3://$BUCKET --recursive

echo ""

# ============================================================
# Task 9: EC2 Deployment Script
# ============================================================
echo "=== Task 9: EC2 Deployment ==="
echo "# Build Go binary for Linux:"
echo "  cd app && GOOS=linux GOARCH=amd64 go build -o employee-api main.go"
echo ""
echo "# Copy to EC2:"
echo "  scp -i key.pem employee-api ec2-user@<IP>:/opt/employee-api/"
echo ""
echo "# SSH and run:"
echo "  ssh -i key.pem ec2-user@<IP>"
echo "  sudo systemctl start employee-api"
echo ""

# ============================================================
# Task 11: CloudWatch Log Commands
# ============================================================
echo "=== Task 11: CloudWatch Logs ==="

# Create log group
aws logs create-log-group --log-group-name /employee-api/app --region $REGION 2>/dev/null
echo "Log group created: /employee-api/app"

# Tail logs (real-time)
echo "To tail logs: aws logs tail /employee-api/app --follow --region $REGION"

# Search logs
echo "To search: aws logs filter-log-events --log-group-name /employee-api/app --filter-pattern 'ERROR' --region $REGION"
