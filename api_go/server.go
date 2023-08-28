package demo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HelloAPI struct {
	e *echo.Echo
}

func NewServer() *HelloAPI {
	return &HelloAPI{e: echo.New()}
}

func (hello *HelloAPI) Start(port string) {
	// Middleware
	hello.e.Use(middleware.Logger())
	hello.e.Use(middleware.Recover())

	// Database
	connection := CreateDbConnection()

	// Handler
	hello.e.GET("/hello", SayHi(connection))

	// Start server
	hello.e.Logger.Fatal(hello.e.Start(port))
}
