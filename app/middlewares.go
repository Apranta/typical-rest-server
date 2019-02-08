package app

import (
	"github.com/labstack/echo/middleware"
)

func (s *server) initMiddlewares() {
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())
	// check list of middleware at https://echo.labstack.com/middleware
}

// Put custom middleware belows
// Example: https://echo.labstack.com/cookbook/middleware
