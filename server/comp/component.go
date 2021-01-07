package comp

import (
	. "github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/model"
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

// BasicResponse ...
func BasicResponse(errorID bool, errorMessage string) model.StatusResponse {
	err := &model.StatusResponse{
		Status:  errorID,
		Message: errorMessage,
	}

	return *err
}
