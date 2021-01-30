package handler

import (
	"fmt"
	"net/http"
	"time"

	c "github.com/rizkysaputra4/moviwiki/server/context"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
	"golang.org/x/crypto/bcrypt"
)

// RegisterNewAdmin ...
func RegisterNewAdmin(w http.ResponseWriter, r *http.Request) {
	newAdmin := &model.UserShortInfo{}
	c := &c.Context{Res: w, Req: r, Data: newAdmin}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	if newAdmin.Password == "" || len(newAdmin.Password) < 6 {
		c.SendError(http.StatusBadRequest, "Password too Short, must be 6 character long", "")
		return
	}

	if err := CheckUserName(newAdmin.UserName); err != nil {
		c.SendError(http.StatusBadRequest, err.Error(), "")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAdmin.Password), bcrypt.DefaultCost)

	newAdmin.Password = string(hashedPassword)
	newAdmin.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")

	_, err = db.DB.Model(newAdmin).Insert()
	if err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	completeUserData := &model.UserInformation{
		UserID:       newAdmin.UserID,
		RegisterDate: time.Now().UTC().Format("2006-01-02"),
	}

	_, err = db.DB.Model(completeUserData).Insert()
	if err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	c.SendSuccess()
}

// ChangeAdminLevel handler to promote user into admin
func ChangeAdminLevel(w http.ResponseWriter, r *http.Request) {
	admin := &model.UserShortInfo{}
	newObj := admin

	c := &c.Context{Res: w, Req: r, Data: admin}
	if err := c.JSONDecoder(); err != nil {
		return
	}

	authStatus, err := RoleOrderPermission(w, r, newObj, admin.Role)
	fmt.Println(authStatus, err)

	if err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "Error when checking authorization order")
		return
	}

	if !authStatus {
		c.SendError(http.StatusUnauthorized, "Unauthorized", "Peasant cannot promote king")
		return
	}

	_, err = db.DB.Model(admin).
		Where("user_id = ?user_id").
		Column("role").
		Update()
	if err != nil {
		c.SendError(http.StatusBadRequest, err.Error(), "Error when update data into db")
	}

	c.SendSuccess()
}

// AddAnotherIdentifier is adding another flag when needed
func AddAnotherIdentifier(w http.ResponseWriter, r *http.Request) {
	newIdentifier := &model.Identifier{}
	c := &c.Context{Res: w, Req: r, Data: newIdentifier}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	fmt.Println(newIdentifier)

	if _, err := db.DB.Model(newIdentifier).Insert(); err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "Error when inserting data into db")
		return
	}

	c.SendSuccess()

}
