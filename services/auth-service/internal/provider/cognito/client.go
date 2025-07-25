package cognito

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type Client struct {
    svc *cognitoidentityprovider.Client
}

func NewClient(ctx context.Context, region string) (*Client, error) {
    cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
    if err != nil {
        return nil, err
    }
    svc := cognitoidentityprovider.NewFromConfig(cfg)
    return &Client{svc: svc}, nil
}
