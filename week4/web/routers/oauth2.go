package routers

import (
	"authmanager/web/handlers"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var clientAuthKey map[string]string

type ThirdPartyClaims struct {
	Uid   int
	scope string
	jwt.StandardClaims
}

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
	auth := c.QueryArray("auth")
	var scope strings.Builder
	for i, v := range auth {
		if i != 0 {
			scope.WriteString(",")
		}
		scope.WriteString(v)
	}
	expireTime := c.Query("expireTime")
	log.Println(auth, expireTime)
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
	userInfo := make(map[string]string)
	userInfo["scope"] = scope.String()
	userInfo["expireTime"] = expireTime
	userInfo["uid"] = strconv.Itoa(uid)
	userInfo["clientId"] = clientId
	handlers.AddClientAuthToken(hashCode, userInfo)
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl+"?"+"code="+hashCode)
}

func Authorize(c *gin.Context) {
	// clientId := c.PostForm("clientId")
	// clientSecretKey := c.PostForm("clientSecretKey")
	// grantType := c.PostForm("grantType")
	// code := c.PostForm("code")
	// redirectUri := c.PostForm("redirectUri")

}

func generateOauthToken(claim *ThirdPartyClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("get token str error!")
		return ""
	}
	return tokenStr
}

func Oauth2(c *gin.Context) {
	clientId := c.PostForm("clientId")
	clientSec := c.PostForm("clientSecret")
	code := c.PostForm("code")
	redirectUrl := c.PostForm("redirectUrl")
	fmt.Println(redirectUrl)
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
	authInfo, err := handlers.GetClientAuthToken(code, "scope", "uid", "expireTime")
	if err != nil {
		c.JSON(404, gin.H{
			"status":  1003,
			"message": "no such auth info",
		})
		return
	}

	scope := authInfo[0]
	uid, _ := strconv.Atoi(authInfo[1])
	expireTime, _ := strconv.Atoi(authInfo[2])

	fmt.Println(uid)
	claim := &ThirdPartyClaims{Uid: uid,
		scope:          scope,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(expireTime) * time.Minute).Unix()}}
	c.JSON(200, gin.H{
		"uid":        uid,
		"token":      generateOauthToken(claim),
		"expierTime": expireTime,
	})

}
