package biz

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"testmgmt/models"
)

const baseFormat = "2006-01-02 15:04:05"

type Response struct {
	Status        string                 `json:"status,omitempty"`
	Message       string                 `json:"message,omitempty"`
	Content       map[string]interface{} `json:"content,omitempty"`
	IsSuccess     bool                   `json:"isSuccess,omitempty"`
	ResultMessage string                 `json:"resultMessage,omitempty"`
	ResultObject  map[string]interface{} `json:"resultObject,omitempty"`
}

type Host struct {
	// gorm.Model
	Project   string `gorm:"column:project" json:"project"`
	Ip        string `gorm:"column:ip" json:"ip"`
	Protocol  string `gorm:"column:protocol" json:"protocol"`
	Auth      string `gorm:"column:auth" json:"auth"`
	Prepath   string `gorm:"column:prepath" json:"prepath"`
	Threading string `gorm:"column:threading" json:"threading"`
	Usermode  string `gorm:"column:usermode" json:"usermode"`
	Dbconfig  string `gorm:"column:dbconfig" json:"dbconfig"`
	Token     string `gorm:"column:token" json:"token"`
	Testmode  string `gorm:"column:testmode" json:"testmode"`
}

// type DbApiTestDetail struct {
// 	Id string `gorm:"column:id" json:"id"`
// 	ApiTestDetail
// }

type ApiTestDetail struct {
	CaseID      string `gorm:"column:case_id" json:"case_id"`
	ApiFunction string `gorm:"column:APIFunction" json:"api_function"`
	Url         string `gorm:"column:url" json:"url"`
	Body        string `gorm:"column:body" json:"body"`
	Response    string `gorm:"column:response" json:"response"`
	TestResult  string `gorm:"column:test_result" json:"test_result"`
	FailReason  string `gorm:"column:fail_reason" json:"fail_reason"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
	Project     string `gorm:"column:project" json:"project"`
}

type ApiTestData struct {
	Id             int    `gorm:"column:id" json:"id"`
	DataDesc       string `gorm:"column:data_desc" json:"data_desc"`
	ApiFunction    string `gorm:"column:apiFunction" json:"apiFunction"`
	CaseID         string `gorm:"column:case_id" json:"case_id"`
	Module         string `gorm:"column:module" json:"module"`
	UrlQuery       string `gorm:"column:urlQuery" json:"urlQuery"`
	Body           string `gorm:"column:body" json:"body"`
	ExpectedResult string `gorm:"column:expected_result" json:"expected_result"`
	ActualResult   string `gorm:"column:actual_result" json:"actual_result"`
	Result         string `gorm:"column:result" json:"result"`
	FailReason     string `gorm:"column:fail_reason" json:"fail_reason"`
	Response       string `gorm:"column:response" json:"response"`
	Project        string `gorm:"column:project" json:"project"`
}

type DbApiTestResult struct {
	Id string `gorm:"column:id" json:"id"`
	ApiTestResult
}

type ApiTestResult struct {
	CaseID      string `gorm:"column:case_id" json:"case_id"`
	RequestVars string `gorm:"column:requestVars" json:"request_vars"`
	Result      string `gorm:"column:result" json:"result"`
	OutVars     string `gorm:"column:outVars" json:"out_vars"`
	UpdatedAt   string `gorm:"column:updated_at" json:"updated_at"`
	Project     string `gorm:"column:project" json:"project"`
}

func GetRandomStr(length int) (ranStr string) {
	baseStr := "ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789"
	bytes := []byte(baseStr)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	ranStr = string(result)
	return
}

type API struct {
	Id          string `gorm:"column:id" json:"id"`
	CaseID      string `gorm:"column:case_id" json:"case_id"`
	HttpMethod  string `gorm:"column:httpMethod" json:"httpMethod"`
	ApiFunction string `gorm:"column:apiFunction" json:"apiFunction"`
	Path        string `gorm:"column:path" json:"path"`
	Header      string `gorm:"column:header" json:"header"`
	QueryParam  string `gorm:"column:queryParameter" json:"queryParameter"`
	PathVar     string `gorm:"column:pathVariable" json:"pathVariable"`
	Body        string `gorm:"column:body" json:"body"`
	Response    string `gorm:"column:response" json:"response"`
	Project     string `gorm:"column:project" json:"project"`
}

func GetAbDef() (inDef map[string]string) {
	inDef = make(map[string]string)
	inDef["intAb"] = "65536,-1"
	inDef["intNor"] = "10,1"
	inDef["strAb"] = GetRandomStr(256) + "," + GetRandomStr(65)
	inDef["strNor"] = GetRandomStr(255) + "," + GetRandomStr(8)
	inDef["arrAb"] = "65536,-1"
	inDef["arrNor"] = "1,2"
	inDef["objAb"] = "{}"
	inDef["bool"] = "true,false"
	return
}

type ComVar struct {
	Name    string `gorm:"column:name" json:"name"`
	Value   string `gorm:"column:value" json:"value"`
	Project string `gorm:"column:project" json:"project"`
}

type APICase struct {
	// gorm.Model
	CaseID     string `gorm:"column:case_id" json:"case_id"`
	Module     string `gorm:"column:module" json:"module"`
	RunNum     int    `gorm:"column:runNum" json:"run_num"`
	BeforeCase string `gorm:"column:beforeCase" json:"before_case"`
	AfterCase  string `gorm:"column:afterCase" json:"after_case"`
	OutVars    string `gorm:"column:outVars" json:"out_vars"`
	ChkVars    string `gorm:"column:chkVars" json:"chk_vars"`
	ParamDef   string `gorm:"column:param_def" json:"param_def"`
	Raw        string `gorm:"column:raw" json:"raw"`
	CreatedAt  string `gorm:"column:created_at" json:"created_at"`
	Project    string `gorm:"column:project" json:"project"`
}

func (apiCase APICase) IsRun() (err error) {
	if apiCase.RunNum == 0 {
		err = fmt.Errorf("%s has %d runNum test ... ", apiCase.CaseID, apiCase.RunNum)
		LogHandle.Println(err)
		return
	}
	return

}

func (apiCase APICase) GetLastValue(mapped map[string]string, data map[string]interface{}) (map[string]string, map[string]interface{}) {
	if data == nil {
		return mapped, nil
	}
	var mKeys, dKeys []string
	for k := range mapped {
		mKeys = append(mKeys, k)
	}

	var tmpMap map[string]interface{}
	tmpMap = make(map[string]interface{})
	var depDefs map[string]string
	depDefs = make(map[string]string)
	i := 0
	for _, v := range mKeys {
		if strings.Contains(mapped[v], "-") || strings.Contains(mapped[v], "*") {
			depDefs[v] = mapped[v]
			dKeys = append(dKeys, v)
			i++
		} else {
			varType := fmt.Sprintf("%T", data[mapped[v]])
			var strValue string
			switch varType {
			case "float64":
				tmpVar := data[mapped[v]].(float64)
				strValue = strconv.FormatFloat(tmpVar, 'f', 0, 64)
			case "string":
				strValue = data[mapped[v]].(string)
			case "bool":
				strValue = strconv.FormatBool(data[mapped[v]].(bool))
			default:
				LogHandle.Printf("A Problem had Occured at Get Out Var: %s", v)
			}
			mapped[v] = strValue
		}
	}
	if i == 0 {
		LogHandle.Printf("mapped: %+v", mapped)
		return mapped, nil
	}

	j := 0
	for _, v := range dKeys {
		items := strings.SplitN(mapped[v], "-", 2)
		mapped[v] = items[1]
		tmpMap = data[items[0]].(map[string]interface{})
		j++
		if j == len(dKeys) {
			return apiCase.GetLastValue(mapped, tmpMap)
		}

	}
	LogHandle.Printf("mapped: %+v", mapped)
	return mapped, nil
}

func (apiCase APICase) SaveTestResult(requestData, response map[string]interface{}) (err error) {
	var testResult ApiTestResult
	var dbResult DbApiTestResult
	testResult.CaseID = apiCase.CaseID
	reqStr, err := json.Marshal(requestData)
	testResult.RequestVars = string(reqStr)
	if value, ok := response["status"]; ok {
		testResult.Result = value.(string)
	}

	if len(apiCase.OutVars) > 0 {
		var mapOutVar map[string]string
		mapOutVar = make(map[string]string)
		err = json.Unmarshal([]byte(apiCase.OutVars), &mapOutVar)
		retOut, _ := apiCase.GetLastValue(mapOutVar, response)
		var outByte []byte
		outByte, err = json.Marshal(retOut)
		testResult.OutVars = string(outByte)
		testResult.OutVars = string(outByte)
	}

	testResult.Project = apiCase.Project
	curTime := time.Now()
	testResult.UpdatedAt = curTime.Format(baseFormat)
	models.Orm.Table("api_test_result").Where("project = ? and case_id = ?", apiCase.Project, apiCase.CaseID).Find(&dbResult)

	if len(dbResult.CaseID) == 0 {
		err = models.Orm.Table("api_test_result").Create(testResult).Error
	} else {
		err = models.Orm.Table("api_test_result").Where("id = ?", dbResult.Id).Update(testResult).Error
	}
	return
}

func (apiCase APICase) SaveDetailResult(url, apiFunction string, requestData map[string]interface{}, response map[string]interface{}) (err error) {
	var testDetail ApiTestDetail
	if value, ok := response["status"]; ok {
		testDetail.TestResult = value.(string)
		if value.(string) == "failure" {
			testDetail.FailReason = response["message"].(string)
		} else {
			testDetail.TestResult = "success"
		}
	} else if value, ok := response["isSuccess"]; ok {
		if !value.(bool) {
			testDetail.FailReason = response["resultMessage"].(string)
			testDetail.TestResult = "failure"
		} else {
			testDetail.TestResult = "success"
		}
	}

	testDetail.CaseID = apiCase.CaseID
	testDetail.ApiFunction = apiFunction
	testDetail.Url = url

	reqByte, err := json.Marshal(requestData)
	testDetail.Body = string(reqByte)

	resByte, err := json.Marshal(response)
	testDetail.Response = string(resByte)

	testDetail.Project = apiCase.Project
	curTime := time.Now()
	testDetail.CreatedAt = curTime.Format(baseFormat)

	err = models.Orm.Table("api_test_detail").Create(testDetail).Error
	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}

	return
}

func (apiCase APICase) GetAPI() (api API, err error) {
	models.Orm.Table("api_detail").Where("project = ? and case_id = ? ", apiCase.Project, apiCase.CaseID).Find(&api)
	if len(api.CaseID) == 0 {
		err = fmt.Errorf("Not Found %s API Detail", apiCase.CaseID)
		LogHandle.Println(err)
		return
	}
	return
}

func ChkUniVar(name, project string) (b bool) {
	var comVar ComVar
	uniVar := "uniVar"
	models.Orm.Table("common_variable").Where("project = ? and name = ? ", project, uniVar).Find(&comVar)
	if len(comVar.Value) > 0 {
		uniVars := strings.Split(comVar.Value, ",")
		for _, v := range uniVars {
			if strings.TrimSpace(v) == strings.TrimSpace(name) {
				b = true
				return
			}
		}
	}
	return
}

func (api API) GetFormatDepVars(depOutVars map[string]string) (retOutVars map[string]string, err error) {
	var allVar map[string]interface{}
	allVar = make(map[string]interface{})
	retOutVars = make(map[string]string)
	if len(api.PathVar) > 0 {
		err = json.Unmarshal([]byte(api.PathVar), &allVar)
		if err != nil {
			LogHandle.Printf("Error: %q", err)
			return
		}
	}

	if len(api.QueryParam) > 0 {
		err = json.Unmarshal([]byte(api.QueryParam), &allVar)
		if err != nil {
			LogHandle.Printf("Error: %q", err)
			return
		}
	}

	if len(api.Body) > 0 {
		err = json.Unmarshal([]byte(api.Body), &allVar)
		if err != nil {
			LogHandle.Printf("Error: %q", err)
			return
		}

	}
	retOutVars = depOutVars
	var host Host
	abDef := GetAbDef()
	models.Orm.Table("host").Where("project = ?", api.Project).Find(&host)
	for k, v := range allVar {
		if host.Testmode == "abnormal" {
			if value, ok := depOutVars[k]; !ok {
				if v == "integer" || v == "int" {
					retOutVars[k] = abDef["intAb"]
				} else if v == "string" {
					retOutVars[k] = abDef["strAb"]
				} else if v == "array" {
					retOutVars[k] = abDef["arrAb"]
				} else if v == "bool" {
					retOutVars[k] = abDef["bool"]
				} else {
					retOutVars[k] = abDef["strAb"]
				}

			} else {
				retOutVars[k] = value
			}
		} else if host.Testmode == "normal" {
			if value, ok := depOutVars[k]; !ok {
				if v == "integer" || v == "int" {
					retOutVars[k] = abDef["intNor"]
				} else if v == "string" {
					retOutVars[k] = abDef["strNor"]
				} else if v == "array" {
					retOutVars[k] = abDef["arrNor"]
				} else if v == "bool" {
					retOutVars[k] = abDef["bool"]
				} else {
					retOutVars[k] = abDef["strNor"]
				}
			} else {
				retOutVars[k] = value
			}
		}
	}
	LogHandle.Printf("retOutVars: %q", retOutVars)
	return
}

func (apiCase APICase) GetDepVars() (depOutVars map[string]string, err error) {
	var depCases []string
	depOutVars = make(map[string]string)
	if len(apiCase.ParamDef) > 0 {
		pds := strings.Split(apiCase.ParamDef, ",")
		depCases = append(depCases, pds...)
	}
	if len(apiCase.BeforeCase) > 0 {
		bcs := strings.Split(apiCase.BeforeCase, ",")
		depCases = append(depCases, bcs...)
	}
	depCases = append(depCases, apiCase.CaseID)
	if len(depCases) > 1 {
		LogHandle.Printf("DepCases: %q", depCases)
	}

	var testResult ApiTestResult
	var comVar ComVar
	for _, dep := range depCases {
		if dep == apiCase.CaseID {
			continue
		}
		models.Orm.Table("api_test_result").Where("project = ? and case_id = ?", apiCase.Project, dep).Find(&testResult)
		if len(testResult.OutVars) > 0 {
			LogHandle.Printf("testResult: %q", testResult)
			err = json.Unmarshal([]byte(testResult.OutVars), &depOutVars)
			if err != nil {
				LogHandle.Printf("Error: %q", err)
				return
			}
		}
	}
	for _, dep := range depCases {
		models.Orm.Table("common_variable").Where("project = ? and name = ?", apiCase.Project, dep).Find(&comVar)
		if len(comVar.Value) > 0 {
			// LogHandle.Printf("comVar: %q", comVar)
			err = json.Unmarshal([]byte(comVar.Value), &depOutVars)
			if err != nil {
				LogHandle.Printf("Error: %q", err)
				return
			}
		}

	}

	if len(depOutVars) > 0 {
		LogHandle.Printf("DepOutVars: %q", depOutVars)
	}

	return
}

func (apiCase APICase) GetRawUrl() (url string) {
	var host Host
	models.Orm.Table("host").Where("project = ?", apiCase.Project).Find(&host)
	paths := strings.Split(apiCase.CaseID, "_")
	url = host.Protocol + "://" + host.Ip + host.Prepath + paths[1]
	LogHandle.Printf("url: %s\n", url)
	return
}

func GetUrl(rawUrl string, depOutVars map[string]string) (urls []string, err error) {
	pathVarsReg := regexp.MustCompile(`{[[:alpha:]]+}`)
	var pathVars []string
	if !strings.Contains(rawUrl, "{") {
		urls = append(urls, rawUrl)
	} else {
		pathVars = pathVarsReg.FindAllString(rawUrl, -1)
		for _, v := range pathVars {
			str1 := v[1 : len(v)-1]
			if value, ok := depOutVars[str1]; ok {
				if !strings.Contains(value, ",") {
					url := strings.Replace(rawUrl, v, value, -1)
					urls = append(urls, url)
				} else {
					strList := strings.Split(value, ",")
					for _, sv := range strList {
						url := strings.Replace(rawUrl, v, sv, -1)
						urls = append(urls, url)
					}
				}
			} else {
				err = fmt.Errorf("Get URL: Not Found [%q] value in DepOutVars", v)
				return
			}

		}
	}
	LogHandle.Printf("URLs: %q", urls)
	return

}

func (api API) GetHeader() (header map[string]string, err error) {
	header = make(map[string]string)
	header["Content-Type"] = "application/json"
	if len(api.Header) == 0 {
		return
	}
	var mapHeader map[string]string
	mapHeader = make(map[string]string)
	err = json.Unmarshal([]byte(api.Header), &mapHeader)
	if err != nil {
		return
	}
	var hKeys []string
	for k := range mapHeader {
		hKeys = append(hKeys, k)
	}
	if len(hKeys) == 0 {
		LogHandle.Printf("Header: %q", header)
		return
	}
	var host Host
	models.Orm.Table("host").Where("project = ?", api.Project).Find(&host)
	if len(host.Token) == 0 {
		err = errors.New("Get Header: Not Found token value in Host")
		return
	}
	for _, v := range hKeys {
		header[v] = host.Token
	}
	LogHandle.Printf("Header: %q", header)
	return
}

func (api API) GetQuery(depOutVars map[string]string) (querys []string, err error) {
	if len(api.QueryParam) == 0 {
		return
	}
	var mapQuery map[string]string
	mapQuery = make(map[string]string)
	err = json.Unmarshal([]byte(api.QueryParam), &mapQuery)
	if err != nil {
		return
	}
	var qKeys []string
	for k := range mapQuery {
		qKeys = append(qKeys, k)
	}
	var comVar ComVar
	for _, v := range qKeys {
		if value, ok := depOutVars[v]; ok {
			if !strings.Contains(value, ",") {
				queryStr := fmt.Sprintf("%s=%s", v, value)
				querys = append(querys, queryStr)
			} else {
				strList := strings.Split(value, ",")
				for _, sv := range strList {
					queryStr := fmt.Sprintf("%s=%s", v, sv)
					querys = append(querys, queryStr)
				}
			}
		} else {
			models.Orm.Table("common_variable").Where("project = ? and name = ?", api.Project, v).Find(&comVar)
			if len(comVar.Value) == 0 {
				err = fmt.Errorf("Get Query: Not Found [%q] value in DepOutVars and ComVars", v)
				return
			} else {
				if !strings.Contains(comVar.Value, ",") {
					queryStr := fmt.Sprintf("%s=%s", v, comVar.Value)
					querys = append(querys, queryStr)
				} else {
					strList := strings.Split(value, ",")
					for _, sv := range strList {
						queryStr := fmt.Sprintf("%s=%s", v, sv)
						querys = append(querys, queryStr)
					}
				}
			}

		}

	}

	if len(querys) > 0 {
		LogHandle.Printf("Querys: %q", querys)
	}

	return
}

func (api API) GetBody(depOutVars map[string]string) (bodys []map[string]interface{}, err error) {
	if len(api.Body) == 0 {
		return
	}
	var mapBody map[string]interface{}
	mapBody = make(map[string]interface{})
	err = json.Unmarshal([]byte(api.Body), &mapBody)
	if err != nil {
		return
	}

	var bKeys []string
	var multiVars []string
	for k := range mapBody {
		bKeys = append(bKeys, k)
	}
	var intValue int
	var boolValue bool

	for _, v := range bKeys {
		if value, ok := depOutVars[v]; ok {
			if mapBody[v] == "array" {
				strList := strings.Split(value, ",")
				var intList []int
				for _, sv := range strList {
					intValue, err := strconv.Atoi(sv)
					if err == nil {
						intList = append(intList, intValue)
					}
				}
				if len(intList) > 0 {
					mapBody[v] = intList
				} else {
					mapBody[v] = strList
				}
			} else {
				if strings.Contains(value, ",") {
					multiVars = append(multiVars, v)
				} else {
					if mapBody[v] == "int" || mapBody[v] == "integer" {
						intValue, err = strconv.Atoi(value)
						if err != nil {
							return
						}
						mapBody[v] = intValue
					} else if mapBody[v] == "boolean" || mapBody[v] == "bool" {
						boolValue, err = strconv.ParseBool(value)
						if err != nil {
							return
						}
						mapBody[v] = boolValue
					} else {
						mapBody[v] = value
					}
				}
			}
		} else {
			err = fmt.Errorf("Get Body: Not Found [%q] value in DepOutVars", v)
			return
		}
	}

	if len(multiVars) == 0 {
		bodys = append(bodys, mapBody)
		LogHandle.Printf("Bodys: %+v", bodys)
		return
	}

	for _, v := range multiVars {
		strList := strings.Split(depOutVars[v], ",")
		for _, sv := range strList {
			var body map[string]interface{}
			body = make(map[string]interface{})
			for mk, mv := range mapBody {
				body[mk] = mv
			}
			if mapBody[v] == "int" || mapBody[v] == "integer" {
				intValue, err := strconv.Atoi(sv)
				if err != nil {
					LogHandle.Printf("Error: %q\n", err)
				}
				body[v] = intValue
			} else if mapBody[v] == "boolean" || mapBody[v] == "bool" {
				boolValue, err = strconv.ParseBool(sv)
				body[v] = boolValue
			} else {
				body[v] = sv
			}
			bodys = append(bodys, body)
		}
	}

	LogHandle.Printf("Bodys: %v", bodys)
	return
}

func (apiCase APICase) RunBeforeCase() (err error) {
	if len(apiCase.BeforeCase) == 0 {
		return
	} else {
		LogHandle.Printf("BeforeCase: %q\n", apiCase.BeforeCase)
	}
	var depCases []string
	var api API
	depCases = strings.Split(apiCase.BeforeCase, ",")
	for _, v := range depCases {
		models.Orm.Table("api_detail").Where("case_id = ?", v).Find(&api)
		LogHandle.Printf("api: %q\n", api)
		if len(api.CaseID) == 0 {
			err = fmt.Errorf("Not Found API Case: [%s] info", v)
			LogHandle.Printf("Error: %s", err)
			return
		}
		err = RunAPI(api.Id, "yes")
		if err != nil {
			return
		}
	}
	return
}

func (apiCase APICase) RunAfterCase() (err error) {
	LogHandle.Printf("AfterCase: %q\n", apiCase.AfterCase)
	if len(apiCase.AfterCase) == 0 {
		return
	}
	var afterCases []string
	var chkApi, delApi API
	afterCases = strings.Split(apiCase.AfterCase, ",")
	var chkId, delId string
	for _, v := range afterCases {
		if strings.HasPrefix(v, "get") {
			chkId = v
		} else if strings.HasPrefix(v, "delete") {
			delId = v
		}
	}
	if len(chkId) > 0 {
		models.Orm.Table("api_detail").Where("case_id = ?", chkId).Find(&chkApi)
		// LogHandle.Printf("chk api: %q\n", chkApi)
		if len(chkApi.CaseID) == 0 {
			err = fmt.Errorf("Not Found API Case: [%s] info", chkId)
			LogHandle.Printf("Error: %s", err)
			return
		}
		err = RunAPI(chkApi.Id, "chk")
		if err != nil {
			return
		}
	}

	if len(delId) > 0 {
		models.Orm.Table("api_detail").Where("case_id = ?", delId).Find(&delApi)
		// LogHandle.Printf("del api: %q\n", delApi)
		if len(delApi.CaseID) == 0 {
			err = fmt.Errorf("Not Found API Case: [%s] info", delId)
			LogHandle.Printf("Error: %s", err)
			return
		}
		err = RunAPI(delApi.Id, "yes")
		if err != nil {
			return
		}
	}

	return
}

func (api API) Run(url string, data map[string]interface{}) (resRaw map[string]interface{}, err error) {

	var req *http.Request
	LogHandle.Printf("Raw Data: %+v", data)
	i := 0
	for k, _ := range data {
		i++
		if ChkUniVar(k, api.Project) {
			var host Host
			models.Orm.Table("host").Where("project = ?", api.Project).Find(&host)
			if host.Testmode == "normal" {
				data[k] = GetRandomStr(12)
			} else if host.Testmode == "abnormal" {
				data[k] = GetRandomStr(256)
			} else {
				data[k] = GetRandomStr(65)
			}

			if i == len(data) {
				LogHandle.Printf("After Data: %+v", data)
			}

		}
	}
	reader, err := json.Marshal(data)
	dest := bytes.NewReader(reader)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	switch api.HttpMethod {
	case "get":
		if data != nil {
			url = url + "?" + data["query"].(string)
		}
		LogHandle.Printf("GET URL: %q", url)
		req, _ = http.NewRequest("GET", url, nil)
	case "post":
		req, _ = http.NewRequest("POST", url, dest)
	case "delete":
		req, _ = http.NewRequest("DELETE", url, dest)
	case "put":
		req, _ = http.NewRequest("PUT", url, dest)
	}
	var hKeys []string
	header, err := api.GetHeader()
	for k := range header {
		hKeys = append(hKeys, k)
	}
	for _, v := range hKeys {
		req.Header.Add(v, header[v])
	}

	resp, err := client.Do(req)
	if err != nil {
		LogHandle.Printf("Request Error: %s", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}

	LogHandle.Printf("response str: %s", string(body))
	resRaw = make(map[string]interface{})
	err = json.Unmarshal(body, &resRaw)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}
	if string(body) == "404 page not found" {
		return
	}
	return
}

func (apiCase APICase) IsThread() (b bool) {
	var host Host
	models.Orm.Table("host").Where("project = ?", apiCase.Project).Find(&host)
	if host.Threading == "yes" {
		b = true
	}
	return
}

func (apiCase APICase) RunThread(api API, urls, querys []string, bodys []map[string]interface{}) (err error) {
	wg := sync.WaitGroup{}
	i := 0
	var tmpMap, response map[string]interface{}
	tmpMap = make(map[string]interface{})
	response = make(map[string]interface{})
	if apiCase.RunNum == 1 {
		for _, url := range urls {
			if api.HttpMethod == "get" && len(querys) == 0 {
				i++
				wg.Add(1)
				go func(i int) {
					LogHandle.Printf("RunTimes: %d", i+1)
					response, err = api.Run(urls[i], nil)
					if err != nil {
						LogHandle.Printf("Error: %s", err)
					}
					apiCase.SaveTestResult(nil, response)
					apiCase.SaveDetailResult(urls[i], api.ApiFunction, nil, response)
					wg.Done()
				}(i)
			} else if api.HttpMethod == "get" && len(querys) > 0 {
				for _, query := range querys {
					wg.Add(1)
					i++
					tmpMap["query"] = query
					go func(i int) {
						LogHandle.Printf("RunTimes: %d", i)
						response, err = api.Run(url, tmpMap)
						if err != nil {
							LogHandle.Printf("Error: %s", err)
						}
						apiCase.SaveTestResult(tmpMap, response)
						apiCase.SaveDetailResult(url, api.ApiFunction, tmpMap, response)
						wg.Done()
					}(i)
				}
			} else {
				for _, body := range bodys {
					wg.Add(1)
					i++
					go func(inVar int, inBody map[string]interface{}) {
						LogHandle.Printf("RunTimes: %d", inVar)
						response, err = api.Run(url, inBody)
						if err != nil {
							LogHandle.Printf("Error: %s", err)
						}
						apiCase.SaveTestResult(inBody, response)
						apiCase.SaveDetailResult(url, api.ApiFunction, inBody, response)
						wg.Done()
					}(i, body)
				}
			}
		}
	} else {
		if api.HttpMethod == "get" && len(querys) == 0 {
			if apiCase.RunNum < len(urls) {
				urls = urls[:apiCase.RunNum]
			} else {
				loopNum := apiCase.RunNum/len(urls) + 1
				for i = 0; i < loopNum; i++ {
					urls = append(urls, urls...)
				}
			}
			for i := 0; i < len(urls); i++ {
				wg.Add(1)
				go func(times int) {
					LogHandle.Printf("RunTimes: %d", times+1)
					response, err = api.Run(urls[times], nil)
					if err != nil {
						LogHandle.Printf("Error: %s", err)
					}
					apiCase.SaveTestResult(nil, response)
					apiCase.SaveDetailResult(urls[times], api.ApiFunction, nil, response)
					wg.Done()
				}(i)
			}
		} else if api.HttpMethod == "get" && len(querys) > 0 {
			if apiCase.RunNum < len(urls)*len(querys) {
				for _, url := range urls {
					for _, query := range querys {
						wg.Add(1)
						i++
						go func(times int) {
							LogHandle.Printf("RunTimes: %d", times)
							tmpMap["query"] = query
							response, err = api.Run(url, tmpMap)
							if err != nil {
								LogHandle.Printf("Error: %s", err)
							}
							apiCase.SaveTestResult(tmpMap, response)
							apiCase.SaveDetailResult(url, api.ApiFunction, tmpMap, response)
							wg.Done()
						}(i)
						if i == apiCase.RunNum {
							break
						}
					}
				}
			} else {
				loopNum := apiCase.RunNum/(len(querys)*len(urls)) + 1
				for j := 0; j < loopNum; j++ {
					for _, url := range urls {
						for _, query := range querys {
							wg.Add(1)
							i++
							go func(times int) {
								LogHandle.Printf("RunTimes: %d", times)
								tmpMap["query"] = query
								response, err = api.Run(url, tmpMap)
								if err != nil {
									LogHandle.Printf("Error: %s", err)
								}
								apiCase.SaveTestResult(tmpMap, response)
								apiCase.SaveDetailResult(url, api.ApiFunction, tmpMap, response)
								wg.Done()
							}(i)
							if i == apiCase.RunNum {
								break
							}
						}
					}
				}
			}
		} else {
			loopNum := (apiCase.RunNum/(len(bodys)*len(urls)) + 1)
			for j := 0; j < loopNum; j++ {
				for _, url := range urls {
					for _, body := range bodys {
						wg.Add(1)
						i++
						go func(times int, inBody map[string]interface{}) {
							LogHandle.Printf("RunTimes: %d", times)
							LogHandle.Printf("body: %+v", inBody)
							response, err = api.Run(url, inBody)
							if err != nil {
								LogHandle.Printf("Error: %s", err)
							}
							apiCase.SaveTestResult(inBody, response)
							apiCase.SaveDetailResult(url, api.ApiFunction, inBody, response)
							wg.Done()
						}(i, body)
					}
				}

			}

		}
	}

	return
}

func RunApiCase(id, depMode string) (err error) {
	var apiCase APICase
	var response map[string]interface{}
	response = make(map[string]interface{})
	s, _ := strconv.Atoi(id)
	models.Orm.Table("api_case").Where("id = ?", s).Find(&apiCase)
	LogHandle.Printf("apiCase: %q\n", apiCase)
	if len(apiCase.CaseID) == 0 {
		err = fmt.Errorf("Not Found APICase info")
		LogHandle.Printf("Error: %s", err)
		return
	}
	err = apiCase.IsRun()
	if err != nil {
		apiCase.SaveTestResult(nil, response)
		LogHandle.Printf("Error: %s", err)
		return
	}
	if depMode != "yes" {
		err = apiCase.RunBeforeCase()
		if err != nil {
			LogHandle.Printf("Error: %s", err)
			return
		}
	}

	api, err := apiCase.GetAPI()
	if err != nil {
		apiCase.SaveTestResult(nil, response)
		LogHandle.Printf("Error: %s", err)
		return
	}
	LogHandle.Printf("api: %q", api)
	rawUrl := apiCase.GetRawUrl()
	depOutVars, err := apiCase.GetDepVars()
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}
	urls, err := GetUrl(rawUrl, depOutVars)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}

	querys, err := api.GetQuery(depOutVars)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}

	bodys, err := api.GetBody(depOutVars)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}
	if apiCase.IsThread() {
		_ = apiCase.RunThread(api, urls, querys, bodys)
		return
	}
	for _, url := range urls {
		if api.HttpMethod == "get" {
			var tmpMap map[string]interface{}
			tmpMap = make(map[string]interface{})
			if len(querys) == 0 {
				response, err = api.Run(url, nil)
				if err != nil {
					LogHandle.Printf("Error: %s", err)
				}
			}
			for _, query := range querys {
				tmpMap["query"] = query
				response, err = api.Run(url, tmpMap)
				if err != nil {
					LogHandle.Printf("Error: %s", err)
				}
				apiCase.SaveTestResult(tmpMap, response)
				apiCase.SaveDetailResult(url, api.ApiFunction, tmpMap, response)
			}

		} else {
			for _, body := range bodys {
				response, err = api.Run(url, body)
				if err != nil {
					LogHandle.Printf("Error: %s", err)
				}
				apiCase.SaveTestResult(body, response)
				apiCase.SaveDetailResult(url, api.ApiFunction, body, response)
			}
		}

	}

	if err != nil {
		if value, ok := response["status"]; ok {
			if value.(string) == "failure" {
				return
			}
		}
	}

	if depMode != "yes" {
		err = apiCase.RunAfterCase()
		if err != nil {
			LogHandle.Printf("Error: %s", err)
		}
	}

	return
}

func RunAPI(id, depMode string) (err error) {
	var api API
	var response map[string]interface{}
	response = make(map[string]interface{})
	s, _ := strconv.Atoi(id)
	models.Orm.Table("api_detail").Where("id = ?", s).Find(&api)
	LogHandle.Printf("api: %q\n", api)
	if len(api.CaseID) == 0 {
		err = fmt.Errorf("Not Found API info")
		LogHandle.Printf("Error: %s", err)
		return
	}
	var apiCase APICase
	models.Orm.Table("api_case").Where("case_id = ?", api.CaseID).Find(&apiCase)
	LogHandle.Printf("apiCase: %q\n", apiCase)
	if len(apiCase.CaseID) == 0 {
		err = fmt.Errorf("Not Found APICase info")
		LogHandle.Printf("Error: %s", err)
		return
	}
	err = apiCase.IsRun()
	if err != nil {
		apiCase.SaveTestResult(nil, response)
		LogHandle.Printf("Error: %s", err)
		return
	}

	if depMode != "yes" && depMode != "chk" {
		err = apiCase.RunBeforeCase()
		if err != nil {
			LogHandle.Printf("Error: %s", err)
			return
		}
	}

	rawUrl := apiCase.GetRawUrl()
	tmpOutVars, err := apiCase.GetDepVars()
	depOutVars, err := api.GetFormatDepVars(tmpOutVars)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}
	urls, err := GetUrl(rawUrl, depOutVars)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}

	querys, err := api.GetQuery(depOutVars)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}

	bodys, err := api.GetBody(depOutVars)
	if err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}
	if apiCase.IsThread() {
		_ = apiCase.RunThread(api, urls, querys, bodys)
		return
	}

	for _, url := range urls {
		if api.HttpMethod == "get" {
			var tmpMap map[string]interface{}
			tmpMap = make(map[string]interface{})
			if len(querys) == 0 {
				response, err = api.Run(url, nil)
				if err != nil {
					LogHandle.Printf("Error: %s", err)
				}
				apiCase.SaveTestResult(nil, response)
				apiCase.SaveDetailResult(url, api.ApiFunction, nil, response)
				if depMode == "chk" {
					var dKeys []string
					for k := range depOutVars {
						dKeys = append(dKeys, k)
					}
					i := 0
					content := response["content"].(map[string]interface{})
					LogHandle.Printf("%-8s\t%-8s\t%-8s\t%-8s\n", "Name", "Expected", "Actual", "Result")
					for _, v := range dKeys {
						if ChkUniVar(v, apiCase.Project) {
							continue
						}
						var chkResult string
						if value, ok := content[v]; ok {
							targetType := fmt.Sprintf("%T", value)
							if targetType == "float64" {
								value = strconv.FormatFloat(value.(float64), 'f', -1, 64)
							}
							if depOutVars[v] != value {
								chkResult = "fail"
								i++
							} else {
								chkResult = "pass"
							}
							LogHandle.Printf("%-8s\t%-8s\t%-8s\t%-8s\n", v, depOutVars[v], value, chkResult)
						}
					}
					if i > 0 {
						err = fmt.Errorf("API[%s] CHECK info fail", api.CaseID)
						return
					}
				}
			} else {
				for _, query := range querys {
					tmpMap["query"] = query
					response, err = api.Run(url, tmpMap)
					if err != nil {
						LogHandle.Printf("Error: %s", err)
					}
					apiCase.SaveTestResult(tmpMap, response)
					apiCase.SaveDetailResult(url, api.ApiFunction, tmpMap, response)
					// if depMode == "chk" {
					// 	i := 0
					// 	LogHandle.Printf("response.Content %s", response.Content)
					// 	for k, v := range response.Content {
					// 		var chkResult string
					// 		if depOutVars[k] != v {
					// 			chkResult = "fail"
					// 			i++
					// 		} else {
					// 			chkResult = "pass"
					// 		}
					// 		LogHandle.Printf("Expected: %s, Actual: %s, Result: %s", depOutVars[k], v, chkResult)
					// 	}
					// 	if i > 0 {
					// 		err = fmt.Errorf("API[%s] check response info fail", api.CaseID)
					// 		return
					// 	}
					// }
				}
			}
		} else {
			for _, body := range bodys {
				response, err = api.Run(url, body)
				if err != nil {
					LogHandle.Printf("Error: %s", err)
				}
				apiCase.SaveTestResult(body, response)
				apiCase.SaveDetailResult(url, api.ApiFunction, body, response)
			}
		}
	}

	if err != nil {
		if value, ok := response["status"]; ok {
			if value.(string) == "failure" {
				return
			}
		}
	}

	if depMode != "yes" && depMode != "chk" {
		err = apiCase.RunAfterCase()
		if err != nil {
			LogHandle.Printf("Error: %s", err)
		}
	}

	return
}

func RunAgain(id string) (err error) {
	var apiTestDetail ApiTestDetail
	s, _ := strconv.Atoi(id)
	models.Orm.Table("api_test_detail").Where("id = ?", s).Find(&apiTestDetail)
	// LogHandle.Printf("apiTestDetail: %q\n", apiTestDetail)
	if len(apiTestDetail.CaseID) == 0 {
		err = fmt.Errorf("Not Found API test detail info")
		LogHandle.Printf("Error: %s", err)
		return
	}
	var api API
	models.Orm.Table("api_detail").Where("case_id = ?", apiTestDetail.CaseID).Find(&api)
	// LogHandle.Printf("api: %q\n", api)
	if len(api.CaseID) == 0 {
		err = fmt.Errorf("Not Found API[%s] info", apiTestDetail.CaseID)
		LogHandle.Printf("Error: %s", err)
		return
	}

	var apiCase APICase
	models.Orm.Table("api_case").Where("case_id = ?", apiTestDetail.CaseID).Find(&apiCase)
	// LogHandle.Printf("apiCase: %q\n", apiCase)
	if len(apiCase.CaseID) == 0 {
		err = fmt.Errorf("Not Found API[%s] info", apiTestDetail.CaseID)
		LogHandle.Printf("Error: %s", err)
		return
	}
	var data, response map[string]interface{}
	response = make(map[string]interface{})
	if err = json.Unmarshal([]byte(apiTestDetail.Body), &data); err != nil {
		LogHandle.Printf("Error: %s", err)
		return
	}
	response, err = api.Run(apiTestDetail.Url, data)
	apiCase.SaveTestResult(data, response)
	apiCase.SaveDetailResult(apiTestDetail.Url, api.ApiFunction, data, response)
	return
}

func RunData(id string) (err error) {
	var apiTestData ApiTestData
	s, _ := strconv.Atoi(id)
	models.Orm.Table("api_test_data").Where("id = ?", s).Find(&apiTestData)
	// LogHandle.Printf("apiTestData: %q\n", apiTestData)
	if len(apiTestData.CaseID) == 0 {
		err = fmt.Errorf("Not Found API test data info")
		LogHandle.Printf("Error: %s", err)
		return
	}

	var api API
	models.Orm.Table("api_detail").Where("case_id = ?", apiTestData.CaseID).Find(&api)
	// LogHandle.Printf("api: %q\n", api)
	if len(api.CaseID) == 0 {
		err = fmt.Errorf("Not Found API[%s] info", apiTestData.CaseID)
		LogHandle.Printf("Error: %s", err)
		return
	}

	var apiCase APICase
	models.Orm.Table("api_case").Where("case_id = ?", apiTestData.CaseID).Find(&apiCase)
	// LogHandle.Printf("apiCase: %q\n", apiCase)
	if len(apiCase.CaseID) == 0 {
		err = fmt.Errorf("Not Found API[%s] info", apiTestData.CaseID)
		LogHandle.Printf("Error: %s", err)
		return
	}
	var data, response map[string]interface{}
	response = make(map[string]interface{})
	if len(apiTestData.Body) > 0 {
		if err = json.Unmarshal([]byte(apiTestData.Body), &data); err != nil {
			LogHandle.Printf("Error: %s", err)
			return
		}
	}

	var url string
	rawUrl := apiCase.GetRawUrl()
	if len(apiTestData.UrlQuery) > 0 {
		if strings.HasPrefix(apiTestData.UrlQuery, "?") {
			url = rawUrl + apiTestData.UrlQuery
		} else {
			url = rawUrl + "?" + apiTestData.UrlQuery
		}
	} else {
		url = rawUrl
	}
	response, err = api.Run(url, data)
	if err != nil {
		return
	}
	if value, ok := response["status"]; ok {
		apiTestData.ActualResult = value.(string)
		if value.(string) == apiTestData.ExpectedResult {
			apiTestData.Result = "pass"
			apiTestData.Response = response["message"].(string)
		}
	} else {
		apiTestData.Result = "fail"
		apiTestData.FailReason = response["message"].(string)
	}

	apiTestData.Module = apiCase.Module
	apiTestData.ApiFunction = api.ApiFunction

	apiCase.SaveTestResult(data, response)
	apiCase.SaveDetailResult(url, api.ApiFunction, data, response)

	err = models.Orm.Table("api_test_data").Where("project = ? and case_id = ? and id = ?", apiTestData.Project, apiTestData.CaseID, s).Update(apiTestData).Error
	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}
	return
}

func RunRandom(id string) (err error) {
	var testcase TestCase
	s, _ := strconv.Atoi(id)
	models.Orm.Table("test_case").Where("id = ?", s).Find(&testcase)
	// LogHandle.Printf("testcase: %q\n", testcase)
	if len(testcase.CaseID) == 0 {
		err = fmt.Errorf("请先关联API，再进行测试")
		LogHandle.Printf("Error: %s", err)
		return
	}

	var api API
	models.Orm.Table("api_detail").Where("case_id = ?", testcase.CaseID).Find(&api)
	// LogHandle.Printf("api: %q\n", api)
	if len(api.CaseID) == 0 {
		err = fmt.Errorf("Not Found API[%s] Test Data", testcase.CaseID)
		LogHandle.Printf("Error: %s", err)
		return
	}

	err = RunAPI(api.Id, "")

	if err != nil {
		testcase.TestResult = "fail"
	} else {
		testcase.TestResult = "pass"
	}

	curTime := time.Now()
	testcase.UpdatedAt = curTime.Format(baseFormat)
	testcase.CaseExecutor = "Robot"
	// LogHandle.Printf("testcase: %s", testcase)
	_ = models.Orm.Table("test_case").Where("id = ?", s).Update(testcase).Error

	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}

	return
}

func RunRegress(id string) (err error) {
	var testcase TestCase
	s, _ := strconv.Atoi(id)
	models.Orm.Table("test_case").Where("id = ?", s).Find(&testcase)
	// LogHandle.Printf("testcase: %q\n", testcase)
	if len(testcase.CaseID) == 0 {
		err = fmt.Errorf("请先关联API，再进行测试")
		LogHandle.Printf("Error: %s", err)
		return
	}

	var apiTestDatas []ApiTestData
	models.Orm.Table("api_test_data").Where("case_id = ?", testcase.CaseID).Find(&apiTestDatas)
	// LogHandle.Printf("apiCase: %q\n", apiCase)
	if len(apiTestDatas) == 0 {
		err = fmt.Errorf("Not Found API[%s] Test Data", testcase.CaseID)
		LogHandle.Printf("Error: %s", err)
		return
	}
	tag := 1
	var errStr string
	for _, item := range apiTestDatas {
		id := strconv.Itoa(item.Id)
		err = RunData(id)
		if err != nil {
			errStr = errStr + fmt.Sprintf("%s", err)
			tag = 0
		}
	}

	if len(errStr) > 0 {
		LogHandle.Printf("Error: %s", errStr)
	}

	if tag == 0 {
		testcase.TestResult = "fail"
		err = fmt.Errorf("%s", errStr)
	} else {
		testcase.TestResult = "pass"
	}

	var afterTests []ApiTestData
	models.Orm.Table("api_test_data").Where("case_id = ?", testcase.CaseID).Find(&afterTests)
	for _, item := range afterTests {
		if item.Result == "fail" {
			testcase.TestResult = "fail"
		}
	}

	curTime := time.Now()
	testcase.UpdatedAt = curTime.Format(baseFormat)
	testcase.CaseExecutor = "Robot"
	// LogHandle.Printf("testcase: %s", testcase)
	_ = models.Orm.Table("test_case").Where("id = ?", s).Update(testcase).Error

	if err != nil {
		LogHandle.Printf("Error: %s", err)
	}

	return
}
