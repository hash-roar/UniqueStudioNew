package router

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"pastebin/cache"
	"pastebin/database/handler"
	"pastebin/database/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// func Getpaste(c *gin.Context) {
// 	urlIndex := c.Param("url")
// 	post, rows := handler.Getpastedata(urlIndex)
// 	if rows == 0 {
// 		c.Redirect(http.StatusMovedPermanently, "/")
// 	}
// 	c.HTML(http.StatusOK, "content.html", gin.H{
// 		"syntax":  post.Syntax,
// 		"content": post.Content,
// 		"poster":  post.Poster,
// 	})
// }

// func Getpost(c *gin.Context) {
// 	var post = model.Pastecode{}
// 	if err := c.ShouldBind(&post); err != nil {
// 		println(err)
// 		panic(err)
// 	}
// 	data_json, _ := json.Marshal(post)
// 	data_md5 := md5.Sum(data_json)
// 	post.UrlIndex = hex.EncodeToString(data_md5[:])
// 	_, err := handler.Addpastedata(&post)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	c.Redirect(http.StatusMovedPermanently, "/pastes/"+hex.EncodeToString(data_md5[:]))
// }

func getExpireTime(data *model.Pastecode) (expireTime time.Time) {
	expireTime = data.CreatedAt
	if data.Expiration == "n" {
		expireTime.AddDate(20, 0, 0)
	} else if strings.HasSuffix(data.Expiration, "d") {
		num, _ := strconv.ParseInt(data.Expiration[0:len(data.Expiration)-1], 10, 32)
		expireTime.AddDate(0, 0, int(num))
	} else {
		t, _ := time.ParseDuration(data.Expiration)
		expireTime = expireTime.Add(t)
	}
	return
}

func getData(key string) (*model.Pastecode, bool) { //获取pastebin数据
	var data *model.Pastecode
	fmt.Println(cache.LruCache)
	if temp, ok := cache.LruCache.Get(key); ok {
		tempData := temp.(*model.Pastecode)
		data = tempData

		expire_time := getExpireTime(data)
		if expire_time.Before(time.Now()) {
			cache.LruCache.Delete(data.UrlIndex)
		} else {
			data = tempData
			return data, true
		}

	}
	if temp, row := handler.Getpastedata(key); row > 0 {
		data = temp
		cache.LruCache.Add(data.UrlIndex, data)
		return data, true
	}
	return data, false
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
	// if data,ok:=getData(post.UrlIndex);ok {

	// }
	post.CreatedAt = time.Now()
	if _, err := handler.Addpastedata(&post); err != nil {
		fmt.Println(err)
	}
	if ok := handler.WriteCache(&post); ok {
		c.Redirect(http.StatusMovedPermanently, "/pastes/"+hex.EncodeToString(data_md5[:]))
	}

}

func GetPastePage(c *gin.Context) {
	urlIndex := c.Param("url")
	var post *model.Pastecode
	if temp, ok := getData(urlIndex); ok {
		post = temp
	} else {
		post = &model.Pastecode{}
	}
	c.HTML(http.StatusOK, "content.html", gin.H{
		"syntax":   post.Syntax,
		"content":  post.Content,
		"poster":   post.Poster,
		"urlindex": post.UrlIndex,
	})
}

func GetRaw(c *gin.Context) {
	urlIndex := c.Param("url")
	var post *model.Pastecode
	if temp, ok := getData(urlIndex); ok {
		post = temp
	} else {
		post = &model.Pastecode{}
	}
	c.String(http.StatusOK, "%s", post.Content)
}
