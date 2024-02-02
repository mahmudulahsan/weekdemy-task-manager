package controllers

import (
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// IAuthorController is an interface that defines the methods implemented by the AuthorController struct.
type IAuthorController interface {
	CreateAuthor(e echo.Context) error
	GetAuthor(e echo.Context) error
	GetFilteredAuthors(e echo.Context) error
	UpdateAuthor(e echo.Context) error
	DeleteAuthor(e echo.Context) error
}

// AuthorController defines the methods of the IAuthorController interface.
type AuthorController struct {
	authorSvc domain.IAuthorService
}

// NewAuthorController returns a new instance of the AuthorController struct.
func NewAuthorController(authorSvc domain.IAuthorService) AuthorController {
	return AuthorController{
		authorSvc: authorSvc,
	}
}

// CreateAuthor handles the create author request.
//
//	request body:
//		{
//			"authorName":  "John Doe",    // required
//			"address":     "123 Main St", // optional
//			"phoneNumber": "1234567890"   // optional
//		}
func (controller *AuthorController) CreateAuthor(e echo.Context) error {
	// bind the request body to the CreateAuthorRequest struct
	request := &types.CreateAuthorRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := request.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := controller.authorSvc.CreateAuthor(request); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "author was created successfully")
}

// GetAuthor handles the get author request.
//
//	path params:
//		id	// required
func (controller *AuthorController) GetAuthor(e echo.Context) error {
	// validate the request
	tempAuthorID := e.Param("id")
	authorID, err := strconv.ParseUint(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "enter a valid author ID")
	}

	// pass the request to the service layer
	response, err := controller.authorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, response)
}

// GetFilteredAuthors handles the get filtered authors request.
//
//	query params:
//		id		// optional
//		authorName	// optional
//		address		// optional
//		phoneNumber	// optional
//
// example: ?authorName=John&address=123%20Main%20St
func (controller *AuthorController) GetFilteredAuthors(e echo.Context) error {
	// prepare the filter request to handle the query params i.e ?id=1&authorName=John
	request := make(map[string]string)
	if e.QueryParam("id") != "" {
		if _, err := strconv.ParseUint(e.QueryParam("id"), 0, 0); err != nil {
			return e.JSON(http.StatusBadRequest, "enter a valid author ID")
		}
		request["ID"] = e.QueryParam("id")
	}
	if e.QueryParam("authorName") != "" {
		request["AuthorName"] = e.QueryParam("authorName")
	}
	if e.QueryParam("address") != "" {
		request["Address"] = e.QueryParam("address")
	}
	if e.QueryParam("phoneNumber") != "" {
		request["PhoneNumber"] = e.QueryParam("phoneNumber")
	}

	// pass the request to the service layer
	response, err := controller.authorSvc.GetFilteredAuthors(request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, response)
}

// UpdateAuthor handles the update author request.
//
//	path params:
//		id	// required
//
//	request body:
//		{
//			"authorName":  "John Doe",    // optional
//			"address":     "123 Main St", // optional
//			"phoneNumber": "1234567890"   // optional
//		}
func (controller *AuthorController) UpdateAuthor(e echo.Context) error {
	// bind the request body to the UpdateAuthorRequest struct
	request := &types.UpdateAuthorRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := request.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempAuthorID := e.Param("id")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "enter a valid author ID")
	}

	// pass the request to the service layer
	if err := controller.authorSvc.UpdateAuthor(uint(authorID), request); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "author was updated successfully")
}

// DeleteAuthor handles the delete author request.
//
//	path params:
//		id	// required
func (controller *AuthorController) DeleteAuthor(e echo.Context) error {
	// validate the request
	tempAuthorID := e.Param("id")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "enter a valid author ID")
	}

	// pass the request to the service layer
	if err := controller.authorSvc.DeleteAuthor(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "author and books of author was deleted successfully")
}
