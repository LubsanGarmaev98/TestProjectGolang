package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(c echo.Context) error {
	daysLeft := e.s.DaysLeft()
	text := fmt.Sprintf("Новый год через %d дней!", daysLeft)

	return c.String(http.StatusOK, text)
}
