package app

import (
	"fmt"
	"github.com/LubsanGarmaev98/TestProjectGolang/internal/app/endpoint"
	"github.com/LubsanGarmaev98/TestProjectGolang/internal/app/middleware"
	"github.com/LubsanGarmaev98/TestProjectGolang/internal/app/service"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.GET("/check", a.e.Status, middleware.CheckRole)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Сервер запущен ...")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
