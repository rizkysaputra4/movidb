package session

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var appSession *sessions.Session;

var authKey = []byte("qwer")
var encKey = []byte("asdf")

var store = sessions.NewCookieStore(authKey, encKey)    

func initSession(r *http.Request) *sessions.Session {

    log.Println("session before get", appSession)

    if appSession != nil {    
        return appSession;    
    }

    session, err := store.Get(r, "golang_cookie")
    appSession = session;

    log.Println("session after get", session)
    if err != nil {
        panic(err)
    }
    return session
}

// SetSessionValue ...
func SetSessionValue(w http.ResponseWriter, r *http.Request, key, value string) {
    session := initSession(r)
    session.Values[key] = value
    fmt.Printf("set session with key %s and value %s\n", key, value)
    session.Save(r, w)
}

// GetSessionValue ...
func GetSessionValue(w http.ResponseWriter, r *http.Request, key string) string {   
    session := initSession(r)
    valWithOutType := session.Values[key]
    fmt.Printf("valWithOutType: %s\n", valWithOutType)
    value, ok := valWithOutType.(string)
    log.Println("returned value: ", value);

    if !ok {
        fmt.Println("cannot get session value by key: " + key)
    }
    return value
}