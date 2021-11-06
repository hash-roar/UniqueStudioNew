package router

import (
	"fmt"
	"net/http"
	"pastebin/cache"
	"pastebin/database/handler"
	"pastebin/database/model"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	cache.LruCache = cache.InitCache(1024*1024*20, func(key string, value cache.Value) {
		data := value.(*model.Pastecode)
		if data.Expiration == "n" {
			_, err := handler.Addpastedata(data)
			if err != nil {
				fmt.Println(err)
			}
		}
		t, _ := time.ParseDuration(data.Expiration)
		expire_time := data.CreatedAt.Add(t)
		data.DeletedAt.Time = expire_time
		if expire_time.After(time.Now()) {
			_, err := handler.Addpastedata(data)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
}

func GetIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "title",
	})
}
