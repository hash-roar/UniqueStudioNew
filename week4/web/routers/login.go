package routers

import (
	"authmanager/web/handlers"
	"authmanager/web/model"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const SECRET_KEY = "ascr1292geaAxeca3122SA"

const maxAge = 60 * 60 * 24

var jwtKey = []byte(SECRET_KEY)

type Claims struct {
	Uid int
	jwt.StandardClaims
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":            "login",
		"loginRedirectUrl": c.Query("loginRedirectUrl"),
		"redirectParma":    c.Request.URL.RawQuery,
	})
}

func Login(c *gin.Context) {
	userName := c.PostForm("name")
	userPassword := c.PostForm("password")
	redirectUrl := c.PostForm("loginRedirectUrl")
	redirectParma := c.PostForm("redirectParma")
	if userName == "" || userPassword == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1001,
			"message": "name and password can not be null",
		})
		return
	}
	userInfo := handlers.GetUserInfo(userName)
	if userInfo == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1002,
			"message": "no such user",
		})
		return
	}
	h := md5.New()
	h.Write([]byte(userPassword))
	if hex.EncodeToString(h.Sum(nil)) != userInfo.Password {
		c.JSON(http.StatusOK, gin.H{
			"status":  1003,
			"message": "wrong password",
		})
		return
	}
	tokenStr := generateToken(int(userInfo.ID))
	c.SetCookie("token", tokenStr, maxAge, "/", "localhost", false, true)
	if redirectUrl != "" {
		c.Redirect(http.StatusMovedPermanently, redirectUrl+"?"+redirectParma)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
	})

}

func isLogin(c *gin.Context) bool {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		return false
	}
	if !checkToken(tokenStr) {
		return false
	}
	return true
}

func registerPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "register",
	})
}

func register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	h := md5.New()
	h.Write([]byte(password))
	userInfo := model.User{
		Name:     name,
		Password: hex.EncodeToString(h.Sum(nil)),
	}
	_, err := handlers.InsertUserInfo(&userInfo)
	if err != nil {
		log.Println("insert user info error")
		c.JSON(http.StatusOK, gin.H{
			"status":  1001,
			"message": "insert user info error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register successfully",
	})
}

func checkToken(tokenStr string) bool {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("parse token error")
		}
		return false
	}
	if !token.Valid {
		return false
	}
	return true
}

func generateToken(id int) string {
	claim := &Claims{
		Uid: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(maxAge * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("get token str error!")
		return ""
	}
	return tokenStr
}

func refreshToken(tokenStr string) string {
	claim := &Claims{}
	_, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println("refresh: parse token err")
	}
	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > 60*60*time.Second {
		return tokenStr
	}
	claim.ExpiresAt = time.Now().Add(maxAge * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStrAlter, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("get token str error!")
		return ""
	}
	return tokenStrAlter
}

func getUserInfoByToken(token string) *model.User {
	claim := &Claims{}
	_, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println(err)
	}
	id := claim.Uid
	userInfo := handlers.GetUserInfo(id)
	return userInfo
}
