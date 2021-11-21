package routers

import (
	"authmanager/web/handlers"
	"authmanager/web/model"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	SelfDailyReport         = "self:dailyReport"
	SelfDailyReportReadOnly = "self:dailyReport:readonly"
	TeamDailyReport         = "team:dailyReport"
	TeamDailyReportReadOnly = "team:dailyReport:readonly"
)

func AdminPage(c *gin.Context) {
	tokenStr, _ := c.Cookie("token")
	userInfo := getUserInfoByToken(tokenStr)
	if !userInfo.Admin {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "not admin user",
		})
		return
	}
	c.HTML(200, "admin.html", gin.H{})
}

func getPrivilege(uid int) ([]string, error) {
	result, err := handlers.GetPrivilege(uid)
	if err != nil {
		log.Println("get privilege error")
		return nil, err
	}
	if result == nil {
		return nil, err
	}
	resultStr := make([]string, 0)
	for _, v := range result {
		resultStr = append(resultStr, v.Name)
	}
	return resultStr, err
}

func checkPrivilege(priv []string, expect []string) bool {
	if len(priv) == 0 {
		return false
	}
	for _, expectPriv := range expect {
		var flag bool
		for _, hasPriv := range priv {
			if hasPriv == expectPriv {
				flag = true
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}

func checkUserPrivilege(uid int, expectPriv []string) bool {
	userPriv, err := getPrivilege(uid)
	if err != nil {
		log.Println(err)
		return false
	}
	return checkPrivilege(userPriv, expectPriv)
}

func AltPrivilege(c *gin.Context) {
	name := c.PostForm("name")
	roleName := c.PostForm("role")
	userInfo := handlers.GetUserInfo(name)
	roleInfo := handlers.GetPrivilegeRole(roleName)
	userRole := &model.UserRole{Uid: int(userInfo.ID), Rid: roleInfo.Rid}
	handlers.AddUserPrivilege(int(userInfo.ID), userRole)
}
