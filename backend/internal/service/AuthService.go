package service

import (
	"task-management/internal/repository"
	"task-management/internal/utils"
)

type AuthService struct {
	userRepository *repository.UserRepository
	passwordService *PasswordService
}

func NewAuthService(userRepo *repository.UserRepository, passwordSvc *PasswordService) *AuthService {
	return &AuthService{
		userRepository: userRepo,
		passwordService: passwordSvc,
	}
}

func (s *AuthService) Login (email, password string) (string, map[string]interface{}, error) {
	if email == "" || password == "" {
		return "", nil, ErrInvalidInput
	}
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", nil, ErrUserNotFound
	}

	err = s.passwordService.CompareHashAndPassword(user.Password, password) 
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Fullname, user.Role)
	if err != nil {
		return "", nil, ErrGenerateTokenFailed
	}

	userData := map[string]interface{}{
		"id":       user.ID,
		"email":    user.Email,
		"fullname": user.Fullname,
		"role":     user.Role,
	}
	return token, userData, nil
}