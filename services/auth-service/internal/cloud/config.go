package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws" 
	"github.com/aws/aws-sdk-go-v2/config"
)

// LoadAWSConfig loads the default AWS configuration with optional customizations.
func LoadAWSConfig(ctx context.Context, region string) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}

	if region != "" {
		opts = append(opts, config.WithRegion(region))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}