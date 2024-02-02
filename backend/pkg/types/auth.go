package types

import validation "github.com/go-ozzo/ozzo-validation"

// LoginRequest defines the request body for the login request.
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse defines the response body for the login request.
type LoginResponse struct {
	Token string `json:"token"`
}

// SignupRequest defines the request body for the signup request.
type SignupRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

// Validate validates the request body for the LoginRequest.
func (request LoginRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.UserName,
			validation.Required.Error("Username cannot be empty")),
		validation.Field(&request.Password,
			validation.Required.Error("Password cannot be empty")))
}

// Validate validates the request body for the SignupRequest.
func (request SignupRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.UserName,
			validation.Required.Error("Username cannot be empty"),
			validation.Length(4, 32)),
		validation.Field(&request.Password,
			validation.Required.Error("Password cannot be empty"),
			validation.Length(8, 128)),
		validation.Field(&request.Name,
			validation.Required.Error("Name cannot be empty"),
			validation.Length(2, 64)),
		validation.Field(&request.Email,
			validation.Required.Error("Email cannot be empty"),
			validation.Length(4, 128)),
		validation.Field(&request.Address,
			validation.Length(2, 128)))
}
