package db

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// DB is global var for DB model
var DB *pg.DB

// Init is initiate db connection
func Init() error {
	DB = pg.Connect(&pg.Options{
		User: "postgres", 
		Password: "postgres",
		Database: "moviwiki",
	})

	ctx := context.Background()

	if err := DB.Ping(ctx); err != nil {
    	return err
	}

	return nil
}