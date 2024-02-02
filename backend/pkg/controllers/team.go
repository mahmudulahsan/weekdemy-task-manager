package controllers

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"weekdemy-task-manager-backend/pkg/domain"
	"weekdemy-task-manager-backend/pkg/types"
)

type ITeamController interface {
	GetFilteredTeams(e echo.Context) error
	GetTeam(e echo.Context) error
	CreateTeam(e echo.Context) error
	UpdateTeam(e echo.Context) error
	DeleteTeam(e echo.Context) error
}

type TeamController struct {
	teamSvc domain.ITeamService
}

func NewTeamController(teamSvc domain.ITeamService) TeamController {
	return TeamController{
		teamSvc: teamSvc,
	}
}

func (controller *TeamController) GetFilteredTeams(e echo.Context) error {
	// prepare the request
	request := make(map[string]string)
	if e.QueryParam("teamName") != "" {
		request["TeamName"] = e.QueryParam("teamName")
	}
	if e.QueryParam("projectName") != "" {
		request["ProjectName"] = e.QueryParam("projectName")
	}
	if e.QueryParam("isFinished") != "" {
		request["IsFinished"] = e.QueryParam("isFinished")
	}
	if e.QueryParam("startTime") != "" {
		request["StartTime"] = e.QueryParam("startTime")
	}
	if e.QueryParam("finishedTime") != "" {
		request["FinishedTime"] = e.QueryParam("finishedTime")
	}

	// pass the request to the service layer
	res, err := controller.teamSvc.GetFilteredTeams(request)
	if err != nil {
		return e.JSON(500, err.Error())
	}

	return e.JSON(200, res)
}

func (controller *TeamController) GetTeam(e echo.Context) error {
	// get teamID from the request
	teamID, err := strconv.Atoi(e.Param("teamID"))
	if err != nil {
		return e.JSON(400, "invalid teamID")
	}

	// pass the request to the service layer
	res, err := controller.teamSvc.GetTeam(uint(teamID))
	if err != nil {
		return e.JSON(500, err.Error())
	}

	return e.JSON(200, res)
}

func (controller *TeamController) CreateTeam(e echo.Context) error {
	// prepare the request
	request := &types.CreateTeamRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(400, "invalid request body")
	}

	// validate the request
	if err := request.Validate(); err != nil {
		return e.JSON(400, err.Error())
	}

	// pass the request to the service layer
	response, err := controller.teamSvc.CreateTeam(request)
	if err != nil {
		return e.JSON(500, err.Error())
	}

	return e.JSON(201, response)
}

func (controller *TeamController) UpdateTeam(e echo.Context) error {
	// get teamID from the request
	teamID, err := strconv.Atoi(e.Param("teamID"))
	if err != nil {
		return e.JSON(400, "invalid teamID")
	}

	// prepare the request
	request := &types.UpdateTeamRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(400, "invalid request body")
	}

	// validate the request
	if err := request.Validate(); err != nil {
		return e.JSON(400, err.Error())
	}

	// pass the request to the service layer
	response, err := controller.teamSvc.UpdateTeam(uint(teamID), request)
	if err != nil {
		return e.JSON(500, err.Error())
	}

	return e.JSON(200, response)
}

func (controller *TeamController) DeleteTeam(e echo.Context) error {
	teamID, err := strconv.Atoi(e.Param("teamID"))
	if err != nil {
		return e.JSON(400, "invalid teamID")
	}

	// pass the request to the service layer
	response, err := controller.teamSvc.DeleteTeam(uint(teamID))
	if err != nil {
		return e.JSON(500, err.Error())
	}
	return e.JSON(200, response)
}
