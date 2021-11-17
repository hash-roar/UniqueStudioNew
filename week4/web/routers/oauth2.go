package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Oauth2Page(c *gin.Context) {
	if !isLogin(c) {
		c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
	} else {
		c.HTML(http.StatusOK, "authorization.html", gin.H{})
	}
}
func Oauth2() {

}
