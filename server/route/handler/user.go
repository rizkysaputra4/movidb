package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/comp"
	. "github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
)

// RegisteringNewUser is handling register request
func RegisteringNewUser(e echo.Context) error {

	shortInfo := &model.UserShortInfo{}

	err := e.Bind(&shortInfo)
	if err != nil {
		return e.JSON(http.StatusBadGateway, comp.BasicResponse(false, err.Error()))
	}

	shortInfo.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = DB.Model(shortInfo).Insert()
	if err != nil {
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, err.Error()),
		)
	}

	completeUserData := &model.UserInformation{
		UserID:       shortInfo.UserID,
		RegisterDate: time.Now().UTC().Format("2006-01-02"),
	}

	_, err = DB.Model(completeUserData).Insert()
	if err != nil {
		return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, "complete userdata error"))
	}

	return e.JSON(http.StatusOK, shortInfo)
}

// CheckIfEmailExist ...
func CheckIfEmailExist(e echo.Context) error {
	shortInfo := &model.UserShortInfo{}

	err := e.Bind(&shortInfo)
	if err != nil {
		return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, err.Error()))
	}

	column := "email"
	isExist := comp.CheckIfExist(column, shortInfo.Email, shortInfo)

	if isExist {
		return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, "This email has been registered"))
	}

	return e.JSON(http.StatusOK, comp.BasicResponse(false, "Ok"))
}

// CheckIfUserNameExist ...
func CheckIfUserNameExist(e echo.Context) error {
	shortInfo := &model.UserShortInfo{}

	err := e.Bind(&shortInfo)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	column := "user_name"
	isExist := comp.CheckIfExist(column, shortInfo.UserName, shortInfo)

	if isExist {
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, "Username has already taken"),
		)
	}

	return e.JSON(
		http.StatusOK, comp.BasicResponse(true, "Ok"),
	)
}

// GetMyProfile ...
func GetMyProfile(e echo.Context) error {
	userProfile := &model.UserInformation{}

	err := e.Bind(&userProfile)
	if err != nil {
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, "fail binding request to variable"),
		)
	}

	err = DB.Model(userProfile).Where("user_id = ?user_id").Select()
	if err != nil {
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, "fail get user info from database"),
		)
	}

	// Updating user last request
	user := &model.UserShortInfo{UserID: userProfile.UserID}
	err = user.UpdateLastRequest()
	if err != nil {
		fmt.Println(err)
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, "fail to update last request"),
		)
	}

	return e.JSON(http.StatusOK, userProfile)
}

// UpdateFullUserInfo ...
func UpdateFullUserInfo(e echo.Context) error {

	userProfile := &model.UserInformation{}

	err := e.Bind(&userProfile)
	if err != nil {
		return err
	}

	_, err = DB.Model(userProfile).
		Where("user_id = ?user_id").
		Column("user_full_name", "birthdate", "bio", "fb_link", "twitter_link", "ig_link", "sex", "last_request").
		Update()
	if err != nil {
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, err.Error()),
		)
	}

	// Updating user last request
	user := &model.UserShortInfo{UserID: userProfile.UserID}
	err = user.UpdateLastRequest()
	if err != nil {
		fmt.Println(err)
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, "fail to update last request"),
		)
	}

	return e.JSON(http.StatusOK, userProfile)
}

// UpdateUserShortInfo ...
func UpdateUserShortInfo(e echo.Context) error {

	userShortInfo := &model.UserShortInfo{}

	err := e.Bind(&userShortInfo)
	if err != nil {
		return err
	}

	userShortInfo.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = DB.Model(userShortInfo).
		Where("user_id = ?user_id").
		Column("user_name", "country_id", "password", "email", "last_request").
		Update()
	if err != nil {
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, err.Error()),
		)
	}

	return e.JSON(http.StatusOK, userShortInfo)
}
