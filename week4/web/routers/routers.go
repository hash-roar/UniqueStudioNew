package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

var authGroup *gin.RouterGroup

func init() {
	if gin.Mode() == gin.TestMode {
		router.LoadHTMLGlob("./../templets/*")
	} else {
		router.LoadHTMLGlob("D:\\working\\UniqueStudioNew\\week4\\web\\templets\\*")
	}
	router.StaticFile("/favicon.ico", "web/assets/girl.ico")
	router.Static("/static", "web/assets")

	authGroup = router.Group("/auth")
	authGroup.Use(AuthMidware())

	// init login api
	initLoginApi()

	// init oauth api
	initOauthApi(authGroup)
	//init daylireportApi
	initDailyReportApi(authGroup)
}

func initLoginApi() {
	router.GET("/loginpage", LoginPage)
	router.POST("/login", Login)
	router.GET("/registerpage", registerPage)
	router.POST("/register", register)
}

func initOauthApi(group *gin.RouterGroup) {
	group.GET("/oauth-page", Oauth2Page)
	group.GET("/send-auth-code", SendAuthCode)
	router.POST("/oauth", Oauth2)
}

func initDailyReportApi(group *gin.RouterGroup) {
	group.GET("/", Index)
	group.GET("/write-page", WriteReportPage)
	group.POST("/add-report", AddReport)
	group.GET("/admin", AdminPage)
	group.POST("/alt-role", AltPrivilege)
	group.GET("/team-report", TeamReport)
}

func Run() {
	router.Run("127.0.0.1:8080")
}

func AuthMidware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isLogin(c) {
			c.Redirect(http.StatusTemporaryRedirect, "/loginpage"+"?"+c.Request.URL.RawQuery)
			c.Abort()
			return
		}
		c.Next()
	}
}
