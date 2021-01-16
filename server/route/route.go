package route

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/env"
	. "github.com/rizkysaputra4/moviwiki/server/route/handler"
)

// InitRoute initialize function
func InitRoute() {
	c := env.GetConfiguration()
	e := echo.New()
	router := chi.NewRouter()

	router.Get("/check-email", CheckIfEmailExist)

	http.ListenAndServe(c.ServerAPIPort, router)


	e.POST("/register", RegisteringNewUser)
	e.PUT("/register", UpdateUserShortInfo)
	e.GET("/login", CheckIfUserExist)
	e.GET("/login-password", CheckIfPasswordMatch)
	e.PUT("/my-profile", UpdateFullUserInfo)
	e.GET("/my-profile", GetMyProfile)
	e.POST("/insert-country", InsertCountry)
	// e.GET("/check-email", CheckIfEmailExist)
	e.GET("/check-username", CheckIfUserNameExist)

	// e.GET("/GetSession", session.SetSessionValue)

	
	// fmt.Println("running on port", c.ServerAPIPort)
	// log.Fatal(e.Start(c.ServerAPIPort))
}
