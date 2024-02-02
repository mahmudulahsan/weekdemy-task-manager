package services

import (
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/models"
	"weekdemy-task-manager-backend/pkg/types"
	"errors"
)

// BookService defines the methods of the domain.IBookService interface.
type bookService struct {
	bookRepo   domain.IBookRepo
	authorRepo domain.IAuthorRepo
}

// BookServiceInstance returns a new instance of the BookService struct.
func BookServiceInstance(bookRepo domain.IBookRepo, authorRepo domain.IAuthorRepo) domain.IBookService {
	return &bookService{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}

// GetFilteredBooks returns a list of books filtered by the request.
func (service *bookService) GetFilteredBooks(request map[string]string) ([]types.ReadBookResponse, error) {
	// get filtered books
	bookDetails, err := service.bookRepo.GetFilteredBooks(request)
	if err != nil {
		return nil, err
	}
	if len(bookDetails) == 0 {
		return nil, errors.New("no book found")
	}

	// convert to responses type
	var responses []types.ReadBookResponse
	for _, val := range bookDetails {
		responses = append(responses, types.ReadBookResponse{
			ID:          val.ID,
			BookName:    val.BookName,
			AuthorID:    val.AuthorID,
			Publication: val.Publication,
		})
	}

	return responses, nil
}

// GetBook returns a book by the bookID.
func (service *bookService) GetBook(bookID uint) (*types.ReadBookResponse, error) {
	// get book from db
	bookDetail, err := service.bookRepo.GetBook(bookID)
	if err != nil {
		return nil, err
	}

	// convert to response type
	response := &types.ReadBookResponse{
		ID:          bookDetail.ID,
		BookName:    bookDetail.BookName,
		AuthorID:    bookDetail.AuthorID,
		Publication: bookDetail.Publication,
	}

	return response, nil
}

// CreateBook handles the create book request.
func (service *bookService) CreateBook(request *types.CreateBookRequest) error {
	// check if author exists
	if _, err := service.authorRepo.GetAuthor(request.AuthorID); err != nil {
		return errors.New("no author found with given author ID. Please create associated Author or give existing author ID")
	}

	// prepare book detail
	bookDetail := &models.BookDetail{
		BookName:    request.BookName,
		AuthorID:    request.AuthorID,
		Publication: request.Publication,
	}

	// create book in db
	if err := service.bookRepo.CreateBook(bookDetail); err != nil {
		return err
	}

	return nil
}

// UpdateBook handles the update book request.
func (service *bookService) UpdateBook(bookID uint, request *types.UpdateBookRequest) error {
	// check if book exists
	existingBook, err := service.bookRepo.GetBook(bookID)
	if err != nil {
		return errors.New("no book found with given ID")
	}

	// update existing book details
	if request.BookName != "" {
		existingBook.BookName = request.BookName
	}
	if request.AuthorID != 0 {
		if _, err := service.authorRepo.GetAuthor(request.AuthorID); err != nil {
			return errors.New("no author found with given author ID. Please create associated Author or give existing author ID")
		}
		existingBook.AuthorID = request.AuthorID
	}
	if request.Publication != "" {
		existingBook.Publication = request.Publication
	}

	// update book in db
	if err := service.bookRepo.UpdateBook(existingBook); err != nil {
		return errors.New("book was not updated")
	}

	return nil
}

// DeleteBook handles the delete book request.
func (service *bookService) DeleteBook(bookID uint) error {
	// check if book exists
	if _, err := service.bookRepo.GetBook(bookID); err != nil {
		return errors.New("no book found with given ID")
	}

	// delete book
	if err := service.bookRepo.DeleteBook(bookID); err != nil {
		return errors.New("book was not deleted")
	}

	return nil
}
