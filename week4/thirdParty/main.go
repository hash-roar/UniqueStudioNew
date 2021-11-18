package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const clientId = "IamPleasantGoat"
const clientSecretKey = "IamtheheadofslowSheepvillage"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("D:\\working\\UniqueStudioNew\\week4\\web\\templets\\*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "thirdPartyPage.html", gin.H{
			"responseType": "code",
			"clientId":     clientId,
			"redirectUri":  "/fetchToken",
			"scope":        "dailyReport",
		})
	})
	router.GET("/fetchToken", func(c *gin.Context) {

	})
	router.Run("127.0.0.1:8081")
}
