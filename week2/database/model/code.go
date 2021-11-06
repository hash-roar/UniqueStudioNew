package model

import (
	"unsafe"

	"gorm.io/gorm"
)

type Pastecode struct {
	gorm.Model
	UrlIndex   string
	Content    string `form:"content"`
	Poster     string `form:"poster"`
	Syntax     string `form:"syntax"`
	Expiration string `form:"expiration"`
}

func (data *Pastecode) Len() int64 {
	return int64(unsafe.Sizeof(*data))
}
