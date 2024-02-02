package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DemoRoutes struct {
	echo *echo.Echo
}

func NewDemoRoutes(echo *echo.Echo) *DemoRoutes {
	return &DemoRoutes{
		echo: echo,
	}
}

func (demoRoutes *DemoRoutes) InitDemoRoutes() {
	e := demoRoutes.echo

	e.GET("/ping", Pong)
	demo := e.Group("/weekdemy-demo")
	demo.GET("/teams", getDemoTasks)
	demo.GET("/teams/:id", getDemoTask)
	demo.POST("/teams", createDemoTask)
	demo.PUT("/teams/:id", updateDemoTask)
	demo.DELETE("/teams/:id", deleteDemoTask)

}

// Pong checks if the server is running.
//func Pong(ctx echo.Context) error {
//	fmt.Println("Pong")
//	return ctx.JSON(http.StatusOK, "Pong")
//}

type DemoTaskCreate struct {
	TeamName     string `json:"teamName"`
	ProjectName  string `json:"projectName"`
	IsFinished   bool   `json:"isFinished"`
	StartTime    string `json:"startTime"`
	FinishedTime string `json:"finishedTime"`
}

type DemoTaskReturn struct {
	ID           int    `json:"id"`
	TeamName     string `json:"teamName"`
	ProjectName  string `json:"projectName"`
	IsFinished   bool   `json:"isFinished"`
	StartTime    string `json:"startTime"`
	FinishedTime string `json:"finishedTime"`
}

func getDemoTasks(ctx echo.Context) error {
	// create 5 demo tasks
	tasks := []DemoTaskReturn{
		{1, "Team 1", "Project 1", false, "2024-01-22 22:11:06", "2024-01-27 22:11:06"},
		{2, "Team 2", "Project 2", false, "2024-01-23 22:12:06", "2024-01-26 22:12:06"},
		{3, "Team 3", "Project 3", false, "2024-01-24 22:13:06", "2024-01-25 22:13:06"},
		{4, "Team 4", "Project 4", false, "2024-01-25 22:14:06", "2024-01-24 22:14:06"},
	}
	return ctx.JSON(http.StatusOK, tasks)
}

func getDemoTask(ctx echo.Context) error {
	id := ctx.Param("id")
	parsedID, _ := strconv.Atoi(id)
	task := DemoTaskReturn{
		ID:           parsedID,
		TeamName:     "Team 1",
		ProjectName:  "Project 1",
		IsFinished:   false,
		StartTime:    "2024-01-28 22:11:06",
		FinishedTime: "2024-01-22 22:11:06",
	}
	return ctx.JSON(http.StatusOK, task)
}

func createDemoTask(ctx echo.Context) error {
	request := &DemoTaskCreate{}
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request body")
	}
	task := DemoTaskReturn{
		ID:           5,
		TeamName:     request.TeamName,
		ProjectName:  request.ProjectName,
		IsFinished:   request.IsFinished,
		StartTime:    request.StartTime,
		FinishedTime: request.FinishedTime,
	}
	return ctx.JSON(http.StatusCreated, task)
}

func updateDemoTask(ctx echo.Context) error {
	id := ctx.Param("id")
	parsedID, _ := strconv.Atoi(id)
	request := &DemoTaskCreate{}
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request body")
	}
	task := DemoTaskReturn{
		ID:           parsedID,
		TeamName:     request.TeamName,
		ProjectName:  request.ProjectName,
		IsFinished:   request.IsFinished,
		StartTime:    request.StartTime,
		FinishedTime: request.FinishedTime,
	}
	return ctx.JSON(http.StatusOK, task)
}

type MSG struct {
	Msg string `json:"msg"`
}

func deleteDemoTask(ctx echo.Context) error {
	id := ctx.Param("id")
	parsedID, _ := strconv.Atoi(id)
	return ctx.JSON(http.StatusOK, MSG{Msg: fmt.Sprintf("Task with ID %d has been deleted", parsedID)})
}
