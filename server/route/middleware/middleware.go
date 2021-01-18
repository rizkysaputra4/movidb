package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/admpub/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/rizkysaputra4/moviwiki/server/comp"
	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/env"
)

// var store = sessions.NewCookieStore([]byte("moviwiki"))

// AdminSession ...
func AdminSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := db.Store.Get(r, "moviwiki-session")
		if err != nil {
			log.Error(err.Error())
		}

		role := session.Values["role"]
		fmt.Println(role)

		next.ServeHTTP(w, r)
	})
}

// StoreSession used when admin logged in
func StoreSession(w http.ResponseWriter, r *http.Request, userID int, role int) {

	session, err := db.Store.Get(r, "moviwiki-session")

	fmt.Println(session.Values["role"])
	if err != nil {
		log.Error(err.Error())
	}

	session.Values["user-id"] = userID
	session.Values["role"] = role

	session.Options = &sessions.Options{
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
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
		comp.BasicResponse(w, http.StatusInternalServerError, false, err.Error())
	}
}

//StoreJWT ...
func StoreJWT(w http.ResponseWriter, r *http.Request, userID int, role int) {
	token, err := CreateToken(uint64(userID), role)
	if err != nil {
		comp.BasicResponse(w, http.StatusInternalServerError, false, err.Error())
	}

	c := &http.Cookie{
		Name:     "Auth-Token",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 168),
		//	Domain:   "localhost",
		Path: "/",
	}

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

	return claims, nil
}

// DeleteJWT ...
func DeleteJWT(w http.ResponseWriter, r *http.Request) {

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
	atClaims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(env.TokenKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
