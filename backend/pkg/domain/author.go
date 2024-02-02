package domain

import (
	"weekdemy-task-manager-backend/pkg/models"
	"weekdemy-task-manager-backend/pkg/types"
)

type IAuthorRepo interface {
	GetFilteredAuthors(request map[string]string) ([]models.AuthorDetail, error)
	GetAuthor(authorID uint) (*models.AuthorDetail, error)
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(author *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}

type IAuthorService interface {
	GetFilteredAuthors(request map[string]string) ([]types.ReadAuthorResponse, error)
	GetAuthor(authorID uint) (*types.ReadAuthorResponse, error)
	CreateAuthor(request *types.CreateAuthorRequest) error
	UpdateAuthor(authorID uint, request *types.UpdateAuthorRequest) error
	DeleteAuthor(authorID uint) error
}
