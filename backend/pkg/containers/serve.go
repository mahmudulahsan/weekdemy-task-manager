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
	bookRepo := repositories.BookDBInstance(db)
	authorRepo := repositories.AuthorDBInstance(db)
	userRepo := repositories.UserDBInstance(db)
	teamRepo := repositories.TeamDBInstance(db)

	// service initialization
	bookService := services.BookServiceInstance(bookRepo, authorRepo)
	authorService := services.AuthorServiceInstance(authorRepo, bookRepo)
	authService := services.AuthServiceInstance(userRepo)
	teamService := services.TeamServiceInstance(teamRepo)

	// controller initialization
	bookCtr := controllers.NewBookController(bookService)
	authorCtr := controllers.NewAuthorController(authorService)
	authCtr := controllers.NewAuthController(authService)
	teamCtr := controllers.NewTeamController(teamService)

	//route initialization
	bookRoutes := routes.NewBookRoutes(e, bookCtr)
	authorRoutes := routes.NewAuthorRoutes(e, authorCtr)
	authRoutes := routes.NewAuthRoutes(e, authCtr)
	demoRoutes := routes.NewDemoRoutes(e)
	teamRoutes := routes.NewTeamRoutes(e, teamCtr)

	//route binding
	bookRoutes.InitBookRoutes()
	authorRoutes.InitAuthorRoutes()
	authRoutes.InitAuthRoutes()
	demoRoutes.InitDemoRoutes()
	teamRoutes.InitTeamRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
