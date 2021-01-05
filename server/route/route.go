package route

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/route/handler"
)

// Init initialize function
func Init(){

	e := echo.New()
	e.POST("/u", handler.InsertNewUser)
	e.POST("/insert-country", handler.InsertCountry)
	e.POST("/insert-user", handler.InsertShortUser)
	
	fmt.Println("running on port 3001")
	log.Fatal(e.Start(":3000"))
}