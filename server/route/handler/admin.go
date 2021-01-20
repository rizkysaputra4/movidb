package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rizkysaputra4/moviwiki/server/comp"
	res "github.com/rizkysaputra4/moviwiki/server/comp"
	. "github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
	"golang.org/x/crypto/bcrypt"
)

// RegisterNewAdmin ...
func RegisterNewAdmin(w http.ResponseWriter, r *http.Request) {
	newAdmin := &model.UserShortInfo{}

	if err := json.NewDecoder(r.Body).Decode(&newAdmin); err != nil {
		res.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when decode request body")
		return
	}

	if newAdmin.Password == "" || len(newAdmin.Password) < 6 {
		res.BasicResponse(w, http.StatusBadRequest, "Password too Short, must be 6 character long", "")
		return
	}

	if err := CheckUserName(newAdmin.UserName); err != nil {
		res.BasicResponse(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAdmin.Password), bcrypt.DefaultCost)

	newAdmin.Password = string(hashedPassword)
	newAdmin.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = DB.Model(newAdmin).Insert()
	if err != nil {
		res.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when inserting user credential into db")
		return
	}

	completeUserData := &model.UserInformation{
		UserID:       newAdmin.UserID,
		RegisterDate: time.Now().UTC().Format("2006-01-02"),
	}

	_, err = DB.Model(completeUserData).Insert()
	if err != nil {
		res.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when inserting user data profile to db")
		return
	}

	comp.BasicResponse(w, http.StatusOK, "OK", newAdmin)
}

// ChangeAdminLevel handler to promote user into admin
func ChangeAdminLevel(w http.ResponseWriter, r *http.Request) {
	admin := &model.UserShortInfo{}
	newObj := admin

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		res.BasicResponse(w, http.StatusBadGateway, "Error when decode request body", err.Error())
		return
	}
	fmt.Println(admin.Role)
	authStatus, err := RoleOrderPermission(w, r, newObj, admin.Role)
	fmt.Println(authStatus, err)

	if err != nil {
		res.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when checking authorization order")
		return
	}

	if !authStatus {
		res.BasicResponse(w, http.StatusUnauthorized, "Unauthorized", "Peasant cannot promote king")
		return
	}

	_, err = DB.Model(&admin).
		Where("user_id = ?user_id").
		Column("role").
		Update()
	if err != nil {
		res.BasicResponse(w, http.StatusBadRequest, err.Error(), "Error when update data into db")
	}

	res.BasicResponse(w, http.StatusOK, "", admin)
}
