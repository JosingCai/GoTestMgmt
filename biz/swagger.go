package biz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"testmgmt/models"
)

type Swagger struct {
	Paths       map[string]PathDef   `json:"paths"`
	Definitions map[string]DefiniDef `json:"definitions"`
}

type PathDef struct {
	Put    ApiDetail `json:"put"`
	Post   ApiDetail `json:"post"`
	Get    ApiDetail `json:"get"`
	Delete ApiDetail `json:"delete"`
}

type ApiDetail struct {
	Description string         `json:"description"`
	Consumes    []string       `json:"consumes"`
	Produces    []string       `json:"produces"`
	Tags        []string       `json:"tags"`
	Summary     string         `json:"summary"`
	Parameters  []ParamDetail  `json:"parameters"`
	Responses   ResponseDetail `json:"responses"`
}

type ParamDetail struct {
	Type        string       `json:"type"`
	Description string       `json:"description"`
	Name        string       `json:"name"`
	In          string       `json:"in"`
	Required    bool         `json:"required"`
	Schema      SchemaDetail `json:"schema"`
}

type DefiniDef struct {
	Type       string                  `json:"type"`
	Required   []string                `json:"required"`
	Properties map[string]ProperDetail `json:"properties"`
}

type ProperDetail struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Ref         string `json:"$ref"`
}

type SchemaDetail struct {
	Type string `json:"type"`
	Ref  string `json:"$ref"`
}

type ResponseDetail struct {
	R200 R200Detail `json:"200"`
}

type R200Detail struct {
	Description string       `json:"description"`
	Schema      SchemaDetail `json:"schema"`
}

type DbApiTable struct {
	Id string `gorm:"column:id" json:"id"`
	ApiTable
}

type DbApiCase struct {
	Id string `gorm:"column:id" json:"id"`
	APICase
}

type ApiTable struct {
	CaseID         string `gorm:"column:case_id" json:"case_id"`
	Module         string `gorm:"column:module" json:"module"`
	ApiFunction    string `gorm:"column:apiFunction" json:"apiFunction"`
	Body           string `gorm:"column:body" json:"body"`
	Response       string `gorm:"column:response" json:"response"`
	Path           string `gorm:"column:path" json:"path"`
	HttpMethod     string `gorm:"column:httpMethod" json:"httpMethod"`
	Header         string `gorm:"column:header" json:"header"`
	PathVariable   string `gorm:"column:pathVariable" json:"pathVariable"`
	QueryParameter string `gorm:"column:queryParameter" json:"queryParameter"`
	Project        string `gorm:"column:project" json:"project"`
}

func GetLastOne(item string) (value string) {
	tmps := strings.Split(item, "/")
	value = tmps[len(tmps)-1]
	return
}

type DoubleMap map[string]map[string]string
type SingleMap map[string]string

func (swagger Swagger) GetAllDefinition() (defAllDict DoubleMap) {
	defAllDict = make(DoubleMap)
	definiKeys := make([]string, len(swagger.Definitions))
	j := 0
	for k := range swagger.Definitions {
		definiKeys[j] = k
		j++
	}

	for _, v := range definiKeys {
		var defDict SingleMap
		defDict = make(SingleMap)
		subProperty := swagger.Definitions[v].Properties
		subProKeys := make([]string, len(subProperty))
		j = 0
		for k := range subProperty {
			subProKeys[j] = k
			j++
		}
		for _, sv := range subProKeys {
			defDict[sv] = subProperty[sv].Type

		}
		defAllDict[v] = defDict
	}

	return

}

func (apiDetail ApiDetail) GetRequestData(allDefini DoubleMap) (mapHeader, mapBody, mapPath, mapQuery, mapResponse SingleMap) {
	mapHeader = make(SingleMap)
	mapBody = make(SingleMap)
	mapPath = make(SingleMap)
	mapQuery = make(SingleMap)
	mapResponse = make(SingleMap)
	parameters := apiDetail.Parameters

	if len(apiDetail.Responses.R200.Schema.Ref) > 0 {
		resDef := GetLastOne(apiDetail.Responses.R200.Schema.Ref)
		mapResponse, _ = allDefini[resDef]
	}

	for _, pv := range parameters {
		var targetValue string
		var tmpMap SingleMap
		tmpMap = make(SingleMap)
		if len(pv.Schema.Ref) > 0 {
			depDef := GetLastOne(pv.Schema.Ref)
			tmpMap, _ = allDefini[depDef]
		} else {
			targetValue = pv.Type
		}

		if len(targetValue) > 0 {
			switch pv.In {
			case "header":
				mapHeader[pv.Name] = targetValue
			case "query":
				mapQuery[pv.Name] = targetValue
			case "path":
				mapPath[pv.Name] = targetValue
			case "body":
				mapBody[pv.Name] = targetValue
			}
		} else {
			switch pv.In {
			case "header":
				mapHeader = tmpMap
			case "query":
				mapQuery = tmpMap
			case "path":
				mapPath = tmpMap
			case "body":
				mapBody = tmpMap
			}
		}

	}
	return
}

func (pathDef PathDef) GetApiDetail(method, path, project string, allDefini DoubleMap) (err error) {
	var mapHeader, mapBody, mapPath, mapQuery, mapResponse SingleMap
	var apiTable ApiTable
	var dbData DbApiTable
	var caseId string
	var apiDetail ApiDetail
	switch method {
	case "put":
		apiDetail = pathDef.Put
		caseId = "put_" + path
		apiTable.HttpMethod = "put"
	case "post":
		apiDetail = pathDef.Post
		caseId = "post_" + path
		apiTable.HttpMethod = "post"
	case "delete":
		apiDetail = pathDef.Delete
		caseId = "delete_" + path
		apiTable.HttpMethod = "delete"
	case "get":
		apiDetail = pathDef.Get
		caseId = "get_" + path
		apiTable.HttpMethod = "get"
	}
	apiTable.CaseID = caseId
	if len(apiDetail.Tags) > 0 {
		apiTable.Module = apiDetail.Tags[0]
	} else {
		apiTable.Module = "other"
	}
	mapHeader, mapBody, mapPath, mapQuery, mapResponse = apiDetail.GetRequestData(allDefini)

	apiTable.ApiFunction = apiDetail.Description
	apiTable.Path = path
	apiTable.Project = project

	mh, _ := json.Marshal(mapHeader)
	apiTable.Header = string(mh)

	mp, _ := json.Marshal(mapPath)
	apiTable.PathVariable = string(mp)

	mq, _ := json.Marshal(mapQuery)
	apiTable.QueryParameter = string(mq)

	mb, _ := json.Marshal(mapBody)
	apiTable.Body = string(mb)

	ms, _ := json.Marshal(mapResponse)
	apiTable.Response = string(ms)

	models.Orm.Table("api_detail").Where("project = ? and case_id = ?", project, caseId).Find(&dbData)
	if len(dbData.CaseID) == 0 {
		err = models.Orm.Table("api_detail").Create(apiTable).Error
	} else {
		err = models.Orm.Table("api_detail").Where("id = ?", dbData.Id).Update(apiTable).Error
	}

	var apiCase APICase
	apiCase.CaseID = caseId
	apiCase.RunNum = 1
	apiCase.Project = project
	apiCase.Raw = apiTable.ApiFunction + "|" + "http" + "|" + apiTable.HttpMethod + "|" + path + "|" + string(mh) + "|" + string(mp) + "|" + string(mq) + "|" + string(mb) + "|" + string(ms)
	apiCase.Module = apiTable.Module
	curTime := time.Now()
	apiCase.CreatedAt = curTime.Format(baseFormat)
	var dbApi DbApiCase
	models.Orm.Table("api_case").Where("project = ? and case_id = ?", project, caseId).Find(&dbApi)
	if len(dbApi.CaseID) == 0 {
		err = models.Orm.Table("api_case").Create(apiCase).Error
	} else {
		err = models.Orm.Table("api_case").Where("id = ?", dbApi.Id).Update(apiCase).Error
	}

	return
}

func GetSwagger(id string) (err error) {
	var host Host
	models.Orm.Table("host").Where("id = ?", id).Find(&host)
	if len(host.Project) == 0 {
		err = fmt.Errorf("Not found related project id:%s", id)
		return
	}
	project := host.Project
	fileName := fmt.Sprintf("%s/testmgmt/api/%s.json", BASEPATH, project)
	content, err := ioutil.ReadFile(fileName)
	// LogHandle.Printf("content: %s", content)
	if err != nil {
		return
	}
	var swagger Swagger
	err = json.Unmarshal([]byte(content), &swagger)
	allDefini := swagger.GetAllDefinition()
	pathKeys := make([]string, len(swagger.Paths))
	j := 0
	for k := range swagger.Paths {
		pathKeys[j] = k
		j++
	}
	var pathDef PathDef
	for _, v := range pathKeys {
		pathDef = swagger.Paths[v]
		if !reflect.DeepEqual(pathDef.Put, ApiDetail{}) {
			err = pathDef.GetApiDetail("put", v, project, allDefini)
		}
		if !reflect.DeepEqual(pathDef.Post, ApiDetail{}) {
			err = pathDef.GetApiDetail("post", v, project, allDefini)
		}
		if !reflect.DeepEqual(pathDef.Delete, ApiDetail{}) {
			err = pathDef.GetApiDetail("delete", v, project, allDefini)
		}
		if !reflect.DeepEqual(pathDef.Get, ApiDetail{}) {
			err = pathDef.GetApiDetail("get", v, project, allDefini)
		}

	}

	return

}
