package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
)

// InsertShortUser ...
func InsertShortUser(e echo.Context) error {
	fmt.Println("Inserting User")

	//var userShortInfo model.UserShortInfo
	userShortInfo := &model.UserShortInfo{
		
	}

	err := e.Bind(&userShortInfo)

	if err != nil {
		fmt.Println("err1 ---------------------")
		return err
	}

	_, err = db.DB.Model(userShortInfo).Insert()

	if err != nil {
		fmt.Println("---------------------------")
		panic(err)
	}
	
	return e.JSON(http.StatusOK, userShortInfo)
}

// InsertNewUser ...
func InsertNewUser(e echo.Context) error {
	fmt.Println("Inserting User")

	//var userShortInfo model.UserShortInfo
	userShortInfo := &model.UserInformation{
		
	}

	err := e.Bind(&userShortInfo)

	if err != nil {
		fmt.Println("err1 ---------------------")
		return err
	}

	_, err = db.DB.Model(userShortInfo).Insert()

	if err != nil {
		fmt.Println("---------------------------")
		panic(err)
	}
	
	return e.JSON(http.StatusOK, userShortInfo)
}