package main

import (
	"github.com/SyarifKA/himbara/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Static("/img/profile", "./img/profile")
	// r.Static("/img/event", "./img/event")
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	// r.Use(cors.New(config))
	routers.RoutersCombine(r)

	r.Run("localhost:8888")
}
