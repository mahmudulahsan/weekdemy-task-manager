package containers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"weekdemy-task-manager-backend/pkg/config"
	"weekdemy-task-manager-backend/pkg/connection"
	"weekdemy-task-manager-backend/pkg/controllers"
	"weekdemy-task-manager-backend/pkg/repositories"
	"weekdemy-task-manager-backend/pkg/routes"
	"weekdemy-task-manager-backend/pkg/services"
)

// Serve is used to initialize the server.
func Serve(e *echo.Echo) {
	//config initialization
	config.SetConfig()

	//database initializations
	db := connection.GetDB()

	// repository initialization
	teamRepo := repositories.TeamDBInstance(db)

	// service initialization
	teamService := services.TeamServiceInstance(teamRepo)

	// controller initialization
	teamCtr := controllers.NewTeamController(teamService)

	//route initialization
	demoRoutes := routes.NewDemoRoutes(e)
	teamRoutes := routes.NewTeamRoutes(e, teamCtr)

	//route binding
	demoRoutes.InitDemoRoutes()
	teamRoutes.InitTeamRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
