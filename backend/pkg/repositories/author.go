package repositories

import (
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
)

// authorRepo defines the methods of the domain.IAuthorRepo interface.
type authorRepo struct {
	db *gorm.DB
}

// AuthorDBInstance returns a new instance of the authorRepo struct.
func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &authorRepo{
		db: d,
	}
}

// GetFilteredAuthors returns a list of authors filtered by the request.
func (repo *authorRepo) GetFilteredAuthors(request map[string]string) ([]models.AuthorDetail, error) {
	// get all authors
	var authorDetails []models.AuthorDetail
	if err := repo.db.Find(&authorDetails).Error; err != nil {
		return nil, err
	}

	// parse the schema
	parsedSchema, err := schema.Parse(&models.AuthorDetail{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return nil, err
	}

	// filter the authors for each field in the request
	for key, value := range request {
		mappedFieldInDB := parsedSchema.FieldsByName[key].DBName
		err := repo.db.Where(mappedFieldInDB+" = ?", value).Find(&authorDetails).Error
		if err != nil {
			return nil, err
		}
	}

	return authorDetails, nil
}

// GetAuthor returns an author by the authorID.
func (repo *authorRepo) GetAuthor(authorID uint) (*models.AuthorDetail, error) {
	authorDetail := &models.AuthorDetail{}
	if err := repo.db.Where("id = ?", authorID).First(authorDetail).Error; err != nil {
		return nil, err
	}
	return authorDetail, nil
}

// CreateAuthor creates a new author with given book details.
func (repo *authorRepo) CreateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}

// UpdateAuthor updates an author with given author details.
func (repo *authorRepo) UpdateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Save(author).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAuthor deletes an author with the given authorID.
func (repo *authorRepo) DeleteAuthor(authorID uint) error {
	authorDetail := &models.AuthorDetail{}
	if err := repo.db.Where("id = ?", authorID).Delete(authorDetail).Error; err != nil {
		return err
	}
	return nil
}
