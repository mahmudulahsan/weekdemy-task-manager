package routes

import (
	"weekdemy-task-manager-backend/pkg/controllers"
	"weekdemy-task-manager-backend/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

// AuthorRoutes stores controller and echo instance for author.
type AuthorRoutes struct {
	echo       *echo.Echo
	controller controllers.AuthorController
}

// NewAuthorRoutes returns a new instance of the AuthorRoutes struct.
func NewAuthorRoutes(echo *echo.Echo, controller controllers.AuthorController) *AuthorRoutes {
	return &AuthorRoutes{
		echo:       echo,
		controller: controller,
	}
}

// InitAuthorRoutes initializes the author routes.
func (authorRoutes *AuthorRoutes) InitAuthorRoutes() {
	e := authorRoutes.echo

	author := e.Group("/bookstore")
	author.GET("/authors", authorRoutes.controller.GetFilteredAuthors)
	author.GET("/authors/:id", authorRoutes.controller.GetAuthor)

	author.Use(middlewares.ValidateToken)

	author.POST("/authors", authorRoutes.controller.CreateAuthor)
	author.PUT("/authors/:id", authorRoutes.controller.UpdateAuthor)
	author.DELETE("/authors/:id", authorRoutes.controller.DeleteAuthor)
}
