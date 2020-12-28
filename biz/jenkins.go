package biz

import (
	"fmt"
	"net/http"
	"testmgmt/models"

	"github.com/gin-gonic/gin"
)

func StartSmoke(ctx *gin.Context) {
	project := ctx.PostForm("project")
	var message, status, failReason string
	var err error
	if err = RunSmoke(project); err == nil {
		status = "success"
		message = "测试成功"
	} else {
		status = "failure"
		message = fmt.Sprintf("测试失败")
		failReason = fmt.Sprintf("%s", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     status,
		"message":    message,
		"failReason": failReason,
	})
	return
}

func RunSmoke(project string) (err error) {
	var apiCases []DbApiCase
	models.Orm.Table("api_case").Where("project = ? and smoketest = ?", project, "yes").Find(&apiCases)
	LogHandle.Printf("apiCases Num: %q\n", len(apiCases))
	if len(apiCases) == 0 {
		err = fmt.Errorf("Not Found APICase has SmokeTest Tag in %s", project)
		LogHandle.Printf("Error: %s", err)
		return
	}
	tag := 0
	var errStr string
	for _, apicase := range apiCases {
		err = RunApiCase(apicase.Id, "")
		if err != nil {
			errStr = errStr + fmt.Sprintf("%s", err)
			tag++
		}
	}
	if tag > 0 {
		LogHandle.Printf("Error: %s", errStr)
		err = fmt.Errorf("Error: %s", errStr)
	}

	return
}
