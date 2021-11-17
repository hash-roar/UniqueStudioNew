package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "login",
	})
}

func Login(c *gin.Context) {

}

func RigisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "rigister",
	})
}

func Rigister(c *gin.Context) {

}
