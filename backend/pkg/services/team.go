package services

import (
	"errors"
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/models"
	"weekdemy-task-manager-backend/pkg/types"
)

// teamService defines the methods of the domain.ITeamService interface.
type teamService struct {
	repo domain.ITeamRepo
}

// TeamServiceInstance returns a new instance of the teamService struct.
func TeamServiceInstance(repo domain.ITeamRepo) domain.ITeamService {
	return &teamService{
		repo: repo,
	}
}

// GetFilteredTeams returns a list of teams filtered by the request.
func (service *teamService) GetFilteredTeams(request map[string]string) ([]types.ReadTeamResponse, error) {
	// get filtered teams
	teamDetails, err := service.repo.GetFilteredTeams(request)
	if err != nil {
		return nil, err
	}

	// convert to response type
	var responses []types.ReadTeamResponse
	for _, val := range teamDetails {
		responses = append(responses, types.ReadTeamResponse{
			ID:           val.ID,
			TeamName:     val.TeamName,
			ProjectName:  val.ProjectName,
			IsFinished:   val.IsFinished,
			StartTime:    val.StartTime,
			FinishedTime: val.FinishedTime,
		})
	}

	if len(responses) == 0 {
		return nil, errors.New("no team found")
	}
	return responses, nil
}

// GetTeam returns a team by the teamID.
func (service *teamService) GetTeam(teamID uint) (*types.ReadTeamResponse, error) {
	// get team from db
	teamDetail, err := service.repo.GetTeam(teamID)
	if err != nil {
		return nil, err
	}

	// convert to response type
	response := &types.ReadTeamResponse{
		ID:           teamDetail.ID,
		TeamName:     teamDetail.TeamName,
		ProjectName:  teamDetail.ProjectName,
		IsFinished:   teamDetail.IsFinished,
		StartTime:    teamDetail.StartTime,
		FinishedTime: teamDetail.FinishedTime,
	}

	return response, nil
}

// CreateTeam creates a new team with given team details and returns the created team.
func (service *teamService) CreateTeam(request *types.CreateTeamRequest) (*types.ReadTeamResponse, error) {
	// prepare team detail
	teamDetail := &models.TeamDetail{
		TeamName:     request.TeamName,
		ProjectName:  request.ProjectName,
		IsFinished:   request.IsFinished,
		StartTime:    request.StartTime,
		FinishedTime: request.FinishedTime,
	}

	// create team in db
	createdTeam, err := service.repo.CreateTeam(teamDetail)
	if err != nil {
		return nil, err
	}

	// convert to response type
	response := &types.ReadTeamResponse{
		ID:           createdTeam.ID,
		TeamName:     createdTeam.TeamName,
		ProjectName:  createdTeam.ProjectName,
		IsFinished:   createdTeam.IsFinished,
		StartTime:    createdTeam.StartTime,
		FinishedTime: createdTeam.FinishedTime,
	}

	return response, nil
}

func (service *teamService) UpdateTeam(teamID uint, request *types.UpdateTeamRequest) (*types.ReadTeamResponse, error) {
	// check if team exists
	existingTeam, err := service.repo.GetTeam(teamID)
	if err != nil {
		return nil, errors.New("no team found with given ID")
	}

	// update existing team details
	existingTeam.TeamName = request.TeamName
	existingTeam.ProjectName = request.ProjectName
	existingTeam.IsFinished = request.IsFinished
	existingTeam.StartTime = existingTeam.StartTime
	existingTeam.FinishedTime = existingTeam.FinishedTime

	// update team in db
	updatedTeam, err := service.repo.UpdateTeam(existingTeam)
	if err != nil {
		return nil, err
	}

	// convert to response type
	response := &types.ReadTeamResponse{
		ID:           updatedTeam.ID,
		TeamName:     updatedTeam.TeamName,
		ProjectName:  updatedTeam.ProjectName,
		IsFinished:   updatedTeam.IsFinished,
		StartTime:    updatedTeam.StartTime,
		FinishedTime: updatedTeam.FinishedTime,
	}

	return response, nil
}

func (service *teamService) DeleteTeam(teamID uint) (*types.DeleteTeamResponse, error) {
	// delete team from db
	err := service.repo.DeleteTeam(teamID)
	if err != nil {
		return nil, err
	}

	// prepare response
	response := &types.DeleteTeamResponse{
		MSG: "Team with ID " + string(teamID) + " deleted successfully.",
	}

	return response, nil
}
