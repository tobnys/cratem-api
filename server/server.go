package server

import (
	"github.com/gin-contrib/cors"
	"github.com/tobnys/cratem-api/cfg"
)

func Initialize() {
	r := Router()
	r.Use(cors.Default())
	r.Run(cfg.HOST + ":" + cfg.PORT)
}
