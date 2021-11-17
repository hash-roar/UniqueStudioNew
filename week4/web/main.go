package main

import (
	"authmanager/web/routers"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	routers.Run()
}
