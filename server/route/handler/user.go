package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo"
	"github.com/rizkysaputra4/moviwiki/server/comp"
	. "github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
	"golang.org/x/crypto/bcrypt"
)

// CheckUserName check whether username is not contains stupid characters
func CheckUserName(s string) error {
	char := `[^a-zA-Z0-9._-]`
	re := regexp.MustCompile(char)
	matched := re.FindAllString(s, -1)
	if matched != nil || len(s) > 20 {
		return fmt.Errorf("Errpr: contain %v or have length more than 20 characters", char)
	}
	return nil
}

// RegisteringNewUser is handling register request
func RegisteringNewUser(e echo.Context) error {

	shortInfo := &model.UserShortInfo{}

	err := e.Bind(&shortInfo)
	if err != nil {
		return e.JSON(http.StatusBadGateway, comp.BasicResponse(false, err.Error()))
	}

	err = CheckUserName(shortInfo.UserName)
	if err != nil {
		return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, err.Error()))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(shortInfo.Password), bcrypt.DefaultCost)

	shortInfo.Password = string(hashedPassword)
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
func CheckIfEmailExist(w http.ResponseWriter, r *http.Request) {
	shortInfo := &model.UserShortInfo{}

	if err := json.NewDecoder(r.Body).Decode(&shortInfo); err != nil {
		http.Error(w, http.StatusText(422), 422) 
		fmt.Println(shortInfo)
		return
	}
	comp.ResJSON(w, http.StatusOK, shortInfo)

	// err := e.Bind(&shortInfo)
	// if err != nil {
	// 	return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, err.Error()))
	// }

	// column := "email"
	// isExist := comp.CheckIfExist(column, shortInfo.Email, shortInfo)

	// if isExist {
	// 	return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, "This email has been registered"))
	// }

	// return e.JSON(http.StatusOK, comp.BasicResponse(false, "Ok"))
}

// CheckIfUserNameExist ...
func CheckIfUserNameExist(e echo.Context) error {
	shortInfo := &model.UserShortInfo{}

	err := e.Bind(&shortInfo)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	err = CheckUserName(shortInfo.UserName)
	if err != nil {
		return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, err.Error()))
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
		Column("user_full_name", "birthdate", "bio", "fb_link", "twitter_link", "ig_link", "sex").
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

	err = CheckUserName(userShortInfo.UserName)
	if err != nil {
		return e.JSON(http.StatusBadRequest, comp.BasicResponse(false, err.Error()))
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

// CheckIfUserExist ...
func CheckIfUserExist(e echo.Context) error {
	info := &model.UserShortInfo{}

	err := e.Bind(&info)
	if err != nil {
		return err
	}

	err = DB.Model(info).
		Column("user_id", "user_name", "email").
		Where("user_name = ?user_name").
		WhereOr("email = ?email").
		Select()

	if err != nil {
		return e.JSON(
			http.StatusBadRequest, comp.BasicResponse(false, err.Error()),
		)
	}

	return e.JSON(http.StatusOK, comp.BasicResponse(true, "OK"))
}

// Login ...
type Login struct {
	Email    string `pg:"email" json:"email"`
	Username string `pg:"user_name" json:"user_name"`
	Pw       string `pg:"password" json:"password"`
}

// CheckIfPasswordMatch ...
func CheckIfPasswordMatch(e echo.Context) error {
	pw := &Login{}
	userInfo := &model.UserShortInfo{}
	err := e.Bind(&pw)
	if err != nil {
		return err
	}
	err = DB.Model(userInfo).
		Column("user_id", "user_name", "email", "password").
		Where("user_name = ?", pw.Username).
		WhereOr("email = ?", pw.Email).
		Select()
	if err != nil {
		return e.JSON(http.StatusBadRequest, "UserName/Email does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(pw.Pw))
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Password does not match")
	}

	user, _, _ := e.Request().BasicAuth()
	method := e.Request().Method
	path := e.Request().URL.Path

	fmt.Println(user, method, path)

	userInfo.UpdateLastRequest()
	return e.JSON(http.StatusOK, userInfo)
}
