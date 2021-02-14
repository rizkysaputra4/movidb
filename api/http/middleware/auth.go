package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/admpub/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	c "github.com/rizkysaputra4/moviwiki/api/context"
	"github.com/rizkysaputra4/moviwiki/api/db"
	"github.com/rizkysaputra4/moviwiki/api/env"
)

//StoreJWT ...
func StoreJWT(w http.ResponseWriter, r *http.Request, userID int, role int) {
	c := &c.Context{Res: w, Req: r}

	token, err := CreateToken(uint64(userID), role)
	if err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "Error when creating JWT")
	}

	cookie := &http.Cookie{
		Name:     "Auth-Token",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 168),
		//Domain:   "localhost",
		Path: "/",
	}

	http.SetCookie(w, cookie)
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

// UpdateSessionExp ...
func UpdateSessionExp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := db.Store.Get(r, "moviwiki-session")
		c := &c.Context{Res: w, Req: r}
		if err != nil {
			c.SendError(http.StatusUnauthorized, err.Error(), "error when trying to get session")
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
		var duration int64 = 86400 // seconds or a day
		claims, err := GetJWTClaims(w, r)

		if err == nil && time.Now().Unix()-int64(claims["time_created"].(float64)) > duration {
			StoreJWT(w, r, int(claims["user_id"].(float64)), int(claims["role"].(float64)))
		}

		next.ServeHTTP(w, r)
	})
}

// DeleteSession ...
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := db.Store.Get(r, "moviwiki-session")
	c := &c.Context{Res: w, Req: r}
	if err != nil {
		log.Error(err.Error())
	}

	session.Options.MaxAge = -1
	if err = sessions.Save(r, w); err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "Error when saving the session")
	}
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
	atClaims["time_created"] = time.Now().Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(env.TokenKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
