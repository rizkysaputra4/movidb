package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rizkysaputra4/moviwiki/server/comp"
	. "github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
	"golang.org/x/crypto/bcrypt"
)

// RegisterNewAdmin ...
func RegisterNewAdmin(w http.ResponseWriter, r *http.Request) {
	newAdmin := &model.UserShortInfo{}

	if err := json.NewDecoder(r.Body).Decode(&newAdmin); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	if newAdmin.Password == "" || len(newAdmin.Password) < 6 {
		comp.BasicResponse(w, http.StatusBadRequest, false, "Password too Short, must be 6 character long")
		return
	}

	if err := CheckUserName(newAdmin.UserName); err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, false, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAdmin.Password), bcrypt.DefaultCost)

	newAdmin.Password = string(hashedPassword)
	newAdmin.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = DB.Model(newAdmin).Insert()
	if err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	completeUserData := &model.UserInformation{
		UserID:       newAdmin.UserID,
		RegisterDate: time.Now().UTC().Format("2006-01-02"),
	}

	_, err = DB.Model(completeUserData).Insert()
	if err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	comp.ResJSON(w, http.StatusOK, newAdmin)
}

// ChangeAdminLevel handler to promote user into admin
func ChangeAdminLevel(w http.ResponseWriter, r *http.Request) {
	admin := &model.UserShortInfo{}

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		comp.BasicResponse(w, http.StatusBadGateway, false, err.Error())
		return
	}

	_, err := DB.Model(admin).
		Where("user_id = ?user_id").
		Column("role").
		Update()
	if err != nil {
		comp.BasicResponse(w, http.StatusBadRequest, false, err.Error())
	}

	comp.ResJSON(w, http.StatusOK, admin)
}
