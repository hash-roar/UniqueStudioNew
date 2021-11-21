package model

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Uid       int    `json:"uid,omitempty"`
	StudyTime int    `json:"study_time,omitempty"`
	Report    string `gorm:"type:text" json:"report,omitempty"`
	Plan      string `gorm:"type:text" json:"plan,omitempty"`
	scope     string `json:"scope,omitempty"`
}
