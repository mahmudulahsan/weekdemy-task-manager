package controllers

import (
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// IBookController is an interface that defines the methods implemented by the BookController struct.
type IBookController interface {
	CreateBook(e echo.Context) error
	GetBook(e echo.Context) error
	GetFilteredBooks(e echo.Context) error
	UpdateBook(e echo.Context) error
	DeleteBook(e echo.Context) error
}

// BookController defines the methods of the IBookController interface.
type BookController struct {
	bookSvc domain.IBookService
}

// NewBookController returns a new instance of the BookController struct.
func NewBookController(bookSvc domain.IBookService) BookController {
	return BookController{
		bookSvc: bookSvc,
	}
}

// CreateBook handles the create book request.
//
//	request body:
//		{
//			"bookName":    "Cool Book",	// required
//			"authorID":    1,		// required
//			"publication": "Cool pub"	// optional
//		}
func (controller *BookController) CreateBook(e echo.Context) error {
	// bind the request body to the CreateBookRequest struct
	request := &types.CreateBookRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := request.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := controller.bookSvc.CreateBook(request); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "book was created successfully")
}

// GetBook handles the get book request.
//
//	path params:
//		id:	// required
func (controller *BookController) GetBook(e echo.Context) error {
	// validate the request
	tempBookID := e.Param("id")
	bookID, err := strconv.ParseUint(tempBookID, 0, 0)
	if err != nil && tempBookID != "" {
		return e.JSON(http.StatusBadRequest, "enter a valid book ID")
	}

	// pass the request to the service layer
	response, err := controller.bookSvc.GetBook(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, response)
}

// GetFilteredBooks handles the get filtered books request.
//
//	query params:
//		id:		// optional
//		bookName:	// optional
//		authorID:	// optional
//		publication:	// optional
//
// example: ?authorID=2&publication=xyz
func (controller *BookController) GetFilteredBooks(e echo.Context) error {
	// prepare the filter request to handle the query params i.e ?authorID=2&publication=xyz
	request := make(map[string]string)
	if e.QueryParam("id") != "" {
		if _, err := strconv.ParseUint(e.QueryParam("id"), 0, 0); err != nil {
			return e.JSON(http.StatusBadRequest, "enter a valid book ID")
		}
		request["ID"] = e.QueryParam("id")
	}
	if e.QueryParam("bookName") != "" {
		request["BookName"] = e.QueryParam("bookName")
	}
	if e.QueryParam("authorID") != "" {
		if _, err := strconv.ParseUint(e.QueryParam("authorID"), 0, 0); err != nil {
			return e.JSON(http.StatusBadRequest, "enter a valid author ID")
		}
		request["AuthorID"] = e.QueryParam("authorID")
	}
	if e.QueryParam("publication") != "" {
		request["Publication"] = e.QueryParam("publication")
	}

	// pass the request to the service layer
	response, err := controller.bookSvc.GetFilteredBooks(request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, response)

}

// UpdateBook handles the update book request.
//
//	path params:
//		id:	// required
//
//	request body:
//		{
//			"bookName":    "Cool Book",	// required
//			"authorID":    1,		// required
//			"publication": "Cool pub"	// optional
//		}
func (controller *BookController) UpdateBook(e echo.Context) error {
	// bind the request body to the CreateBookRequest struct
	request := &types.UpdateBookRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	tempBookID := e.Param("id")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "enter a valid book ID")
	}

	// pass the request to the service layer
	if err := controller.bookSvc.UpdateBook(uint(bookID), request); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "book was updated successfully")
}

// DeleteBook handles the delete book request.
//
//	path params:
//		id:	// required
func (controller *BookController) DeleteBook(e echo.Context) error {
	// validate the request
	tempBookID := e.Param("id")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "enter a valid book ID")
	}

	// pass the request to the service layer
	if err := controller.bookSvc.DeleteBook(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "book was deleted successfully")
}
