package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/comp"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
)

// InsertShortUser ...
func InsertShortUser(e echo.Context) error {
	fmt.Println("Inserting User")

	//var userShortInfo model.UserShortInfo
	userShortInfo := &model.UserShortInfo{}

	err := e.Bind(&userShortInfo)

	if err != nil {
		return err
	}

	_, err = db.DB.Model(userShortInfo).Insert()

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, userShortInfo)
}

// InsertNewUser is inserting full user information
func InsertNewUser(e echo.Context) error {
	fmt.Println("Inserting User")

	//var userShortInfo model.UserShortInfo
	userShortInfo := &model.UserInformation{}

	err := e.Bind(&userShortInfo)
	if err != nil {
		return err
	}

	_, err = db.DB.Model(userShortInfo).Insert()
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, userShortInfo)
}

// RegisteringNewUser is handling register request
func RegisteringNewUser(e echo.Context) error {

	registerNewUserData := &model.UserShortInfo{}

	response := &model.StatusResponse{}
	err := e.Bind(&registerNewUserData)
	if err != nil {
		response.Message = err.Error()
		response.Status = false
		return e.JSON(http.StatusBadGateway, response)
	}

	registerNewUserData.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = db.DB.Model(registerNewUserData).Insert()
	if err != nil {
		response.Message = err.Error()
		response.Status = false
		return e.JSON(http.StatusBadRequest, response)
	}

	return e.JSON(http.StatusOK, registerNewUserData)
}

// CheckIfEmailExist ...
func CheckIfEmailExist(e echo.Context) error {
	registerNewUserData := &model.UserShortInfo{}
	response := &model.StatusResponse{
		Status:  true,
		Message: "Ok",
	}

	err := e.Bind(&registerNewUserData)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		return e.JSON(http.StatusBadRequest, response)
	}

	column := "email"
	isExist := comp.CheckIfExist(column, registerNewUserData.Email, registerNewUserData)

	if isExist {
		response.Status = false
		response.Message = "This email has been registered"
		return e.JSON(http.StatusBadRequest, response)
	}

	return e.JSON(http.StatusOK, response)
}

// CheckIfUserNameExist ...
func CheckIfUserNameExist(e echo.Context) error {
	registerNewUserData := &model.UserShortInfo{}
	response := &model.StatusResponse{
		Message: "Ok",
	}

	err := e.Bind(&registerNewUserData)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	column := "user_name"
	isExist := comp.CheckIfExist(column, registerNewUserData.UserName, registerNewUserData)

	if isExist {
		response.Status = false
		response.Message = "Username has already taken"
		return e.JSON(http.StatusBadRequest, response)
	}
	return e.JSON(http.StatusOK, response)
}
