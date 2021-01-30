package handler

import (
	"fmt"
	"net/http"

	c "github.com/rizkysaputra4/moviwiki/server/context"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/http/middleware"
	"github.com/rizkysaputra4/moviwiki/server/model"
)

// AddNewMovieType is adding new movie type
func AddNewMovieType(w http.ResponseWriter, r *http.Request) {
	newMovieType := &model.MovieType{}
	c := &c.Context{Res: w, Req: r, Data: newMovieType}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	if _, err := db.DB.Model(newMovieType).Insert(); err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	c.SendSuccess()
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

	if subjectRole > objRoleInDB || subjectRole >= requestedRole {
		return false, nil
	}

	return true, nil
}

// type movie struct {
// 	description model.MovieInformation `json:"description"`
// 	characters  []model.MovieCharacter `json:"characters"`
// 	genre
// }

// AddNewMovieByAdmin ...
func AddNewMovieByAdmin(w http.ResponseWriter, r *http.Request) {
	newMovie := &model.MovieInformation{}
	// character := &model.MovieCharacter{}
	c := &c.Context{Res: w, Req: r, Data: newMovie}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	if _, err := db.DB.Model(newMovie).Insert(); err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	c.SendSuccess()
}

// EditMovieData ...
func EditMovieData(w http.ResponseWriter, r *http.Request) {
	movie := &model.MovieInformation{}
	c := &c.Context{Res: w, Req: r, Data: movie}
	if err := c.JSONDecoder(); err != nil {
		return
	}

	if _, err := db.DB.Model(movie).Where("movie_id = ?movie_id").Update(); err != nil {
		c.ErrorUpdatingDB(err)
		return
	}

	c.SendSuccess()
}

// AddNewRoleListByAdmin is adding new role like actress, actor, storywriter, director, etc.
func AddNewRoleListByAdmin(w http.ResponseWriter, r *http.Request) {
	newRole := &model.MoviePeopleRole{}
	c := &c.Context{Res: w, Req: r, Data: newRole}

	if err := c.JSONDecoder(); err != nil {
		return
	}

	if _, err := db.DB.Model(newRole).Insert(); err != nil {
		c.ErrorInsertingDataIntoDB(err)
		return
	}

	c.SendSuccess()
}

// UpdateNewRoleListByAdmin ...
func UpdateNewRoleListByAdmin(w http.ResponseWriter, r *http.Request) {
	newRole := &model.MoviePeopleRole{}
	c := &c.Context{Res: w, Req: r, Data: newRole}
	if err := c.JSONDecoder(); err != nil {
		return
	}

	if newRole.RoleID < 1 {
		c.SendError(http.StatusBadRequest, "role_id cannot nil", "")
		return
	}

	if _, err := db.DB.Model(newRole).Where("role_id = ?role_id").Update(); err != nil {
		c.ErrorUpdatingDB(err)
		return
	}

	c.SendSuccess()
}
