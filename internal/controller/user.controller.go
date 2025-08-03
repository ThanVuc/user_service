package controller

import (
	"context"
	"user_service/internal/services"
	"user_service/internal/utils"
	"user_service/proto/user"
)

type UserController struct {
	user.UnimplementedUserServiceServer
	userService services.UserService
}

func NewUserController(
	userService services.UserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error) {
	println("GetUserProfile called with request:", req)
	return utils.WithSafePanic(ctx, req, uc.userService.GetUserProfile)
}
