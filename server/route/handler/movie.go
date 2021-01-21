package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rizkysaputra4/moviwiki/server/comp"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
)

// AddNewMovieByAdmin ...
func AddNewMovieByAdmin(w http.ResponseWriter, r *http.Request) {
	newMovie := &model.MovieInformation{}

	if err := json.NewDecoder(r.Body).Decode(newMovie); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when decoding request body")
		return
	}

	if _, err := db.DB.Model(newMovie).Insert(); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when inserting data into db")
		return
	}

	comp.BasicResponse(w, http.StatusOK, "OK", newMovie)
}

// EditMovieData ...
func EditMovieData(w http.ResponseWriter, r *http.Request) {
	movie := &model.MovieInformation{}

	if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when decoding request body")
		return
	}

	if _, err := db.DB.Model(movie).Where("movie_id = ?movie_id").Update(); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when updating data into db")
		return
	}

	comp.BasicResponse(w, http.StatusOK, "OK", movie)
}

// AddNewRoleList is adding new role like actress, actor, storywriter, director, etc.
