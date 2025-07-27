package cognito

import (
	internalAws "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/aws"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/config"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/aws" 
	
	"fmt"
	"context"
)

type CognitoService struct {
	client *cognitoidentityprovider.Client
}

// verify performs a simple call to check if the Cognito client is properly configured and authorized.
func (cs *CognitoService) verify(ctx context.Context) error {
	config := config.LoadConfig()

	input := &cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: aws.String(config.UserPoolId),
	}

	_, err := cs.client.DescribeUserPool(ctx, input)
	return err
}

// NewClient initializes the CognitoService and verifies the client by making a lightweight API call.
func NewClient(ctx context.Context, region string) (*CognitoService, error) {
	awsConfig, err := internalAws.LoadAWSConfig(ctx, region)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := cognitoidentityprovider.NewFromConfig(awsConfig)
	cognitoService := &CognitoService{client: client}

	if err := cognitoService.verify(ctx); err != nil {
		return nil, fmt.Errorf("failed to verify Cognito client: %w", err)
	}

	return cognitoService, nil
}