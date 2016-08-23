package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"

	"github.com/edwinlab/inventory/user"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", user.GetUsers)
	e.POST("/users", user.CreateUser)
	e.GET("/users/:id", user.GetUser)
	e.PUT("/users/:id", user.UpdateUser)
	e.DELETE("/users/:id", user.DeleteUser)

	// Start server
	e.Run(fasthttp.New(":1323"))
}