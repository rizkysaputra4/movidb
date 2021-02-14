package middleware

import (
	"github.com/rizkysaputra4/moviwiki/api/model"
)

// UpdateLastRequest is middleware that updating last request for every api request
func UpdateLastRequest(uid int) {

	user := model.UserShortInfo{UserID: uid}
	user.UpdateLastRequest()
}
