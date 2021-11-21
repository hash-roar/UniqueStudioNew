package routers

import (
	"authmanager/web/handlers"
	"authmanager/web/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
		return
	}
	uid := int(getUserInfoByToken(tokenStr).ID)
	if !checkUserPrivilege(uid, []string{SelfDailyReportReadOnly}) {
		c.String(404, "permission denied")
		return
	}
	reports, err := handlers.GetReport(uid)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "index.html", reports)
}

func GetList(c *gin.Context) {
	tokenStr, _ := c.Cookie("token")
	// if err != nil {
	// 	c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
	// 	return
	// }
	uid := int(getUserInfoByToken(tokenStr).ID)
	if !checkUserPrivilege(uid, []string{SelfDailyReportReadOnly}) {
		c.String(404, "permission denied")
		return
	}

}

func TeamReport(c *gin.Context) {
	tokenStr, _ := c.Cookie("token")
	uid := int(getUserInfoByToken(tokenStr).ID)
	if !checkUserPrivilege(uid, []string{TeamDailyReportReadOnly}) {
		c.String(404, "permission denied")
		return
	}
	reports, err := handlers.GetTeamReport()
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "teamReport.html", reports)
}

func WriteReportPage(c *gin.Context) {
	c.HTML(http.StatusOK, "writeReport.html", gin.H{})
}

func AddReport(c *gin.Context) {
	report := &model.Report{}
	report.StudyTime, _ = strconv.Atoi(c.PostForm("time"))
	report.Report = c.PostForm("report")
	report.Plan = c.PostForm("plan")
	tokenStr, _ := c.Cookie("token")
	uid := int(getUserInfoByToken(tokenStr).ID)
	report.Uid = uid
	if !checkUserPrivilege(uid, []string{SelfDailyReport}) {
		c.String(404, "permission denied")
		return
	}
	if _, err := handlers.AddReport(report); err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "write error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
