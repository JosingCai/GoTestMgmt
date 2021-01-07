package biz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"testmgmt/models"
	"time"
)

type TestCase struct {
	CaseNumber   string `gorm:"column:case_number" json:"case_number"`
	CaseName     string `gorm:"column:case_name" json:"case_name"`
	CaseType     string `gorm:"column:case_type" json:"case_type"`
	Priority     string `gorm:"column:priority" json:"priority"`
	PreCondition string `gorm:"column:pre_condition" json:"pre_condition"`
	TestRange    string `gorm:"column:test_range" json:"test_range"`
	TestSteps    string `gorm:"column:test_steps" json:"test_steps"`
	ExpectResult string `gorm:"column:expect_result" json:"expect_result"`
	Auto         string `gorm:"column:auto" json:"auto"`
	CaseID       string `gorm:"column:case_id" json:"case_id"`
	// FunDeveloper string `gorm:"column:fun_developer" json:"fun_developer"`
	// CaseDesigner string `gorm:"column:case_designer" json:"case_designer"`
	CaseExecutor string `gorm:"column:case_executor" json:"case_executor"`
	// TestTime     string `gorm:"column:test_time" json:"test_time"`
	TestResult string `gorm:"column:test_result" json:"test_result"`
	Module     string `gorm:"column:module" json:"module"`
	UpdatedAt  string `gorm:"column:updated_at" json:"updated_at"`
	// CaseId       string `gorm:"column:case_id" json:"case_id"`
	Project string `gorm:"column:project" json:"project"`
	// Remark       string `gorm:"column:remark" json:"remark"`
}

type XmindTestCase struct {
	Name                  string      `json:"name"`
	Version               int         `json:"version"`
	Summary               string      `json:"summary"`
	Preconditions         string      `json:"preconditions"`
	ExecutionType         int         `json:"execution_type"`
	Importance            int         `json:"importance"`
	EstimatedExecDuration int         `json:"estimated_exec_duration"`
	Status                int         `json:"status"`
	Result                int         `json:"result"`
	Steps                 []TestSteps `json:"steps"`
	Product               string      `json:"product"`
	Suite                 string      `json:"suite"`
}

type TestSteps struct {
	StepNumber      int    `json:"step_number"`
	Actions         string `json:"actions"`
	ExpectedResults string `json:"expectedresults"`
	ExecutionType   int    `json:"execution_type"`
	Result          int    `json:"result"`
}

func GetFileName(dirName, project string) (fileName string, err error) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}
	var allNames []string
	for _, file := range files {
		tmpName := file.Name()
		if strings.HasPrefix(tmpName, project) && strings.HasSuffix(tmpName, ".xmind") {
			rawName := dirName + "/" + tmpName
			allNames = append(allNames, rawName)
		}
	}

	if len(allNames) > 0 {
		// LogHandle.Printf("allFileNames: %s", allNames)
		fileName = allNames[len(allNames)-1]
	}

	if len(fileName) == 0 {
		err = fmt.Errorf("Not Found file [%s*.xmind] in directory[%s]", project, dirName)
		LogHandle.Printf("err: %v", err)
	}
	LogHandle.Printf("xmindFileName: %s", fileName)
	return
}

func GetJSON(id string) (err error) {
	StatusDef := map[int]string{1: "草稿", 2: "待评审", 3: "评审中", 4: "重做", 5: "废弃", 6: "特性", 7: "终稿"}
	PriorityDef := map[int]string{1: "高", 2: "中", 3: "低"}
	AutoDef := map[int]string{1: "否", 2: "是"}
	TestTypeDef := map[int]string{1: "功能测试", 2: "异常测试", 3: "场景测试"}
	var host Host
	models.Orm.Table("host").Where("id = ?", id).Find(&host)
	if len(host.Project) == 0 {
		err = fmt.Errorf("Not found related project id:%s", id)
		return
	}
	project := host.Project
	basePath := fmt.Sprintf("%s/testmgmt/test", BASEPATH)
	fileName, err := GetFileName(basePath, project)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}

	var ms string
	if strings.Contains(fileName, "_v") || strings.Contains(fileName, "_V") {
		items := strings.Split(fileName, "_")
		tmpName := items[1]
		ms = strings.Replace(tmpName, ".xmind", "", -1)
	}

	output, err := exec.Command("xmind2testcase", fileName, "-json").Output()
	if err != nil {
		LogHandle.Printf("output: %s", output)
		LogHandle.Printf("Error: %s", err)
	}

	jsonFileName := fileName[:len(fileName)-len(".xmind")] + ".json"
	LogHandle.Printf("jsonFileName: %s", jsonFileName)
	content, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		return
	}
	var xmindTestCases []XmindTestCase
	err = json.Unmarshal([]byte(content), &xmindTestCases)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}

	var caseNumberPrefix, sufixNum string
	for _, item := range xmindTestCases {
		var testcase, testCaseDb, testCaseDb2 TestCase
		modules := strings.Split(item.Suite, "-")
		testcase.CaseName = item.Name
		testcase.Auto = AutoDef[item.ExecutionType]
		testcase.TestResult = StatusDef[item.Status]
		testcase.CaseType = TestTypeDef[item.Importance]
		testcase.Priority = PriorityDef[item.Importance]
		testcase.PreCondition = item.Preconditions
		testcase.Project = project
		var stepStr, resultStr string
		for _, step := range item.Steps {
			tmpAction := fmt.Sprintf("%d. %s\n", step.StepNumber, step.Actions)
			tmpResult := fmt.Sprintf("%d. %s\n", step.StepNumber, step.ExpectedResults)
			stepStr = stepStr + tmpAction
			resultStr = resultStr + tmpResult
		}

		testcase.TestSteps = stepStr
		testcase.ExpectResult = resultStr
		if len(modules) > 1 {
			testcase.Module = modules[0]
			if len(ms) > 0 {
				caseNumberPrefix = modules[1] + "_" + ms
			} else {
				caseNumberPrefix = modules[1]
			}

		} else {
			testcase.CaseName = item.Name
			if len(ms) > 0 {
				caseNumberPrefix = item.Product + "_" + ms
			} else {
				caseNumberPrefix = item.Product + "_" + "other"
			}

		}

		chkStr := "%" + caseNumberPrefix + "%"
		models.Orm.Table("test_case").Where("project = ? AND case_number LIKE ?", project, chkStr).Find(&testCaseDb)
		if len(testCaseDb.CaseNumber) == 0 {
			sufixNum = strconv.Itoa(1)
		} else {
			tmps := strings.Split(testCaseDb.CaseNumber, "_")
			numStr := tmps[len(tmps)-1]
			s, _ := strconv.Atoi(numStr)
			if err != nil {
				LogHandle.Printf("Error: %s", err)
			}
			sufixNum = strconv.Itoa(s + 1)
		}

		testcase.CaseNumber = caseNumberPrefix + "_" + sufixNum
		// LogHandle.Printf("testcase: %+v", testcase)
		curTime := time.Now()
		testcase.UpdatedAt = curTime.Format(baseFormat)
		models.Orm.Table("test_case").Where("project = ? and case_number = ?", project, testcase.CaseNumber).Find(&testCaseDb2)
		if len(testCaseDb2.CaseNumber) == 0 {
			err = models.Orm.Table("test_case").Create(testcase).Error
			if err != nil {
				LogHandle.Printf("Error: %s", err)
			}
		} else {
			err = models.Orm.Table("test_case").Where("project = ? and case_number = ?", project, testcase.CaseNumber).Update(testcase).Error
			if err != nil {
				LogHandle.Printf("Error: %s", err)
			}
		}
	}

	return

}
