package config

import (
    "os"
)

type Config struct {
    AwsRegion   string
    UserPoolId  string
    ClientId    string
    Port        string
}

func LoadConfig() *Config {
    return &Config{
        AwsRegion:  os.Getenv("AWS_REGION"),
        UserPoolId: os.Getenv("COGNITO_USER_POOL_ID"),
        ClientId:   os.Getenv("COGNITO_CLIENT_ID"),
        Port:       os.Getenv("PORT"),
    }
}
