package comp

import (
	"github.com/rizkysaputra4/moviwiki/server/db"
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
