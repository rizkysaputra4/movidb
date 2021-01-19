package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rizkysaputra4/moviwiki/server/comp"
	. "github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
	"github.com/rizkysaputra4/moviwiki/server/route/middleware"
	"golang.org/x/crypto/bcrypt"
)

// CheckIfEmailExist ...
func CheckIfEmailExist(w http.ResponseWriter, r *http.Request) {
	shortInfo := &model.UserShortInfo{}

	if err := json.NewDecoder(r.Body).Decode(&shortInfo); err != nil || shortInfo.Email == "" {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when decode request body")
		return
	}

	column := "email"
	if isExist := comp.CheckIfExist(column, shortInfo.Email, shortInfo); isExist {
		comp.BasicResponse(w, http.StatusBadRequest, "Email is already exist", "")
		return
	}

	comp.ResJSON(w, http.StatusOK, shortInfo)
}

// CheckIfUserNameExist ...
func CheckIfUserNameExist(w http.ResponseWriter, r *http.Request) {
	shortInfo := &model.UserShortInfo{}

	if err := json.NewDecoder(r.Body).Decode(&shortInfo); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when decode request body")
		return
	}

	if err := CheckUserName(shortInfo.UserName); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	column := "user_name"
	if isExist := comp.CheckIfExist(column, shortInfo.UserName, shortInfo); isExist {
		comp.BasicResponse(w, http.StatusBadRequest, "UserName is already exist", "")
		return
	}

	comp.ResJSON(w, http.StatusOK, shortInfo)
}

// CheckIfUserExist ...
func CheckIfUserExist(w http.ResponseWriter, r *http.Request) {
	info := &model.UserShortInfo{}
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when decode request body")
		return
	}

	err := DB.Model(info).
		Column("user_id", "user_name", "email").
		Where("user_name = ?user_name").
		WhereOr("email = ?email").
		Select()

	if err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "User does not exist")
		return
	}

	comp.BasicResponse(w, http.StatusOK, "", "User Exist")
}

// Login ...
type Login struct {
	Email    string `pg:"email" json:"email"`
	Username string `pg:"user_name" json:"user_name"`
	Pw       string `pg:"password" json:"password"`
}

// CheckIfPasswordMatch ...
func CheckIfPasswordMatch(w http.ResponseWriter, r *http.Request) {
	pw := &Login{}
	userInfo := &model.UserShortInfo{}
	if err := json.NewDecoder(r.Body).Decode(&pw); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when decode request body")
		return
	}
	err := DB.Model(userInfo).
		Where("user_name = ?", pw.Username).
		WhereOr("email = ?", pw.Email).
		Column("user_id", "user_name", "email", "password", "role").
		Select()

	if err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when selecting user cred from db")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(pw.Pw)); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when bcrypting user password")
		return
	}

	middleware.StoreJWT(w, r, userInfo.UserID, userInfo.Role)

	if userInfo.Role < 11 {
		middleware.StoreSession(w, r, userInfo.UserID, userInfo.Role)
	}

	userInfo.UpdateLastRequest()

	comp.BasicResponse(w, http.StatusOK, "", userInfo)
}

// LogOut ...
func LogOut(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteSession(w, r)
	middleware.DeleteJWTFromCookie(w, r)
	comp.BasicResponse(w, http.StatusOK, "", "logout success")
}

// RegisteringNewUser is handling register request
func RegisteringNewUser(w http.ResponseWriter, r *http.Request) {

	shortInfo := &model.UserShortInfo{}
	if err := json.NewDecoder(r.Body).Decode(&shortInfo); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when decode request body")
		return
	}

	if shortInfo.Password == "" || len(shortInfo.Password) < 6 {
		comp.BasicResponse(w, http.StatusBadRequest, "Password too Short, must be 6 character long", "")
		return
	}

	if err := CheckUserName(shortInfo.UserName); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(shortInfo.Password), bcrypt.DefaultCost)

	shortInfo.Password = string(hashedPassword)
	shortInfo.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")
	if err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when bcrypting password")
		return
	}

	_, err = DB.Model(shortInfo).
		Column("user_name", "country_id", "password", "email", "last_request").
		Insert()

	if err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when inserting user credential into db")
		return
	}

	completeUserData := &model.UserInformation{
		UserID:       shortInfo.UserID,
		RegisterDate: time.Now().UTC().Format("2006-01-02"),
	}

	_, err = DB.Model(completeUserData).Insert()
	if err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when inserting user profile into db")
		return
	}

	comp.ResJSON(w, http.StatusOK, shortInfo)
}
