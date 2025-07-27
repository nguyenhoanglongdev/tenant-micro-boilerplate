package cognito

import (
	"context"
	"log"
	
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/model"
)

type CognitoUserPool struct {
	Client     *cognitoidentityprovider.Client
	UserPoolID string
	ClientID   string
}

func (c *CognitoUserPool) Register(ctx context.Context, user *model.User) error {
	_, err := c.Client.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: &c.ClientID,
		Username: &user.Email,
		Password: &user.Password,
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(user.Email)},
		},
	})
	
	if err != nil {
		return err
	}

	// AdminConfirmSignUp uses the UserPoolID and user's email from the struct and argument
	_, err = c.Client.AdminConfirmSignUp(ctx, &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(c.UserPoolID),
		Username:   aws.String(user.Email),
	})
	if err != nil {
		log.Printf("AdminConfirmSignUp failed for user %s: %v", user.Email, err)
	} else {
		log.Printf("AdminConfirmSignUp succeeded for user %s", user.Email)
	}

	return err

}

func (c *CognitoUserPool) Login(ctx context.Context, email, password string) (*model.User, error) {
	resp, err := c.Client.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: &c.ClientID,
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		Email: email,
		AccessToken:  *resp.AuthenticationResult.AccessToken,
		// optionally include tokens from resp.AuthenticationResult
	}, nil
}

func NewCognitoUserPool(client *CognitoService, userPoolID, clientID string) provider.AuthInterface {
	return &CognitoUserPool{
		Client:     client.client,
		UserPoolID: userPoolID,
		ClientID:   clientID,
	}
}