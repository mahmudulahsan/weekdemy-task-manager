package types

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ReadAuthorResponse defines the response body for the read author request.
type ReadAuthorResponse struct {
	ID          uint   `json:"id"`
	AuthorName  string `json:"authorName"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// CreateAuthorRequest defines the request body for the create author request.
// ID is redacted because it is auto-incremented.
type CreateAuthorRequest struct {
	ReadAuthorResponse
	ID uint `json:"-"`
}

// UpdateAuthorRequest defines the request body for the update author request.
// ID is not allowed to be updated.
type UpdateAuthorRequest struct {
	ReadAuthorResponse
	ID uint `json:"-"`
}

// Validate validates the request body for the CreateAuthorRequest.
// AuthorName is required.
// Address and PhoneNumber are optional.
func (request CreateAuthorRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.AuthorName,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(2, 64)),
		validation.Field(&request.Address,
			validation.Length(2, 128)),
		validation.Field(&request.PhoneNumber,
			validation.Length(8, 32)))

}

// Validate validates the request body for the UpdateAuthorRequest.
func (request UpdateAuthorRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.AuthorName,
			validation.Length(2, 64)),
		validation.Field(&request.Address,
			validation.Length(2, 128)),
		validation.Field(&request.PhoneNumber,
			validation.Length(8, 32)))
}
