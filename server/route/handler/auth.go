package handler

import (
	"net/http"
	"time"

	"github.com/rizkysaputra4/moviwiki/server/comp"
	c "github.com/rizkysaputra4/moviwiki/server/context"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
	"github.com/rizkysaputra4/moviwiki/server/route/middleware"
	"golang.org/x/crypto/bcrypt"
)

// CheckIfEmailExist ...
func CheckIfEmailExist(w http.ResponseWriter, r *http.Request) {
	shortInfo := &model.UserShortInfo{}
	c := &c.Context{Res: w, Req: r, Data: shortInfo}
	if err := c.JSONDecoder(); err != nil || shortInfo.Email == "" {
		return
	}

	column := "email"
	if isExist := comp.CheckIfExist(column, shortInfo.Email, shortInfo); isExist {
		c.SendError(http.StatusBadRequest, "Email is already exist", "")
		return
	}

	c.SendSuccess()
}

// CheckIfUserNameExist ...
func CheckIfUserNameExist(w http.ResponseWriter, r *http.Request) {
	shortInfo := &model.UserShortInfo{}
	c := &c.Context{Res: w, Req: r, Data: shortInfo}
	if err := c.JSONDecoder(); err != nil {
		return
	}

	if err := CheckUserName(shortInfo.UserName); err != nil {
		c.SendError(http.StatusBadRequest, err.Error(), "")
		return
	}

	column := "user_name"
	if isExist := comp.CheckIfExist(column, shortInfo.UserName, shortInfo); isExist {
		c.SendError(http.StatusBadRequest, "UserName is already exist", "")
		return
	}

	c.SendSuccess()
}

// CheckIfUserExist ...
func CheckIfUserExist(w http.ResponseWriter, r *http.Request) {
	info := &model.UserShortInfo{}
	c := &c.Context{Res: w, Req: r, Data: info}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	err := db.DB.Model(info).
		Column("user_id", "user_name", "email").
		Where("user_name = ?user_name").
		WhereOr("email = ?email").
		Select()

	if err != nil {
		c.ErrorGettingDataFromDB(err)
		return
	}

	c.SendSuccess("User Exist")
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
	c := &c.Context{Res: w, Req: r, Data: pw}
	if err := c.JSONDecoder(); err != nil {
		return
	}

	err := db.DB.Model(userInfo).
		Where("user_name = ?", pw.Username).
		WhereOr("email = ?", pw.Email).
		Column("user_id", "user_name", "email", "password", "role").
		Select()

	if err != nil {
		c.ErrorGettingDataFromDB(err)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(pw.Pw)); err != nil {
		c.SendError(http.StatusBadRequest, err.Error(), "Error when bcrypting user password")
		return
	}

	middleware.StoreJWT(w, r, userInfo.UserID, userInfo.Role)

	if userInfo.Role < 11 {
		middleware.StoreSession(w, r, userInfo.UserID, userInfo.Role)
	}

	userInfo.UpdateLastRequest()

	c.SendSuccess()
}

// LogOut ...
func LogOut(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteSession(w, r)
	middleware.DeleteJWTFromCookie(w, r)

	c := &c.Context{Res: w, Req: r, Data: "Logout success"}
	c.SendSuccess()
}

// RegisteringNewUser is handling register request
func RegisteringNewUser(w http.ResponseWriter, r *http.Request) {
	user := &model.UserShortInfo{}
	c := &c.Context{
		Res:  w,
		Req:  r,
		Data: user,
	}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	if user.Password == "" || len(user.Password) < 6 {
		c.SendError(http.StatusBadRequest, "Password too Short, must be 6 character long", "")
		return
	}

	if err := CheckUserName(user.UserName); err != nil {
		c.SendError(http.StatusBadRequest, err.Error(), "")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)
	user.LastRequest = time.Now().UTC().Format("2006-01-02 15:04:05")
	if err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "Error when bcrypting password")
		return
	}

	col := []string{"user_id", "user_name", "country_id", "password", "email", "last_request"}
	_, err = db.DB.Model(c.Data).
		Column(col...).
		Insert()

	if err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	completeUserData := &model.UserInformation{
		UserID:       user.UserID,
		RegisterDate: time.Now().UTC().Format("2006-01-02"),
	}

	_, err = db.DB.Model(completeUserData).Insert()
	if err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	c.SendSuccess()
}
