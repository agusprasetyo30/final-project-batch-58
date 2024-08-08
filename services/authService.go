package services

import (
	"final-project/model"
	"final-project/repository"
)

type AuthService interface {
	Authenticate(username string) (*model.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) Authenticate(username string) (*model.User, error) {
	// hashedPassword := helper.GeneratePassword(password)
	user, err := s.userRepository.GetUserByUsernameAndPassword(username)

	if err != nil {
		return nil, err
	}
	return user, nil
}
