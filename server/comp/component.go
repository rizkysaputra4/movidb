package comp

import (
	"encoding/json"
	"net/http"

	. "github.com/rizkysaputra4/moviwiki/server/db"
)

// CheckIfExist ...
func CheckIfExist(column string, element string, c interface{}) bool {
	err := DB.Model(c).Column(column).
		Where(column+"= ?", element).
		Select()

	if err == nil {
		return true
	}

	return false
}

// StatusResponse ...
type StatusResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// BasicResponse ...
func BasicResponse(w http.ResponseWriter, code int, isError bool, errorMessage string) {
	err := &StatusResponse{
		Status:  isError,
		Message: errorMessage,
	}
	response, _ := json.Marshal(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ResJSON ...
func ResJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

