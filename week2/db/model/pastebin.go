package model

type Pastecode struct {
	// ID         int
	Content    string `form:"content"`
	Poster     string `form:"poster"`
	Syntax     string `form:"syntax"`
	Expiration string `form:"expiration"`
	// CreateTime time.Time
}
