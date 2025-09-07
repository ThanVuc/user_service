package services

import (
	"context"
	"user_service/global"
	"user_service/internal/grpc/mapper"
	"user_service/internal/grpc/repos"
	"user_service/proto/common"
	"user_service/proto/user"
)

type (
	UserService interface {
		GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error)
		UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*common.EmptyResponse, error)
	}
)


func NewUserService(
	userRepo repos.UserRepo,
	userMapper mapper.UserMapper,
) UserService {
	return &userService{
		userRepo: userRepo,
		userMapper : userMapper,
		logger:     global.Logger,
	}
}



