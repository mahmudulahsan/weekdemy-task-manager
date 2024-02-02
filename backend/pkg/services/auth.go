package services

import (
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/models"
	"weekdemy-task-manager-backend/pkg/types"
	"weekdemy-task-manager-backend/pkg/utils"
	"errors"
)

// authService defines the methods of the domain.IAuthService interface.
type authService struct {
	userRepo domain.IUserRepo
}

// AuthServiceInstance returns a new instance of the authService struct.
func AuthServiceInstance(userRepo domain.IUserRepo) domain.IAuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// LoginUser returns a JWT token for the user if the credentials are correct.
func (service *authService) LoginUser(loginRequest *types.LoginRequest) (*types.LoginResponse, error) {
	// Check if user exists
	existingUser, err := service.userRepo.GetUser(&loginRequest.UserName)
	if err != nil {
		return nil, errors.New("user does not exist")
	}

	// Check if password is correct
	if err := utils.CheckPassword(existingUser.PasswordHash, loginRequest.Password); err != nil {
		return nil, errors.New("incorrect password")
	}

	// Generate JWT token
	token, err := utils.GetJwtForUser(existingUser.Username)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token: token,
	}, nil

}

// SignupUser creates a new user with the given user details.
func (service *authService) SignupUser(registerRequest *types.SignupRequest) error {
	// get hashed password
	passwordHash, err := utils.GetPasswordHash(registerRequest.Password)
	if err != nil {
		return err
	}

	// create user
	user := &models.UserDetail{
		Username:     registerRequest.UserName,
		PasswordHash: passwordHash,
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		Address:      registerRequest.Address,
	}
	if err := service.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}
