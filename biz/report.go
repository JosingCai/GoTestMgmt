package biz

import (
	"fmt"
	"testmgmt/models"
)

type AllAPI struct {
	AllCount           int     `gorm:"column:all_count" json:"all_count"`
	AutomatableCount   int     `gorm:"column:automatable_count" json:"automatable_count"`
	UnautomatableCount int     `gorm:"column:unautomatable_count" json:"unautomatable_count"`
	AutoTestCount      int     `gorm:"column:auto_test_count" json:"auto_test_count"`
	UntestCount        int     `gorm:"column:untest_count" json:"untest_count"`
	PassCount          int     `gorm:"column:pass_count" json:"pass_count"`
	FailCount          int     `gorm:"column:fail_count" json:"fail_count"`
	AutoPer            float64 `gorm:"column:auto_per" json:"auto_per"`
	PassPer            float64 `gorm:"column:pass_per" json:"pass_per"`
	FailPer            float64 `gorm:"column:fail_per" json:"fail_per"`
	Project            string  `gorm:"column:project" json:"project"`
}

type APIsCount struct {
	CaseId      string `gorm:"column:case_id" json:"case_id"`
	APIFunction string `gorm:"column:API_function" json:"API_function"`
	RunTimes    int    `gorm:"column:run_times" json:"run_times"`
	TestTimes   int    `gorm:"column:test_times" json:"test_times"`
	PassTimes   int    `gorm:"column:pass_times" json:"pass_times"`
	FailTimes   int    `gorm:"column:fail_times" json:"fail_times"`
	UntestTimes int    `gorm:"column:untest_times" json:"untest_times"`
	TestResult  string `gorm:"column:test_result" json:"test_result"`
	FailReason  string `gorm:"column:fail_reason" json:"fail_reason"`
	Project     string `gorm:"column:project" json:"project"`
}

func CountAPIs(id string) (err error) {
	var apiCount APIsCount
	var host Host
	models.Orm.Table("host").Where("id = ?", id).Find(&host)
	if len(host.Project) == 0 {
		err = fmt.Errorf("Not found related project id:%s", id)
		return
	}
	project := host.Project
	var apiTestDetails []ApiTestDetail
	var apiTestresults []ApiTestResult
	models.Orm.Table("api_test_result").Where("project = ?", project).Find(&apiTestresults)
	for _, apitestresult := range apiTestresults {
		models.Orm.Table("api_test_detail").Where("project = ? and case_id = ?", project, apitestresult.CaseID).Find(&apiTestDetails).Count(&apiCount.TestTimes)
		for k, apiTestDetail := range apiTestDetails {
			if apiTestDetail.TestResult == "failure" {
				apiCount.FailReason = apiCount.FailReason + apiTestDetail.FailReason
			}
			if k == 0 {
				apiCount.APIFunction = apiTestDetail.ApiFunction
			}
		}

		models.Orm.Table("api_test_detail").Where("project = ? and case_id = ? and test_result = ? or test_result = ?", project, apitestresult.CaseID, "success", "pass").Find(&apiTestDetails).Count(&apiCount.PassTimes)
		models.Orm.Table("api_test_detail").Where("project = ? and case_id = ? and test_result = ? or test_result = ?", project, apitestresult.CaseID, "failure", "fail").Find(&apiTestDetails).Count(&apiCount.FailTimes)
		models.Orm.Table("api_test_detail").Where("project = ? and case_id = ? and test_result = ?", project, apitestresult.CaseID, "untest").Find(&apiTestDetails).Count(&apiCount.UntestTimes)
		apiCount.RunTimes = apiCount.PassTimes + apiCount.FailTimes
		apiCount.TestTimes = apiCount.PassTimes + apiCount.FailTimes + apiCount.UntestTimes
		if apiCount.PassTimes > 0 {
			apiCount.TestResult = "pass"
		} else {
			apiCount.TestResult = "fail"
		}
		apiCount.CaseId = apitestresult.CaseID
		apiCount.Project = project
		err = models.Orm.Table("case_test_count").Create(apiCount).Error
		if err != nil {
			LogHandle.Printf("err: %s", err)
		}
	}

	return

}

func CountAllAPI(id string) (err error) {
	var host Host
	models.Orm.Table("host").Where("id = ?", id).Find(&host)
	if len(host.Project) == 0 {
		err = fmt.Errorf("Not found related project id:%s", id)
		return
	}
	project := host.Project

	var allAPI AllAPI
	var apitestresults []ApiTestResult
	var apicases []APICase
	models.Orm.Table("api_case").Where("project = ?", project).Find(&apicases).Count(&allAPI.AllCount)
	models.Orm.Table("api_case").Where("project = ? and runNum >= ?", project, 1).Find(&apicases).Count(&allAPI.AutomatableCount)
	models.Orm.Table("api_case").Where("project = ? and runNum = ?", project, 0).Find(&apicases).Count(&allAPI.UnautomatableCount)
	models.Orm.Table("api_test_result").Where("project = ? and result = ? or result = ?", project, "failure", "fail").Find(&apitestresults).Count(&allAPI.FailCount)
	models.Orm.Table("api_test_result").Where("project = ? and result = ? or result = ?", project, "success", "pass").Find(&apitestresults).Count(&allAPI.PassCount)
	allAPI.AutoTestCount = allAPI.PassCount + allAPI.FailCount
	allAPI.UntestCount = allAPI.AllCount - allAPI.PassCount - allAPI.FailCount
	if allAPI.AllCount > 0 {
		allAPI.AutoPer = float64(allAPI.AutomatableCount) / float64(allAPI.AllCount)
	}

	if allAPI.AutomatableCount > 0 {
		allAPI.PassPer = float64(allAPI.PassCount) / float64(allAPI.AutomatableCount)
		allAPI.FailPer = float64(allAPI.FailCount) / float64(allAPI.AutomatableCount)
	}

	allAPI.Project = project
	// LogHandle.Printf("allAPI: %+v", allAPI)
	err = models.Orm.Table("api_sum_up").Create(allAPI).Error
	if err != nil {
		LogHandle.Printf("err: %s", err)
	}
	return
}
