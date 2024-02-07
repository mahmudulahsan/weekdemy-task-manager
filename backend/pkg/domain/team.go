package domain

import (
	"weekdemy-task-manager-backend/pkg/models"
	"weekdemy-task-manager-backend/pkg/types"
)

type ITeamRepo interface {
	GetFilteredTeams(request map[string]string) ([]models.TeamDetail, error)
	GetTeam(teamID uint) (*models.TeamDetail, error)
	CreateTeam(team *models.TeamDetail) (*models.TeamDetail, error)
	UpdateTeam(team *models.TeamDetail) (*models.TeamDetail, error)
	DeleteTeam(teamID uint) error
}

type ITeamService interface {
	GetFilteredTeams(request map[string]string) ([]types.ReadTeamResponse, error)
	GetTeam(teamID uint) (*types.ReadTeamResponse, error)
	CreateTeam(request *types.CreateTeamRequest) (*types.ReadTeamResponse, error)
	UpdateTeam(teamID uint, request *types.UpdateTeamRequest) (*types.ReadTeamResponse, error)
	DeleteTeam(teamID uint) (*types.DeleteTeamResponse, error)
}
