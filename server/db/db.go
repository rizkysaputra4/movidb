package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/rizkysaputra4/moviwiki/server/env"
)

// DB is global var for DB model
var DB *pg.DB

// Init is initiate db connection
func init() {
	c := env.GetConfiguration()
	DB = pg.Connect(&pg.Options{
		User:     c.PostgresUserName,
		Password: c.PostgresPass,
		Database: c.PostgresDB,
	})

	fmt.Println("DB INIT ...")
}
