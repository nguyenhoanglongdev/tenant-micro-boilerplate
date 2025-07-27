package cognito

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/aws" 
	
	"fmt"
	"context"
)

type CognitoService struct {
	client *cognitoidentityprovider.Client
}

func (cs *CognitoService) verify(ctx context.Context, userPoolIdStr string) error {
	input := &cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: aws.String(userPoolIdStr),
	}

	_, err := cs.client.DescribeUserPool(ctx, input)
	return err
}

func NewClient(ctx context.Context, awsConfig aws.Config, userPoolId string) (*CognitoService, error) {
	client := cognitoidentityprovider.NewFromConfig(awsConfig)
	cognitoService := &CognitoService{client: client}

	if err := cognitoService.verify(ctx, userPoolId); err != nil {
		return nil, fmt.Errorf("failed to verify Cognito client: %w", err)
	}

	return cognitoService, nil
}