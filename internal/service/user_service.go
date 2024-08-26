package service

import (
	"errors"
	"fmt"
	"user-management-backend/internal/model"
	"user-management-backend/internal/repository"
	"user-management-backend/pkg/validator"
)

type UserService interface {
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	GetUserByID(id uint) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) CreateUser(user *model.User) error {
	// Validate user data
	if err := validator.ValidateUser(user); err != nil {
		return err
	}

	fmt.Println("Validated the user", user)

	// Check if user_name already exists
	existingUser, _ := s.repo.GetByUserName(user.UserName)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user *model.User) error {
	// Validate user data
	if err := validator.ValidateUser(user); err != nil {
		return err
	}

	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
