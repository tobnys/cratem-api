package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/tobnys/cratem-api/cfg"
)

func Initialize() {
	r := Router()
	r.Use(cors.Default())
	fmt.Println(cfg.HOST + ":" + cfg.PORT)
	r.Run(cfg.HOST + ":" + cfg.PORT)
}
