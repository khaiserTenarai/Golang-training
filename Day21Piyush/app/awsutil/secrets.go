// Task 8: Retrieve a secret from AWS Secrets Manager

package awsutil

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretsClient struct {
	svc *secretsmanager.SecretsManager
}

func NewSecretsClient(region string) (*SecretsClient, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, fmt.Errorf("AWS session failed: %w", err)
	}
	return &SecretsClient{svc: secretsmanager.New(sess)}, nil
}

// Task 8: Get secret value by name
func (c *SecretsClient) GetSecret(secretName string) (map[string]string, error) {
	result, err := c.svc.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return nil, fmt.Errorf("Secrets Manager failed: %w", err)
	}
	var secrets map[string]string
	if err := json.Unmarshal([]byte(*result.SecretString), &secrets); err != nil {
		return nil, fmt.Errorf("failed to parse secret JSON: %w", err)
	}
	return secrets, nil
}

// Simulated response for local dev
func SimulatedSecretResponse(secretName string) map[string]interface{} {
	return map[string]interface{}{
		"simulated":   true,
		"message":     "Secrets Manager simulated (no AWS credentials)",
		"secret_name": secretName,
		"cli_cmd":     fmt.Sprintf("aws secretsmanager get-secret-value --secret-id %s", secretName),
		"sample_value": map[string]string{
			"db_host":     "day21-postgres.xxxxx.ap-south-1.rds.amazonaws.com",
			"db_user":     "postgres",
			"db_password": "********",
			"db_name":     "day21db",
		},
	}
}
