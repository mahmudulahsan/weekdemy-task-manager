package routes

import (
	"github.com/labstack/echo/v4"
	"weekdemy-task-manager-backend/pkg/controllers"
	"weekdemy-task-manager-backend/pkg/middlewares"
)

// BookRoutes stores controller and echo instance for book.
type BookRoutes struct {
	echo       *echo.Echo
	controller controllers.BookController
}

// NewBookRoutes returns a new instance of the BookRoutes struct.
func NewBookRoutes(echo *echo.Echo, controller controllers.BookController) *BookRoutes {
	return &BookRoutes{
		echo:       echo,
		controller: controller,
	}
}

// InitBookRoutes initializes the book routes.
func (bookRoutes *BookRoutes) InitBookRoutes() {
	e := bookRoutes.echo

	e.GET("/ping", Pong)

	book := e.Group("/bookstore")
	book.GET("/books", bookRoutes.controller.GetFilteredBooks)
	book.GET("/books/:id", bookRoutes.controller.GetBook)

	book.Use(middlewares.ValidateToken)

	book.POST("/books", bookRoutes.controller.CreateBook)
	book.PUT("/books/:id", bookRoutes.controller.UpdateBook)
	book.DELETE("/books/:id", bookRoutes.controller.DeleteBook)
}

// Pong checks if the server is running.
//func Pong(ctx echo.Context) error {
//	fmt.Println("Pong")
//	return ctx.JSON(http.StatusOK, "Pong")
//}
