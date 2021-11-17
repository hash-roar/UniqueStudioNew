package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
		return
	}
	if !checkToken(tokenStr) {
		c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
	}
	if tokenAlter := refreshToken(tokenStr); tokenAlter != tokenStr {
		c.SetCookie("token", tokenAlter, maxAge, "/", "localhost", false, true)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
