package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const clientId = "IamPleasantGoat"
const clientSecretKey = "IamtheheadofslowSheepvillage"

type userTokenInfo struct {
	userId     int    `json:"uid"`
	userToken  string `json:"token"`
	scope      string `json:"scope"`
	expireTime int    `json:"expireTime"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("D:\\working\\UniqueStudioNew\\week4\\web\\templets\\*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "thirdPartyPage.html", gin.H{
			"responseType": "code",
			"clientId":     clientId,
			"redirectUrl":  "http://localhost:8081/fetch-token",
			"scope":        "selfDailyReport",
		})
	})
	router.GET("/fetch-token", func(c *gin.Context) {
		code := c.Query("code")
		resp, err := http.PostForm("http://localhost:8080/oauth", url.Values{
			"code":         {code},
			"clientId":     {clientId},
			"clientSecret": {clientSecretKey},
		})
		if resp.StatusCode != 200 {
			c.String(404, "authorize fail")
			return
		}
		body, _ := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		userInfo := &userTokenInfo{}
		if err3 := json.Unmarshal(body, userInfo); err3 != nil {
			fmt.Println(err3)
		}

		c.SetCookie("client_token", userInfo.userToken, userInfo.expireTime*3600, "/", "localhost", false, true)
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/auth/write-page")
	})
	router.Any("/write-report", func(c *gin.Context) {
		uid := c.Query("uid")
		c.HTML(200, "writeReport.html", gin.H{"id": uid})
	})
	router.Any("/add-report", func(c *gin.Context) {

	})
	router.Run("127.0.0.1:8081")
}
