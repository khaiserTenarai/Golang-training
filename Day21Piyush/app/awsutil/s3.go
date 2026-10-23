// Tasks 3, 5: Configure AWS SDK for Go + S3 operations from Go
// Task 3 — AWS SDK session creation
// Task 5 — Upload, download, list S3 objects from Go

package awsutil

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	svc    *s3.S3
	bucket string
	region string
}

// Task 3: Configure AWS SDK — create session
func NewS3Client(region, bucket string) (*S3Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("AWS session failed: %w", err)
	}
	return &S3Client{svc: s3.New(sess), bucket: bucket, region: region}, nil
}

// Task 5: Upload to S3
func (c *S3Client) Upload(key string, data []byte, contentType string) (string, error) {
	_, err := c.svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(c.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("S3 upload failed: %w", err)
	}
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", c.bucket, c.region, key)
	return url, nil
}

// Task 5: Download from S3
func (c *S3Client) Download(key string) ([]byte, error) {
	result, err := c.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("S3 download failed: %w", err)
	}
	defer result.Body.Close()
	return io.ReadAll(result.Body)
}

// Task 5: List S3 objects
func (c *S3Client) ListObjects(prefix string) ([]map[string]interface{}, error) {
	result, err := c.svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("S3 list failed: %w", err)
	}
	var objects []map[string]interface{}
	for _, obj := range result.Contents {
		objects = append(objects, map[string]interface{}{
			"key":           *obj.Key,
			"size":          *obj.Size,
			"last_modified": obj.LastModified.Format(time.RFC3339),
		})
	}
	return objects, nil
}

// Simulated response when AWS isn't configured (local dev)
func SimulatedS3Response(bucket, key, region string) map[string]interface{} {
	return map[string]interface{}{
		"simulated": true,
		"message":   "S3 operation simulated (no AWS credentials)",
		"bucket":    bucket,
		"key":       key,
		"region":    region,
		"cli_cmd":   fmt.Sprintf("aws s3 cp file.csv s3://%s/%s", bucket, key),
	}
}
