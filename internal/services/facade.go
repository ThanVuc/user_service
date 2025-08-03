package services

import (
	"context"
	"user_service/internal/repos"
	"user_service/proto/user"
)

type (
	UserService interface {
		GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error)
	}
)

func NewUserService(userRepo repos.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
