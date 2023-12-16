package services

import (
	"errors"

	"github.com/BerdanAkbulut/task-app-backend/entity"
	"github.com/BerdanAkbulut/task-app-backend/jwt"
	"github.com/BerdanAkbulut/task-app-backend/repository"
	"github.com/BerdanAkbulut/task-app-backend/requests"
	"github.com/BerdanAkbulut/task-app-backend/responses"
)

type authService struct {
	userRepository repository.UserRepository
}

type AuthService interface {
	Register(registerRequest *requests.RegisterRequest) (*responses.AuthResponse, error)
	Authenticate(authenticateRequest *requests.AuthenticateRequest) (*responses.AuthResponse, error)
}

func NewAuthService(ur repository.UserRepository) AuthService {
	return &authService{
		userRepository: ur,
	}
}

func (s *authService) Register(registerRequest *requests.RegisterRequest) (*responses.AuthResponse, error) {
	u := &entity.User{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Email:     registerRequest.Email,
		Password:  registerRequest.Password,
	}
	err := s.userRepository.Save(u)
	if err != nil {
		return nil, err
	}
	token, ex := jwt.GenerateJWT(u.Email)

	if ex != nil {
		return nil, ex
	}

	return &responses.AuthResponse{
		Token: *token,
	}, nil
}

func (s *authService) Authenticate(authenticateRequest *requests.AuthenticateRequest) (*responses.AuthResponse, error) {
	user, err := s.userRepository.FindByEmail(authenticateRequest.Email)
	if err != nil {
		return nil, errors.New("Email with user not found: " + authenticateRequest.Email)
	}
	isPwRight := user.CheckPassword(authenticateRequest.Password)
	if !isPwRight {
		return nil, errors.New("Wrong password")
	}
	token, err := jwt.GenerateJWT(user.Email)
	if err != nil {
		return nil, err
	}
	return &responses.AuthResponse{Token: *token}, nil
}
