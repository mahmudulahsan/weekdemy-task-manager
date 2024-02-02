package types

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ReadTeamResponse struct {
	ID           uint   `json:"id"`
	TeamName     string `json:"teamName"`
	ProjectName  string `json:"projectName"`
	IsFinished   bool   `json:"isFinished"`
	StartTime    string `json:"startTime,omitempty"`
	FinishedTime string `json:"finishedTime,omitempty"`
}

type DeleteTeamResponse struct {
	MSG string `json:"msg"`
}

type CreateTeamRequest struct {
	ReadTeamResponse
	ID uint `json:"-"`
}

type UpdateTeamRequest struct {
	ReadTeamResponse
	ID uint `json:"-"`
}

func (request CreateTeamRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.TeamName,
			validation.Required.Error("Team name cannot be empty"),
			validation.Length(1, 256)),
		validation.Field(&request.ProjectName,
			validation.Required.Error("Project name cannot be empty"),
			validation.Length(1, 256)))
}

func (request UpdateTeamRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.TeamName,
			validation.Length(1, 256)),
		validation.Field(&request.ProjectName,
			validation.Length(1, 256)))
}
