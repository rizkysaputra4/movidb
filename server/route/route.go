package route

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/route/handler"
)

// Init initialize function
func Init() {

	e := echo.New()
	e.POST("/register", handler.RegisteringNewUser)
	e.POST("/insert-country", handler.InsertCountry)
	e.POST("/insert-short-user", handler.InsertShortUser)
	e.POST("/check-email", handler.CheckIfEmailExist)
	e.POST("/check-username", handler.CheckIfUserNameExist)

	fmt.Println("running on port 3000")
	log.Fatal(e.Start(":3000"))
}
