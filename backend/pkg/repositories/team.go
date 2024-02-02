package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/models"
)

// teamRepo defines the methods of the domain.ITeamRepo interface.
type teamRepo struct {
	db *gorm.DB
}

// TeamDBInstance returns a new instance of the teamRepo struct.
func TeamDBInstance(d *gorm.DB) domain.ITeamRepo {
	return &teamRepo{
		db: d,
	}
}

// GetFilteredTeams returns a list of teams filtered by the request.
func (repo *teamRepo) GetFilteredTeams(request map[string]string) ([]models.TeamDetail, error) {
	var teamDetails []models.TeamDetail
	// get all teams
	if err := repo.db.Find(&teamDetails).Error; err != nil {
		return nil, err
	}

	// parse schema
	parsedSchema, err := schema.Parse(&models.TeamDetail{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return nil, err
	}

	// filter the teams for each field in the request
	for key, value := range request {
		mappedFieldInDB := parsedSchema.FieldsByName[key].DBName
		err = repo.db.Where(mappedFieldInDB+" = ?", value).Find(&teamDetails).Error
		if err != nil {
			return nil, err
		}
	}

	return teamDetails, nil
}

// GetTeam returns a team by the teamID.
func (repo *teamRepo) GetTeam(teamID uint) (*models.TeamDetail, error) {
	teamDetail := &models.TeamDetail{}
	if err := repo.db.Where("id = ?", teamID).First(teamDetail).Error; err != nil {
		return nil, err
	}
	return teamDetail, nil
}

// CreateTeam creates a new team with given team details and returns the created team.
func (repo *teamRepo) CreateTeam(team *models.TeamDetail) (*models.TeamDetail, error) {
	if err := repo.db.Create(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

// UpdateTeam updates a team with given team details and returns the updated team.
func (repo *teamRepo) UpdateTeam(team *models.TeamDetail) (*models.TeamDetail, error) {
	if err := repo.db.Save(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

// DeleteTeam deletes a team with the given teamID
func (repo *teamRepo) DeleteTeam(teamID uint) error {
	teamDetail := &models.TeamDetail{}
	if err := repo.db.Where("id = ?", teamID).Delete(teamDetail).Error; err != nil {
		return err
	}
	return nil
}
