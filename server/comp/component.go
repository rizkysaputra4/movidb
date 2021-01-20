package comp

import (
	"encoding/json"
	"fmt"
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
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// BasicResponse ...
func BasicResponse(w http.ResponseWriter, code int, errorMessage string, data interface{}) {
	err := &StatusResponse{
		Status:  code,
		Message: errorMessage,
		Data:    data,
	}
	fmt.Println("basic response")
	response, _ := json.Marshal(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
