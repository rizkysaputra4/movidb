package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	h "github.com/rizkysaputra4/moviwiki/img/http/handler"
	mid "github.com/rizkysaputra4/moviwiki/img/http/middleware"
)

// RouteInit ...
func RouteInit() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4000", "http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	s := &http.Server{
		Addr:         ":3001",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.GET("/", hello)

	post := e.Group("/post")

	post.Use(mid.JWTAuth)

	post.GET("/post", hello)
	post.POST("/upload", h.UploadPP)

	e.Logger.Fatal(e.StartServer(s))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
