package services

import (
	"context"
	"user_service/internal/grpc/mapper"
	"user_service/internal/grpc/repos"
	"user_service/internal/grpc/utils"
	"user_service/proto/common"
	"user_service/proto/user"

	"github.com/thanvuc/go-core-lib/log"
)

type userService struct {
	userRepo   repos.UserRepo
	userMapper mapper.UserMapper
	logger     log.Logger
}

func (us *userService) GetUserProfile(ctx context.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error) {
    userPr, err := us.userRepo.GetUserProfile(ctx, req)
    if err != nil {
        return &user.GetUserProfileResponse{
            Error: utils.DatabaseError(ctx, err),
        }, nil
    }

    if userPr == nil  {
        return &user.GetUserProfileResponse{
            Error: &common.Error{
                Code:    1,
                Message: "User not found",
            },
        }, nil
    }

    resp := &user.GetUserProfileResponse{
        Profiles: us.userMapper.ConvertDbUserPrifileRowToGrpcUser(userPr),
    }
    return resp, nil
}


func (us *userService) UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*common.EmptyResponse, error) {
	userId, err := us.userRepo.UpdateUserProfile(ctx,req)
	if err != nil {
		return &common.EmptyResponse{
			Success: utils.ToBoolPointer(false),
			Message: utils.ToStringPointer("Failed update User!"),
			Error:   utils.DatabaseError(ctx, err),
		}, err
	}

	if userId == nil {
		return &common.EmptyResponse{
			Success: utils.ToBoolPointer(false),
			Message: utils.ToStringPointer("User not found!"),
			Error:   utils.NotFoundError(ctx, err),
		}, nil
	}

	return &common.EmptyResponse{
		Success: utils.ToBoolPointer(true),
		Message: utils.ToStringPointer("Update User successfully!"),
	}, nil
}
