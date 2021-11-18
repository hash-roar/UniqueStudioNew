package routers

import (
	"authmanager/web/handlers"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var clientAuthKey map[string]string

func init() {
	clientAuthKey = make(map[string]string)
	clientAuthKey["IamPleasantGoat"] = "IamtheheadofslowSheepvillage"
}

func Oauth2Page(c *gin.Context) {
	if !isLogin(c) {
		c.Redirect(http.StatusTemporaryRedirect, "/loginpage"+"?"+c.Request.URL.RawQuery+"&loginRedirectUrl=oauth-page")
	} else {
		clientId := c.Query("clientId")
		redirectUrl := c.Query("redirectUrl")
		scope := c.Query("scope")
		c.HTML(http.StatusOK, "authorization.html", gin.H{
			"clientId":    clientId,
			"redirectUrl": redirectUrl,
			"scope":       scope,
		})
	}
}

func SendAuthCode(c *gin.Context) {
	clientId := c.Query("clientId")
	redirectUrl := c.Query("redirectUrl")
	// scope := c.Query("scope")
	tokenStr, er := c.Cookie("token")
	if er != nil {
		c.Redirect(http.StatusMovedPermanently, "/loginpage")
	}
	claim := &Claims{}
	_, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println("parse token error")
	}
	uid := claim.Uid
	h := md5.New()
	n, err2 := h.Write([]byte(strconv.Itoa(uid) + ":" + clientId))
	if err2 != nil {
		log.Println("generate md5 code error", n)
	}
	hashCode := hex.EncodeToString(h.Sum(nil))
	handlers.AddClientTokenAuth(hashCode, strconv.Itoa(uid)+":"+clientId)
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl+"?"+"code="+hashCode)
}

func Authorize(c *gin.Context) {
	// clientId := c.PostForm("clientId")
	// clientSecretKey := c.PostForm("clientSecretKey")
	// grantType := c.PostForm("grantType")
	// code := c.PostForm("code")
	// redirectUri := c.PostForm("redirectUri")

}

func Oauth2(c *gin.Context) {
	clientId := c.PostForm("clientId")
	clientSec := c.PostForm("ClientSecret")
	code := c.PostForm("code")
	redirectUrl := c.PostForm("redirectUrl")

	clientSecKey, ok := clientAuthKey[clientId]
	if !ok {
		c.JSON(404, gin.H{
			"status":  1001,
			"message": "wrong clientId",
		})
		return
	}
	if clientSecKey != clientSec {
		c.JSON(404, gin.H{
			"status":  1002,
			"message": "wrong clientSecrect",
		})
		return
	}
	authInfo, err := handlers.GetClientAuthToken(code)
	if err != nil {
		c.JSON(404, gin.H{
			"status":  1003,
			"message": "no such auth info",
		})
		return
	}

}
