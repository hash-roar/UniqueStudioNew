package router

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"pastebin/database/handler"
	"pastebin/database/model"

	"github.com/gin-gonic/gin"
)

func Getpaste(c *gin.Context) {
	urlIndex := c.Param("url")
	post, rows := handler.Getpastedata(urlIndex)
	if rows == 0 {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	c.HTML(http.StatusOK, "content.html", gin.H{
		"syntax":  post.Syntax,
		"content": post.Content,
		"poster":  post.Poster,
	})
}

func Getpost(c *gin.Context) {
	var post = model.Pastecode{}
	if err := c.ShouldBind(&post); err != nil {
		println(err)
		panic(err)
	}
	data_json, _ := json.Marshal(post)
	data_md5 := md5.Sum(data_json)
	post.UrlIndex = hex.EncodeToString(data_md5[:])
	_, err := handler.Addpastedata(&post)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/pastes/"+hex.EncodeToString(data_md5[:]))
}
