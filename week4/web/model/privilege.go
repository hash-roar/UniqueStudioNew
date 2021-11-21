package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	Uid    int `gorm:"autoIncrement"`
	Rid    int
	IsRoot bool
}

type Role struct {
	gorm.Model
	Rid         int `gorm:"autoIncrement"`
	Name        string
	description string
}
type RolePriv struct {
	gorm.Model
	Rid int
	Pid int
}

type Privilege struct {
	gorm.Model
	Pid         int `gorm:"autoIncrement"`
	Name        string
	description string
	note        string
}
