package handlers

import (
	dbhandlers "authmanager/dbhandle"
	"authmanager/web/model"
	"log"

	"gorm.io/gorm"
)

func init() {
	conn := dbhandlers.GetDbConn()
	conn.AutoMigrate(&model.User{})
	conn.AutoMigrate(&model.UserRole{})
	conn.AutoMigrate(&model.Role{})
	conn.AutoMigrate(&model.RolePriv{})
	conn.AutoMigrate(&model.Privilege{})
	conn.AutoMigrate(&model.Report{})
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

func AddReport(report *model.Report) (uint, error) {
	conn := dbhandlers.GetDbConn()
	result := conn.Create(report)
	return report.ID, result.Error
}

func GetReport(uid int) ([]model.Report, error) {
	conn := dbhandlers.GetDbConn()
	var reports []model.Report

	result := conn.Where("uid = ?", uid).Find(&reports)
	return reports, result.Error
}

func GetTeamReport() ([]model.Report, error) {
	conn := dbhandlers.GetDbConn()
	var reports []model.Report
	result := conn.Find(&reports)
	return reports, result.Error
}

func GetPrivilege(uid int) ([]model.Privilege, error) {
	conn := dbhandlers.GetDbConn()
	var roleIds []int
	reTemp := conn.Table("user_roles").Where("uid = ?", uid).Select("rid").Find(&roleIds)
	if reTemp.Error != nil {
		return nil, reTemp.Error
	}
	var pids []int
	reTemp2 := conn.Table("role_privs").Where("rid IN ?", roleIds).Select("pid").Find(&pids)
	if reTemp2.Error != nil {
		return nil, reTemp2.Error
	}
	var userPrivilege []model.Privilege
	reTemp1 := conn.Table("privileges").Where("pid IN ?", pids).Find(&userPrivilege)
	if reTemp1.Error != nil {
		return nil, reTemp1.Error
	}
	return userPrivilege, nil
}

func GetPrivilegeRole(name string) *model.Role {
	conn := dbhandlers.GetDbConn()
	role := &model.Role{}
	result := conn.Table("roles").Where("name = ?", name).Find(role)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return role
}

func AddUserPrivilege(uid int, userRole *model.UserRole) {
	conn := dbhandlers.GetDbConn()
	conn.Create(userRole)
}

func DeleteUserPrivilege() {

}
