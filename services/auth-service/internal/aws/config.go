package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws" 
	"github.com/aws/aws-sdk-go-v2/config"
)

// LoadAWSConfig loads the default AWS configuration with optional customizations.
func LoadAWSConfig(ctx context.Context, region string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}
