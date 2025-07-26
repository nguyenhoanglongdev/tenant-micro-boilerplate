package cognito

import (
	"context"
	"errors"
)

type CognitoUserPool struct {
	client     *Client
	userPoolId string
	clientId   string
}

func NewUserPool(client *Client, userPoolId, clientId string) *CognitoUserPool {
	return &CognitoUserPool{client: client, userPoolId: userPoolId, clientId: clientId}
}

func (u *CognitoUserPool) Register(ctx context.Context, username, password string) error {
	// Implement AWS Cognito sign-up logic here,
	// use u.client.svc.SignUp API call
	return errors.New("not implemented")
}

func (u *CognitoUserPool) Login(ctx context.Context, username, password string) (string, error) {
	// Implement AWS Cognito initiate auth logic,
	// return JWT token on success
	return "", errors.New("not implemented")
}
