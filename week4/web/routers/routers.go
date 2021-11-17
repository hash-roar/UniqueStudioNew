package routers

import "github.com/gin-gonic/gin"

var router = gin.Default()

func init() {
	router.LoadHTMLGlob("web/templets/*")
	router.StaticFile("/favicon.ico", "web/assets/girl.ico")
	router.Static("/static", "web/assets")

	// init auth api
	initAuthApi()

	//
}

func initAuthApi() {
	router.GET("/loginpage", LoginPage)
	router.POST("/login", Login)
	router.GET("/rigisterpage", RigisterPage)
	router.POST("/rigister", Rigister)
}

func Run() {
	router.Run("127.0.0.1:8080")
}
