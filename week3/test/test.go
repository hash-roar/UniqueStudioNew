package main

import (
	"fmt"
	"lform"

	_ "github.com/lib/pq"
)

type Pastecode struct {
	UrlIndex   string
	Content    string `form:"content"`
	Poster     string `form:"poster"`
	Syntax     string `form:"syntax"`
	Expiration string `form:"expiration"`
}

func main() {
	dsn := "user=dbuser password=zlf dbname=pastebin port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	engine, _ := lform.NewConnent("postgres", dsn)
	defer engine.Close()
	s := engine.NewSession()
	codes := Pastecode{UrlIndex: "index", Content: "test", Poster: "poster", Syntax: "test", Expiration: "h"}
	row, _ := s.Insert(&codes)
	fmt.Println(row)
}
