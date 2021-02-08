package handler

import (
	"net/http"

	c "github.com/rizkysaputra4/moviwiki/server/context"
	"github.com/rizkysaputra4/moviwiki/server/http/middleware"
)

// GetMyRole ...
func GetMyRole(w http.ResponseWriter, r *http.Request) {
	claims, errJWT := middleware.GetJWTClaims(w, r)
	c := &c.Context{Res: w, Req: r}
	claimRole := claims["role"]
	if claimRole == nil && errJWT == nil {
		c.SendError(http.StatusUnauthorized, "Token claims empty", "")
		return
	}

	type Res struct {
		Role int `json:"role"`
	}

	var role Res

	if errJWT != nil {
		role.Role = 41
	} else {
		role.Role = int(claimRole.(float64))
	}

	c.SendSuccess(role)
}
