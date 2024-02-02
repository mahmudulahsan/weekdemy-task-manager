package domain

import (
	"weekdemy-task-manager-backend/pkg/models"
	"weekdemy-task-manager-backend/pkg/types"
)

type IBookRepo interface {
	GetFilteredBooks(request map[string]string) ([]models.BookDetail, error)
	GetBook(bookID uint) (*models.BookDetail, error)
	CreateBook(book *models.BookDetail) error
	UpdateBook(book *models.BookDetail) error
	DeleteBook(bookID uint) error
	DeleteBooksByAuthorID(authorID uint) error
}

type IBookService interface {
	GetFilteredBooks(request map[string]string) ([]types.ReadBookResponse, error)
	GetBook(bookID uint) (*types.ReadBookResponse, error)
	CreateBook(request *types.CreateBookRequest) error
	UpdateBook(bookID uint, request *types.UpdateBookRequest) error
	DeleteBook(bookID uint) error
}
