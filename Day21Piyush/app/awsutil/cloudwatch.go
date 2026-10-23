// Task 11: CloudWatch log operations from Go

package awsutil

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type CloudWatchClient struct {
	svc      *cloudwatchlogs.CloudWatchLogs
	logGroup string
}

func NewCloudWatchClient(region, logGroup string) (*CloudWatchClient, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, fmt.Errorf("AWS session failed: %w", err)
	}
	return &CloudWatchClient{svc: cloudwatchlogs.New(sess), logGroup: logGroup}, nil
}

// Write a log event to CloudWatch
func (c *CloudWatchClient) PutLogEvent(streamName, message string) error {
	// Ensure log stream exists
	c.svc.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(c.logGroup),
		LogStreamName: aws.String(streamName),
	})

	_, err := c.svc.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(c.logGroup),
		LogStreamName: aws.String(streamName),
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(message),
				Timestamp: aws.Int64(time.Now().UnixMilli()),
			},
		},
	})
	return err
}

// Read recent log events
func (c *CloudWatchClient) GetLogs(streamName string, limit int) ([]map[string]interface{}, error) {
	result, err := c.svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String(c.logGroup),
		LogStreamName: aws.String(streamName),
		Limit:         aws.Int64(int64(limit)),
	})
	if err != nil {
		return nil, err
	}
	var logs []map[string]interface{}
	for _, event := range result.Events {
		logs = append(logs, map[string]interface{}{
			"message":   *event.Message,
			"timestamp": time.UnixMilli(*event.Timestamp).Format(time.RFC3339),
		})
	}
	return logs, nil
}

func SimulatedCloudWatchResponse(logGroup string) map[string]interface{} {
	return map[string]interface{}{
		"simulated": true,
		"message":   "CloudWatch simulated (no AWS credentials)",
		"log_group": logGroup,
		"cli_cmd":   fmt.Sprintf("aws logs tail %s --follow", logGroup),
	}
}
