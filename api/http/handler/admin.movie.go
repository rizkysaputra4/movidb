package handler

import (
	"net/http"

	c "github.com/rizkysaputra4/moviwiki/api/context"
	"github.com/rizkysaputra4/moviwiki/api/db"
	"github.com/rizkysaputra4/moviwiki/api/model"
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
