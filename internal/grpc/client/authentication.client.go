package grpc_client

import (
	"context"
	"fmt"
	"user_service/global"
	v1 "user_service/internal/grpc/gen_code/authentication.v1"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IAuthenticationClient interface {
	ValidateToken(jwtToken string) (bool, error)
}

type AuthenticationClient struct {
	authenticationClient v1.AuthenticationServiceClient
}

func NewAuthenticationClient() (IAuthenticationClient, error) {
	authenClientConfig := global.Config.GrpcAuthenticationClient
	logger := global.Logger

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", authenClientConfig.Host, authenClientConfig.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.ErrorString("grpc", zap.String("error", fmt.Sprintf("Failed to connect to authentication service: %v", err)))
		return nil, err
	}

	client := v1.NewAuthenticationServiceClient(conn)
	if client == nil {
		conn.Close()
		logger.ErrorString("grpc", zap.String("error", "Failed to create authentication service client"))
		return nil, fmt.Errorf("failed to create authentication service client: %w", err)
	}

	return &AuthenticationClient{
		authenticationClient: v1.NewAuthenticationServiceClient(conn),
	}, nil
}

func (c *AuthenticationClient) ValidateToken(jwtToken string) (bool, error) {
	req := &v1.ValidateTokenRequest{
		JwtToken: jwtToken,
	}

	resp, err := c.authenticationClient.ValidateToken(context.Background(), req)
	if err != nil {
		global.Logger.ErrorString("grpc", zap.String("error", fmt.Sprintf("Failed to validate token: %v", err)))
		return false, err
	}

	if resp == nil {
		return false, fmt.Errorf("received nil response from authentication service")
	}

	return resp.Result, nil
}
