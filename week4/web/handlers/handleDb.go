package handlers

import (
	dbhandlers "authmanager/dbhandle"
	"authmanager/web/model"

	"gorm.io/gorm"
)

func init() {
	conn := dbhandlers.GetDbConn()
	conn.AutoMigrate(&model.User{})
}

func GetUserInfo(key interface{}) *model.User {
	var userInfo model.User
	var result *gorm.DB
	conn := dbhandlers.GetDbConn()
	switch k := key.(type) {
	case int:
		result = conn.Where("id = ?", k).Find(&userInfo)
	case string:
		result = conn.Where("name = ?", k).Find(&userInfo)
	default:
		return nil
	}
	if result.RowsAffected == 0 {
		return nil
	} else {
		return &userInfo
	}

}

func InsertUserInfo(user *model.User) (uint, error) {
	conn := dbhandlers.GetDbConn()
	result := conn.Create(user)
	return user.ID, result.Error
}
