package types

import validation "github.com/go-ozzo/ozzo-validation"

// ReadBookResponse defines the response body for the read book request.
type ReadBookResponse struct {
	ID          uint   `json:"id"`
	BookName    string `json:"bookName"`
	AuthorID    uint   `json:"authorID"`
	Publication string `json:"publication,omitempty"`
}

// CreateBookRequest defines the request body for the create book request.
// ID is redacted because it is auto-incremented.
type CreateBookRequest struct {
	ReadBookResponse
	ID uint `json:"-"`
}

// UpdateBookRequest defines the request body for the update book request.
// ID is not allowed to be updated.
type UpdateBookRequest struct {
	ReadBookResponse
	ID uint `json:"-"`
}

// Validate validates the request body for the CreateBookRequest.
// BookName is required.
// AuthorID is required.
// Publication is optional.
func (request CreateBookRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 64)),
		validation.Field(&request.AuthorID,
			validation.Required.Error("Author ID cannot be empty")),
		validation.Field(&request.Publication,
			validation.Length(2, 64)))
}

// Validate validates the request body for the UpdateBookRequest.
func (request UpdateBookRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.BookName,
			validation.Length(1, 64)),
		validation.Field(&request.Publication,
			validation.Length(2, 64)))
}
