package main

import (
	routers "pastebin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", routers.GetIndex)
	router.POST("/post", routers.Getpost)
	router.GET("/pastes/:url", routers.Getpaste)
	router.StaticFile("/favicon.ico", "./web/assets/girl.ico")
	router.Static("/static", "./web/assets")
	// router.GET("/test", routers.Test)
	router.Run("127.0.0.1:8080")
}
