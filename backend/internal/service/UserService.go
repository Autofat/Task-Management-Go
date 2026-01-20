package service

import (
	"errors"
	"strings"
	"task-management/internal/model"
	"task-management/internal/repository"
)

type UserService struct{
	userRepository  *repository.UserRepository
	passwordService *PasswordService
}

func NewUserService(userRepo *repository.UserRepository, passwordService *PasswordService) *UserService {
	return &UserService{
		userRepository:  userRepo,
		passwordService: passwordService,
	}
}

func (s *UserService) CreateUser(email, password, fullname, role string) (*model.User,error) {
	if email == "" || password == "" || fullname == ""{
		return nil, errors.New("all fields are required")
	}
	if role == "" {
	role = "USER"
}

	existingEmail, _ := s.userRepository.FindByEmail(email)
	if existingEmail != nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword,err := s.passwordService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
		Fullname: fullname,
		Role:     role,
	}
	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepository.FindByEmail(email)
}

func (s *UserService) UpdateUser(id uint, fullname, role string) error {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return ErrUserNotFound
	}

	if strings.TrimSpace(fullname) == "" || strings.TrimSpace(role) == "" {
		return ErrInvalidInput
	}
	
	if fullname != "" {
		user.Fullname = fullname
	}
	if role != "" {
		user.Role = role
	}
		
	return s.userRepository.Update(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	_, err := s.userRepository.FindByID(id)
	if err != nil {
		return ErrUserNotFound
	}
	return s.userRepository.DeleteById(id)
}