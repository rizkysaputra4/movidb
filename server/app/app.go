package app

import (
	"fmt"
	"time"

	"github.com/rizkysaputra4/moviwiki/server/route"
)

// Init initialize function
func Init() {
	route.InitRoute()
	fmt.Println(time.Now().Unix())

}
