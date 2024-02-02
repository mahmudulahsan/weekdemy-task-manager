package services

import (
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/models"
	"weekdemy-task-manager-backend/pkg/types"
	"errors"
)

// AuthorService defines the methods of the domain.IAuthorService interface.
type authorService struct {
	authorRepo domain.IAuthorRepo
	bookRepo   domain.IBookRepo
}

// AuthorServiceInstance returns a new instance of the AuthorService struct.
func AuthorServiceInstance(authorRepo domain.IAuthorRepo, bookRepo domain.IBookRepo) domain.IAuthorService {
	return &authorService{
		authorRepo: authorRepo,
		bookRepo:   bookRepo,
	}
}

// GetFilteredAuthors returns a list of authors filtered by the request.
func (service *authorService) GetFilteredAuthors(request map[string]string) ([]types.ReadAuthorResponse, error) {
	// get filtered authors
	authorDetails, err := service.authorRepo.GetFilteredAuthors(request)
	if err != nil {
		return nil, err
	}
	if len(authorDetails) == 0 {
		return nil, errors.New("no author found with given query")
	}

	// convert to response type
	var responses []types.ReadAuthorResponse
	for _, val := range authorDetails {
		responses = append(responses, types.ReadAuthorResponse{
			ID:          val.ID,
			AuthorName:  val.AuthorName,
			Address:     val.Address,
			PhoneNumber: val.PhoneNumber,
		})
	}

	return responses, nil
}

// GetAuthor returns an author by the authorID.
func (service *authorService) GetAuthor(authorID uint) (*types.ReadAuthorResponse, error) {
	// get author from db
	authorDetail, err := service.authorRepo.GetAuthor(authorID)
	if err != nil {
		return nil, err
	}

	// convert to response type
	response := &types.ReadAuthorResponse{
		ID:          authorDetail.ID,
		AuthorName:  authorDetail.AuthorName,
		Address:     authorDetail.Address,
		PhoneNumber: authorDetail.PhoneNumber,
	}

	return response, nil
}

// CreateAuthor handles the create author request.
func (service *authorService) CreateAuthor(request *types.CreateAuthorRequest) error {
	// prepare author detail
	authorDetail := &models.AuthorDetail{
		AuthorName:  request.AuthorName,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}

	// create author in db
	if err := service.authorRepo.CreateAuthor(authorDetail); err != nil {
		return err
	}

	return nil
}

// UpdateAuthor handles the update author request.
func (service *authorService) UpdateAuthor(authorID uint, request *types.UpdateAuthorRequest) error {
	// check if author exists
	existingAuthor, err := service.authorRepo.GetAuthor(authorID)
	if err != nil {
		return errors.New("no author found with given author ID")
	}

	// update existing author details
	if request.AuthorName != "" {
		existingAuthor.AuthorName = request.AuthorName
	}
	if request.Address != "" {
		existingAuthor.Address = request.Address
	}
	if request.PhoneNumber != "" {
		existingAuthor.PhoneNumber = request.PhoneNumber
	}

	// update author in db
	if err := service.authorRepo.UpdateAuthor(existingAuthor); err != nil {
		return errors.New("author was not updated")
	}

	return nil
}

// DeleteAuthor handles the delete author request.
func (service *authorService) DeleteAuthor(authorID uint) error {
	// check if author exists
	if _, err := service.authorRepo.GetAuthor(authorID); err != nil {
		return errors.New("no author found with given author ID")
	}

	// delete author and books of author
	if err := service.authorRepo.DeleteAuthor(authorID); err != nil {
		return err
	}
	if err := service.bookRepo.DeleteBooksByAuthorID(authorID); err != nil {
		return err
	}

	return nil
}
