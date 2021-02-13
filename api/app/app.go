package app

import (
	"fmt"
	"time"

	"github.com/rizkysaputra4/moviwiki/api/http/route"
)

// Init initialize function
func Init() {
	route.InitRoute()
	fmt.Println(time.Now().Unix())

}
