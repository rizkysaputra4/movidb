package app

import (
	"fmt"
	"time"

	"github.com/rizkysaputra4/moviwiki/server/db"
	"github.com/rizkysaputra4/moviwiki/server/route"
)

// Init initialize function
func Init(){
	db.Init()
	fmt.Println(time.Now().Unix())

	route.Init()
}
