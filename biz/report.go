package biz

import (
	"fmt"
	"testmgmt/models"
)

type SimpleCase struct {
	Module     string `gorm:"column:module" json:"module"`
	TestResult string `gorm:"column:test_result" json:"test_result"`
	Project    string `gorm:"column:project" json:"project"`
}

type CasesCount struct {
	Module          string  `gorm:"column:module" json:"module"`
	Allcase         int     `gorm:"column:allcase" json:"allcase"`
	UnincludeUntest int     `gorm:"column:uninclude_untest" json:"uninclude_untest"`
	Pass            int     `gorm:"column:pass" json:"pass"`
	Fail            int     `gorm:"column:fail" json:"fail"`
	PassPer         float64 `gorm:"column:pass_per" json:"pass_per"`
	Project         string  `gorm:"column:project" json:"project"`
}

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

func CountCases(id string) (err error) {
	var casesCounts []CasesCount
	var caseCount CasesCount
	var tmpCases []SimpleCase
	var host Host
	models.Orm.Table("host").Where("id = ?", id).Find(&host)
	if len(host.Project) == 0 {
		err = fmt.Errorf("Not found related project id:%s", id)
		return
	}
	project := host.Project

	// var testcase TestCase
	//
	var modules []string

	models.Orm.Table("test_case").Model(&TestCase{}).Where("project = ?", project).Pluck("distinct(module)", &modules)
	// LogHandle.Printf("modules: %+v", modules)
	if len(modules) == 0 {
		LogHandle.Printf("Not Found: project[%s] case ", project)
		return
	}

	var untest, deprecated *int
	for _, v := range modules {
		models.Orm.Table("test_case").Where("project = ? and module = ?", project, v).Find(&tmpCases).Count(&caseCount.Allcase)
		models.Orm.Table("test_case").Where("project = ? and module = ? and test_result = ?", project, v, "pass").Find(&tmpCases).Count(&caseCount.Pass)
		models.Orm.Table("test_case").Where("project = ? and module = ? and test_result = ?", project, v, "fail").Find(&tmpCases).Count(&caseCount.Fail)
		models.Orm.Table("test_case").Where("project = ? and module = ? and test_result = ?", project, v, "untest").Find(&tmpCases).Count(&untest)
		models.Orm.Table("test_case").Where("project = ? and module = ? and test_result = ?", project, v, "deprecated").Find(&tmpCases).Count(&deprecated)
		caseCount.Project = project
		caseCount.Module = v
		caseCount.PassPer = float64(caseCount.Pass) / float64(caseCount.Allcase-*untest-*deprecated)
		casesCounts = append(casesCounts, caseCount)
	}
	caseCount.Module = "全部用例"
	models.Orm.Table("test_case").Where("project = ?", project).Find(&tmpCases).Count(&caseCount.Allcase)
	models.Orm.Table("test_case").Where("project = ? and test_result = ?", project, "pass").Find(&tmpCases).Count(&caseCount.Pass)
	models.Orm.Table("test_case").Where("project = ? and test_result = ?", project, "fail").Find(&tmpCases).Count(&caseCount.Fail)
	models.Orm.Table("test_case").Where("project = ? and test_result = ?", project, "untest").Find(&tmpCases).Count(&untest)
	models.Orm.Table("test_case").Where("project = ? and test_result = ?", project, "deprecated").Find(&tmpCases).Count(&deprecated)
	caseCount.PassPer = float64(caseCount.Pass) / float64(caseCount.Allcase-*untest-*deprecated)
	casesCounts = append(casesCounts, caseCount)

	for _, v := range casesCounts {
		err = models.Orm.Table("testcase_count").Create(v).Error
		if err != nil {
			LogHandle.Printf("err: %s", err)
		}
	}

	return

}

func CreateReport(id string) (err error) {
	return
}
