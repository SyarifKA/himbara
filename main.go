package main

import (
	"os"

	"github.com/SyarifKA/himbara/lib"
	"github.com/SyarifKA/himbara/logs/logger"
	"github.com/SyarifKA/himbara/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	lib.ConnectDB()
	r := gin.Default()
	// r.Static("/img/profile", "./img/profile")
	// r.Static("/img/event", "./img/event")
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	// r.Use(cors.New(config))

	_ = godotenv.Load()

	// initialize config log
	os.MkdirAll("logs/log", os.ModePerm)

	// Init logger config
	logger.InitLogger(&logger.Config{
		Formatter: &logger.TextFormatter,
		Level:     logger.InfoLevel,
		LogName:   "application.log",
	})
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	routers.RoutersCombine(r)

	r.Run("localhost:8888")
}
