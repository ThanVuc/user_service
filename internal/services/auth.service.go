package services

import (
	"context"
	"user_service/internal/repos"
)

type IAuthService interface {
	CreateAuthor(c context.Context) bool
	CreateUserAndAuthor(c context.Context) bool
}

type AuthService struct {
	authRepo repos.IAuthRepo
}

func NewAuthService(
	authRepo repos.IAuthRepo,
) IAuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

func (us *AuthService) CreateAuthor(c context.Context) bool {
	return us.authRepo.CreateAuthor(c)
}

func (us *AuthService) CreateUserAndAuthor(c context.Context) bool {
	return us.authRepo.CreateUserAndAuthor(c)
}
