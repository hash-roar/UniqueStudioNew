package routers

import "github.com/gin-gonic/gin"

var router = gin.Default()

func init() {
	if gin.Mode() == gin.TestMode {
		router.LoadHTMLGlob("./../templets/*")
	} else {
		router.LoadHTMLGlob("D:\\working\\UniqueStudioNew\\week4\\web\\templets\\*")
	}
	router.StaticFile("/favicon.ico", "web/assets/girl.ico")
	router.Static("/static", "web/assets")

	// init login api
	initLoginApi()

	// init oauth api
	initOauthApi()
	//init daylireportApi
	initDailyReportApi()
}

func initLoginApi() {
	router.GET("/loginpage", LoginPage)
	router.POST("/login", Login)
	router.GET("/registerpage", registerPage)
	router.POST("/register", register)
}

func initOauthApi() {
	router.GET("/oauth-page", Oauth2Page)
}

func initDailyReportApi() {
	router.GET("/", Index)
}

func Run() {
	router.Run("127.0.0.1:8080")
}
