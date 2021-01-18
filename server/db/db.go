package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/rizkysaputra4/moviwiki/server/env"
	"gopkg.in/boj/redistore.v1"
)

// DB is global var for DB model
var DB *pg.DB

// Store ...
var Store *redistore.RediStore

// Init is initiate db connection
func init() {
	Store, _ = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))

	c := env.GetConfiguration()
	DB = pg.Connect(&pg.Options{
		User:     c.PostgresUserName,
		Password: c.PostgresPass,
		Database: c.PostgresDB,
	})

	fmt.Println("DB INIT ...")
}
