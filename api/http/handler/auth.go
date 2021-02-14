package handler

import (
	"net/http"
	"time"

	"github.com/rizkysaputra4/moviwiki/api/comp"
	c "github.com/rizkysaputra4/moviwiki/api/context"
	"github.com/rizkysaputra4/moviwiki/api/db"
	"github.com/rizkysaputra4/moviwiki/api/http/middleware"
	"github.com/rizkysaputra4/moviwiki/api/model"
	"golang.org/x/crypto/bcrypt"
)

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

// ReceivedUserInfo ...
type ReceivedUserInfo struct {
	model.UserShortInfo
	model.UserInformation
}

// RegisteringNewUser is handling register request
func RegisteringNewUser(w http.ResponseWriter, r *http.Request) {
	fullInfo := &ReceivedUserInfo{}
	c := &c.Context{
		Res:  w,
		Req:  r,
		Data: fullInfo,
	}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	user := fullInfo.UserShortInfo

	// user := &model.UserShortInfo{
	// 	UserName:  fullInfo.UserName,
	// 	CountryID: fullInfo.CountryID,
	// 	Password:  fullInfo.Password,
	// 	Email:     fullInfo.Email,
	// }

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
	_, err = db.DB.Model(&user).
		Column(col...).
		Insert()

	if err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	// completeUserData := &model.UserInformation{
	// 	UserFullName: fullInfo.UserFullName,

	// }
	completeUserData := fullInfo.UserInformation
	completeUserData.UserID = user.UserID
	completeUserData.RegisterDate = time.Now().UTC().Format("2006-01-02")

	_, err = db.DB.Model(&completeUserData).Insert()
	if err != nil {
		c.ErrorInsertingDataIntoDB(err)
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
	userInfo := &UserFull{}
	c := &c.Context{Res: w, Req: r, Data: pw}
	if err := c.JSONDecoder(); err != nil {
		return
	}

	_, err := db.DB.Query(userInfo,
		`select  
			user_short_info.user_id, 
			user_information.user_id,
			user_name,
			user_full_name,
			country_id,
			email,
			password,
			sex,
			role 
		from 
			user_short_info 
		inner join 
			user_information 
		on 
			user_short_info.user_id = user_information.user_id
		where 
			user_name = ?`, pw.Username)

	// err := db.DB.Model(userInfo).
	// 	Where("user_name = ?", pw.Username).
	// 	WhereOr("email = ?", pw.Email).
	// 	Column("user_id", "user_name", "email", "password", "role").
	// 	Select()

	if err != nil {
		c.ErrorGettingDataFromDB(err)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(pw.Pw)); err != nil {
		c.SendError(http.StatusBadRequest, err.Error(), "Error when bcrypting user password")
		return
	}

	userInfo.UserShortInfo.UserID = userInfo.UserInformation.UserID

	middleware.StoreJWT(w, r, userInfo.UserShortInfo.UserID, userInfo.Role)

	if userInfo.Role < 11 {
		middleware.StoreSession(w, r, userInfo.UserShortInfo.UserID, userInfo.Role)
	}

	userInfo.UpdateLastRequest()
	userInfo.Password = ""
	c.SendSuccess(userInfo)
}

// LogOut ...
func LogOut(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteSession(w, r)
	middleware.DeleteJWTFromCookie(w, r)

	c := &c.Context{Res: w, Req: r, Data: "Logout success"}
	c.SendSuccess()
}
