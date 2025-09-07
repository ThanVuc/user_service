package mapper

import (
	"user_service/internal/grpc/database"
	"user_service/proto/user"
)
type (
	UserMapper interface { 
		ConvertDbUserPrifileRowToGrpcUser(user *[]database.GetUserProfileRow) *user.UserProfileItem
	}
)


func NewUserMapper() UserMapper {
	return &userMapper{}
}
