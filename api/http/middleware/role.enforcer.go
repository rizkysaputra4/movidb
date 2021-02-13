package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	c "github.com/rizkysaputra4/moviwiki/api/context"
	"github.com/rizkysaputra4/moviwiki/api/db"
)

// var store = sessions.NewCookieStore([]byte("moviwiki"))

// RoleEnforcer ...
func RoleEnforcer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, errJWT := GetJWTClaims(w, r)
		c := &c.Context{Res: w, Req: r}
		claimRole := claims["role"]
		if claimRole == nil && errJWT == nil {
			c.SendError(http.StatusUnauthorized, "Token claims empty", "")
			return
		}

		var role int
		uid := claims["user_id"]

		path := r.URL.Path
		method := r.Method

		if errJWT != nil {
			role = 41
		} else {
			role = int(claimRole.(float64))
		}

		if role < 11 {
			session, err := db.Store.Get(r, "moviwiki-session")
			if err != nil {
				c.SendError(http.StatusUnauthorized, err.Error(), "Session expired")
				return
			}

			sessionRole := session.Values["role"]
			if sessionRole == nil {
				DeleteJWTFromCookie(w, r)
				c.SendError(http.StatusUnauthorized, "Session nil", "Role admin but session not found")
				return
			}

			role = sessionRole.(int)
		}

		// fmt.Printf("uid %v. role %v, path %v, method %v \n", uid, role, path, method)

		access, err := DefineAccess(role, path, method)
		// fmt.Println("Casbin access; ", access, err)

		if err != nil {
			c.SendError(http.StatusInternalServerError, err.Error(), "Error when define access")
			return
		}

		if access {

			if role < 25 {
				UpdateLastRequest(int(uid.(float64)))
			}

			next.ServeHTTP(w, r)
			return
		}

		c.SendError(http.StatusUnauthorized, "Unauthorized", "Token not found")
		return
	})
}

var e, err = casbin.NewEnforcer("./config/casbin_model.conf", "./config/casbin_policy.csv")

// DefineAccess ...
func DefineAccess(role int, path string, method string) (bool, error) {
	if err != nil {
		fmt.Println(err)
	}

	var sub string
	if role < 11 {
		sub = "admin"
	} else if role > 11 && role < 21 {
		sub = "member-admin"
	} else if role > 21 && role < 25 {
		sub = "member"
	} else {
		sub = "anonymous"
	}
	// fmt.Println("sub: ", sub)
	ok, err := e.Enforce(sub, path, method)
	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	return false, nil
}
