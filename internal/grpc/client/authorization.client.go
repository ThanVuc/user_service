package grpc_client

import (
	"context"
	"fmt"
	"user_service/global"
	v1 "user_service/internal/grpc/gen_code/authorization.v1"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IAuthorizationClient interface {
	CheckPerm(resource string, action string, jwtToken string) (bool, error)
}

type AuthorizationClient struct {
	authorizationClient v1.AuthorizationServiceClient
}

func NewAuthorizationClient() (IAuthorizationClient, error) {
	authorClientConfig := global.Config.GrpcAuthorizationClient
	logger := global.Logger

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", authorClientConfig.Host, authorClientConfig.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.ErrorString("grpc", zap.String("error", fmt.Sprintf("Failed to connect to authorization service: %v", err)))
		return nil, err
	}

	client := v1.NewAuthorizationServiceClient(conn)
	if client == nil {
		conn.Close()
		logger.ErrorString("grpc", zap.String("error", "Failed to create authorization service client"))
		return nil, fmt.Errorf("failed to create authorization service client: %w", err)
	}

	return &AuthorizationClient{
		authorizationClient: v1.NewAuthorizationServiceClient(conn),
	}, nil
}

func (c *AuthorizationClient) CheckPerm(resource string, action string, jwtToken string) (bool, error) {
	req := &v1.CheckPermRequest{
		JwtToken: jwtToken,
		Resource: resource,
		Action:   action,
	}

	resp, err := c.authorizationClient.CheckPerm(context.Background(), req)
	if err != nil {
		global.Logger.ErrorString("grpc", zap.String("error", fmt.Sprintf("Failed to validate token: %v", err)))
		return false, err
	}

	if resp == nil {
		return false, fmt.Errorf("received nil response from authorization service")
	}

	return resp.Allowed, nil
}
