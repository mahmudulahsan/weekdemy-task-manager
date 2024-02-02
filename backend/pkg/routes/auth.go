package routes

import (
	"weekdemy-task-manager-backend/pkg/controllers"
	"github.com/labstack/echo/v4"
)

// AuthRoutes stores controller and echo instance for authentication.
type AuthRoutes struct {
	echo    *echo.Echo
	authCtr controllers.AuthController
}

// NewAuthRoutes returns a new instance of the AuthRoutes struct.
func NewAuthRoutes(echo *echo.Echo, authCtr controllers.AuthController) *AuthRoutes {
	return &AuthRoutes{
		echo:    echo,
		authCtr: authCtr,
	}
}

// InitAuthRoutes initializes the authentication routes.
func (routes *AuthRoutes) InitAuthRoutes() {
	e := routes.echo

	e.POST("/login", routes.authCtr.Login)
	e.POST("/signup", routes.authCtr.Signup)
}
