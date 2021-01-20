package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

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

// GetMyProfile ...
func GetMyProfile(w http.ResponseWriter, r *http.Request) {
	userProfile := &model.UserInformation{}
	if err := json.NewDecoder(r.Body).Decode(&userProfile); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when decode request body")
		return
	}

	if err := DB.Model(userProfile).Where("user_id = ?user_id").Select(); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when selecting user profile from db")
		return
	}

	// Updating user last request
	user := &model.UserShortInfo{UserID: userProfile.UserID}
	if err := user.UpdateLastRequest(); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when updating last request")
		return
	}

	comp.BasicResponse(w, http.StatusOK, "OK", userProfile)
}

// UpdateFullUserInfo ...
func UpdateFullUserInfo(w http.ResponseWriter, r *http.Request) {

	userProfile := &model.UserInformation{}
	if err := json.NewDecoder(r.Body).Decode(&userProfile); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, "Error when decode request body", err.Error())
		return
	}

	_, err := DB.Model(userProfile).
		Where("user_id = ?user_id").
		Column("user_full_name", "birthdate", "bio", "fb_link", "twitter_link", "ig_link", "sex").
		Update()
	if err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when updating user profile into db")
		return
	}

	// Updating user last request
	user := &model.UserShortInfo{UserID: userProfile.UserID}

	if err = user.UpdateLastRequest(); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when updating last request")
		return
	}

	comp.BasicResponse(w, http.StatusOK, "", userProfile)
}

// UpdateUserShortInfo ...
func UpdateUserShortInfo(w http.ResponseWriter, r *http.Request) {

	userShortInfo := &model.UserShortInfo{}

	if err := json.NewDecoder(r.Body).Decode(&userShortInfo); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when decode request body")
		return
	}

	if err := CheckUserName(userShortInfo.UserName); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userShortInfo.Password), bcrypt.DefaultCost)

	userShortInfo.Password = string(hashedPassword)
	userShortInfo.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = DB.Model(userShortInfo).
		Where("user_id = ?user_id").
		Column("user_name", "country_id", "password", "email", "last_request").
		Update()
	if err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when updating user credetials into db")
		return
	}

	comp.BasicResponse(w, http.StatusOK, "OK", userShortInfo)
}
