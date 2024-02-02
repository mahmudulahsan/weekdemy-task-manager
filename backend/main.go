package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	"weekdemy-task-manager-backend/pkg/containers"
)

func main() {
	t := time.Now()
	println(t.Format("2/17/2024 3:47:06 PM"))
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	containers.Serve(e)
}
