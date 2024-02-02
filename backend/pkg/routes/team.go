package routes

import (
	"github.com/labstack/echo/v4"
	"weekdemy-task-manager-backend/pkg/controllers"
)

type TeamRoutes struct {
	echo       *echo.Echo
	controller controllers.TeamController
}

func NewTeamRoutes(echo *echo.Echo, controller controllers.TeamController) *TeamRoutes {
	return &TeamRoutes{
		echo:       echo,
		controller: controller,
	}
}

func (teamRoutes *TeamRoutes) InitTeamRoutes() {
	e := teamRoutes.echo

	e.GET("/ping", Pong)

	team := e.Group("/weekdemy")
	team.GET("/teams", teamRoutes.controller.GetFilteredTeams)
	team.GET("/teams/:teamID", teamRoutes.controller.GetTeam)
	team.POST("/teams", teamRoutes.controller.CreateTeam)
	team.PUT("/teams/:teamID", teamRoutes.controller.UpdateTeam)
	team.DELETE("/teams/:teamID", teamRoutes.controller.DeleteTeam)
}

// Pong checks if the server is running.
func Pong(ctx echo.Context) error {
	return ctx.JSON(200, "Pong")
}
