package route

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/env"
	. "github.com/rizkysaputra4/moviwiki/server/route/handler"
)

// InitRoute initialize function
func InitRoute() {

	e := echo.New()
	e.POST("/add-user", RegisteringNewUser)
	e.PUT("/add-user", UpdateUserShortInfo)
	e.PUT("/my-profile", UpdateFullUserInfo)
	e.GET("/my-profile", GetMyProfile)
	e.POST("/insert-country", InsertCountry)
	e.POST("/check-email", CheckIfEmailExist)
	e.POST("/check-username", CheckIfUserNameExist)

	c := env.GetConfiguration()
	fmt.Println("running on port", c.ServerAPIPort)
	log.Fatal(e.Start(c.ServerAPIPort))
}
