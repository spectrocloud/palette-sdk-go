package awsinfra

import (
	"fmt"
	"os"
)

// ResolveCredentials reads AWS credentials from environment variables.
func ResolveCredentials(region string) (Credentials, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if accessKey == "" || secretKey == "" {
		return Credentials{}, fmt.Errorf("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are required")
	}

	if region == "" {
		region = os.Getenv("AWS_REGION")
	}
	if region == "" {
		region = "us-east-1"
	}

	return Credentials{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
	}, nil
}
