package handler

import (
	"fmt"
	"net/http"
	"time"

	c "github.com/rizkysaputra4/moviwiki/api/context"
	"github.com/rizkysaputra4/moviwiki/api/db"
	"github.com/rizkysaputra4/moviwiki/api/http/middleware"
	"github.com/rizkysaputra4/moviwiki/api/model"
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
		c.SendError(http.StatusForbidden, "Unauthorized", "Peasant cannot promote king")
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

// UserFull combine usershortinfo and userinformation
type UserFull struct {
	model.UserShortInfo
	model.UserInformation
	IsOnline string `json:"is_online"`
}

func (u *UserFull) addIsOnline(time string) {
	u.IsOnline = time
}

// GetAdminList ...
func GetAdminList(w http.ResponseWriter, r *http.Request) {

	c := &c.Context{Res: w, Req: r}

	var users []UserFull

	_, err := db.DB.Query(&users,
		`select  
			user_short_info.user_id, 
			user_name,
			user_full_name,
			country_id, 
			email,
			sex,
			role, 
			last_request, 
			ig_link, 
			fb_link, 
			monthly_contributor_points
		from 
			user_short_info 
		inner join 
			user_information 
		on 
			user_short_info.user_id = user_information.user_id
		where 
			role < 21`)

	if err != nil {
		c.ErrorGettingDataFromDB(err)
		return
	}

	for i := range users {
		lastReq, err := time.Parse("2006-01-02 15:04:05", users[i].LastRequest)

		if err != nil {
			c.SendError(http.StatusBadRequest, err.Error(), "error when parsing time")
			return
		}

		fromLastRequest := time.Now().Sub(lastReq)

		switch {
		case fromLastRequest < time.Minute:
			users[i].addIsOnline("online")
		case fromLastRequest > (time.Hour * 730): // 730 hours == 1 month
			users[i].addIsOnline(fmt.Sprint(users[i].LastRequest))
		case fromLastRequest > (time.Hour * 48):
			users[i].addIsOnline(fmt.Sprint(int(fromLastRequest/(time.Hour*24))) + " days ago")
		case fromLastRequest > (time.Hour * 24):
			users[i].addIsOnline("Yesterday")
		case fromLastRequest > (time.Hour):
			users[i].addIsOnline(fmt.Sprint(int(fromLastRequest.Hours())) + " hours ago")
		case fromLastRequest < time.Minute*2:
			users[i].addIsOnline("One minute ago")
		case fromLastRequest < time.Hour:
			users[i].addIsOnline(fmt.Sprint(int(fromLastRequest.Minutes())) + " minutes ago")
		}

	}

	c.SendSuccess(users)
}

// ResultStruct ...
type ResultStruct struct {
	model.UserShortInfo
	Count int `pg:"result" json:"-"`
}

// SendSearchResult ...
type SendSearchResult struct {
	SearchResult []ResultStruct `json:"result"`
	Count        int            `json:"count"`
}

// SearchUser ...
func SearchUser(w http.ResponseWriter, r *http.Request) {

	var users []ResultStruct
	c := &c.Context{Res: w, Req: r, Data: users}

	uid := r.URL.Query().Get("uid")
	uid = "%" + uid + "%"
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	if limit == "" {
		limit = "10"
	}

	if offset == "" {
		offset = "0"
	}

	if uid == "" {
		uid = "admin"
	}

	_, err := db.DB.Query(&users,
		`SELECT
			user_id,
			user_name,
			role, COUNT(*) OVER() AS result
		FROM
			user_short_info
		WHERE
			user_name like ?0

		ORDER BY user_name ASC
		LIMIT ?1 OFFSET ?2`, uid, limit, offset)

	if err != nil {
		c.ErrorGettingDataFromDB(err)
		return
	}

	if users == nil {
		c.SendError(http.StatusNotFound, "Not Found", "UserName Not Found")
		return
	}

	resData := &SendSearchResult{SearchResult: users, Count: users[0].Count}
	c.SendSuccess(resData)

}

// RoleOrderPermission ...
func RoleOrderPermission(w http.ResponseWriter, r *http.Request, obj interface{}, requestedRole int) (bool, error) {

	claims, _ := middleware.GetJWTClaims(w, r)

	claimRole := claims["role"]
	if claimRole == nil && claims != nil {
		err := fmt.Errorf("invalid token")
		return false, err
	}

	var subjectRole int
	if claimRole == nil {
		subjectRole = 41
	} else {
		subjectRole = int(claimRole.(float64))
	}

	var objRoleInDB int

	err := db.DB.Model(obj).
		Where("user_id = ?user_id").
		Column("role").Select(&objRoleInDB)
	fmt.Println("objRole", objRoleInDB)
	fmt.Println("requested role", requestedRole)
	fmt.Println("subjectRole", subjectRole)
	fmt.Println("obj", obj)
	if err != nil {
		return false, err
	}

	if subjectRole >= objRoleInDB || subjectRole >= requestedRole {
		return false, nil
	}

	return true, nil
}
