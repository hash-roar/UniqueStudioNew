package lform_test

import (
	"fmt"
	"lform"
)

type User struct {
	Id   int `lfrom:"PRIMARY KEY"`
	Name string
}

var dsn string = "user=dbuser password=zlf dbname=pastebin port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func Connect_test() {
	engine1, err := lform.NewConnent("postgres", dsn)
	if err != nil {
		fmt.Println("connect_test postgres error")
		fmt.Println(err)
	}
	defer engine1.Close()
	engine2, err2 := lform.NewConnent("mysql", dsn)
	if err != nil {
		fmt.Println("connect_test mysql error")
		fmt.Println(err2)
	}
	defer engine2.Close()
}

func Session_test() {
	engine, _ := lform.NewConnent("postgres", dsn)
	defer engine.Close()
	s := engine.NewSession()
	if s == nil {
		fmt.Println("new session error")
	}
}
