package services

import (
	"user_service/internal/repos"
)

type UserService struct {
	userRepo *repos.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repos.NewUserRepo(),
	}
}

func (us *UserService) GetUserInfo() map[string]interface{} {
	return us.userRepo.GetUserInfo()
}
