package services

import (
	"user_service/internal/database"
	"user_service/internal/repos"
)

type IUserService interface {
	GetUsers() ([]database.UserProfile, error)
}

type UserService struct {
	repo repos.IUserRepo
}

func NewUserService(repo repos.IUserRepo) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers() ([]database.UserProfile, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
