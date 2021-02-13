package comp

import (
	"github.com/rizkysaputra4/moviwiki/api/db"
)

// CheckIfExist ...
func CheckIfExist(column string, element string, c interface{}) bool {
	err := db.DB.Model(c).Column(column).
		Where(column+"= ?", element).
		Select()

	if err == nil {
		return true
	}

	return false
}

// StatusResponse ...
// type StatusResponse struct {
// 	Status  int         `json:"status"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// // BasicResponse ...
// func BasicResponse(w http.ResponseWriter, code int, errorMessage string, data interface{}) {
// 	err := &StatusResponse{
// 		Status:  code,
// 		Message: errorMessage,
// 		Data:    data,
// 	}

// 	response, _ := json.Marshal(err)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)

// }

// // JSONDecoder ...
// func JSONDecoder(w http.ResponseWriter, r *http.Request, model interface{}) error {
// 	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
// 		BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when decoding request body")
// 		return err
// 	}

// 	return nil
// }
