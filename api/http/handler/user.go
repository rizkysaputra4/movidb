package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/go-chi/chi"
	c "github.com/rizkysaputra4/moviwiki/api/context"
	"github.com/rizkysaputra4/moviwiki/api/db"
	"github.com/rizkysaputra4/moviwiki/api/http/middleware"
	"github.com/rizkysaputra4/moviwiki/api/model"
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
	c := &c.Context{Res: w, Req: r, Data: userProfile}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	if err := db.DB.Model(userProfile).Where("user_id = ?user_id").Select(); err != nil {
		c.ErrorGettingDataFromDB(err)
		return
	}

	// Updating user last request
	user := &model.UserShortInfo{UserID: userProfile.UserID}
	if err := user.UpdateLastRequest(); err != nil {
		c.SendError(http.StatusBadRequest, err.Error(), "Error when updating last request")
		return
	}

	c.SendSuccess()
}

// UpdateFullUserInfo ...
func UpdateFullUserInfo(w http.ResponseWriter, r *http.Request) {

	userProfile := &model.UserInformation{}
	c := &c.Context{Res: w, Req: r, Data: userProfile}
	// if err := c.JSONDecoder(); err != nil {
	// 	return
	// }
	userProfile.UserFullName = "boymens cool"

	userID := chi.URLParam(r, "user-id")

	claims, err := middleware.GetJWTClaims(w, r)
	if err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "")
		return
	}

	if fmt.Sprint(claims["user_id"].(float64)) != (userID) {
		c.SendError(http.StatusForbidden, "", "Cannot modify someone else data")
		return
	}

	_, err = db.DB.Model(userProfile).
		Where("user_id = ?", userID).
		Column("user_full_name", "birthdate", "bio", "fb_link", "twitter_link", "ig_link", "sex").
		Update()
	if err != nil {
		c.ErrorUpdatingDB(err)
		return
	}

	// Updating user last request
	user := &model.UserShortInfo{UserID: userProfile.UserID}

	if err = user.UpdateLastRequest(); err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "Error when updating last request")
		return
	}

	c.SendSuccess()
}

// UpdateUserShortInfo ...
func UpdateUserShortInfo(w http.ResponseWriter, r *http.Request) {

	userShortInfo := &model.UserShortInfo{}
	c := &c.Context{Res: w, Req: r, Data: userShortInfo}
	if err := c.JSONDecoder(); err != nil {
		return
	}

	if err := CheckUserName(userShortInfo.UserName); err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userShortInfo.Password), bcrypt.DefaultCost)

	userShortInfo.Password = string(hashedPassword)
	userShortInfo.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = db.DB.Model(userShortInfo).
		Where("user_id = ?user_id").
		Column("user_name", "country_id", "password", "email", "last_request").
		Update()
	if err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	c.SendSuccess()
}
