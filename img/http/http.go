package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RouteInit ...
func RouteInit() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}))

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("moviwiki"),
		TokenLookup: "query:Auth-Token",
	}))

	s := &http.Server{
		Addr:         ":3001",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.GET("/", hello)
	e.Logger.Fatal(e.StartServer(s))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
