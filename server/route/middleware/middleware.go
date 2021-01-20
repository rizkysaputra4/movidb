package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/admpub/log"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/rizkysaputra4/moviwiki/server/comp"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/env"
)

// var store = sessions.NewCookieStore([]byte("moviwiki"))

// RoleEnforcer ...
func RoleEnforcer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, errJWT := GetJWTClaims(w, r)

		claimRole := claims["role"]
		if claimRole == nil && errJWT == nil {
			comp.BasicResponse(w, http.StatusUnauthorized, "Token claims empty", "")
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
				comp.BasicResponse(w, http.StatusUnauthorized, err.Error(), "Session expired")
				return
			}

			sessionRole := session.Values["role"]
			if sessionRole == nil {
				DeleteJWTFromCookie(w, r)
				comp.BasicResponse(w, http.StatusUnauthorized, "Session nil", "Role admin but session not found")
				return
			}

			role = sessionRole.(int)
		}

		fmt.Printf("uid %v. role %v, path %v, method %v \n", uid, role, path, method)

		access, err := DefineAccess(role, path, method)
		fmt.Println("Casbin access; ", access, err)

		if access {
			next.ServeHTTP(w, r)
			return
		}

		comp.BasicResponse(w, http.StatusUnauthorized, "Unauthorized", "Token not found")
		return
	})
}

// StoreSession used when admin logged in
func StoreSession(w http.ResponseWriter, r *http.Request, userID int, role int) {

	session, err := db.Store.Get(r, "moviwiki-session")

	if err != nil {
		log.Error(err.Error())
	}

	session.Values["user-id"] = userID
	session.Values["role"] = role

	session.Options = &sessions.Options{
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	if err = sessions.Save(r, w); err != nil {
		fmt.Println(err.Error())
	}

	db.Store.SetMaxAge(3600)
}

// DeleteSession ...
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := db.Store.Get(r, "moviwiki-session")

	if err != nil {
		log.Error(err.Error())
	}

	session.Options.MaxAge = -1
	if err = sessions.Save(r, w); err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when saving the session")
	}
}

//StoreJWT ...
func StoreJWT(w http.ResponseWriter, r *http.Request, userID int, role int) {
	token, err := CreateToken(uint64(userID), role)
	if err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, err.Error(), "Error when creating JWT")
	}

	c := &http.Cookie{
		Name:     "Auth-Token",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 168),
		//	Domain:   "localhost",
		Path: "/",
	}
	fmt.Println(token)
	http.SetCookie(w, c)
}

// GetJWTClaims ...
func GetJWTClaims(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, error) {
	Token, err := r.Cookie("Auth-Token")
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(Token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.GetConfiguration().TokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	if Token == nil {
		return nil, err
	}

	return claims, nil
}

// DeleteJWTFromCookie ...
func DeleteJWTFromCookie(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:   "Auth-Token",
		MaxAge: -1,
		Value:  "",
		Path:   "/",
	}

	http.SetCookie(w, c)
}

// CreateToken ...
func CreateToken(userid uint64, role int) (string, error) {
	env := env.GetConfiguration()
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["role"] = role

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(env.TokenKey))
	if err != nil {
		return "", err
	}
	return token, nil
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
	fmt.Println("sub: ", sub)
	ok, err := e.Enforce(sub, path, method)
	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	return false, nil
}

// UpdateSessionExp ...
func UpdateSessionExp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := db.Store.Get(r, "moviwiki-session")
		if err != nil {
			comp.BasicResponse(w, http.StatusUnauthorized, err.Error(), "error when trying to get session")
			return
		}

		session.Options.MaxAge = 3600
		session.Options.HttpOnly = true
		db.Store.SetMaxAge(3600)
		sessions.Save(r, w)
		next.ServeHTTP(w, r)

	})
}

// UpdateJWTExp ...
func UpdateJWTExp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t, err := r.Cookie("Auth-Token")
		if err == nil {
			t.Expires = time.Now().Add(time.Hour * 168)
			t.HttpOnly = true
			t.Path = "/"
			http.SetCookie(w, t)
		}

		// c := &http.Cookie{
		// 	Name:     "Auth-Token",
		// 	HttpOnly: true,
		// 	Expires:  time.Now().Add(time.Hour * 168),
		// 	//	Domain:   "localhost",
		// 	Path: "/",
		// }

		next.ServeHTTP(w, r)
	})
}
