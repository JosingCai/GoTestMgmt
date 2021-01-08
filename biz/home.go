package biz

import (
	"html/template"
	"testmgmt/models"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type Project struct {
	Project string `gorm:"column:project" json:"project"`
}

func GetProjects() (projects []types.FieldOption) {
	var dbProjects []Project
	var project types.FieldOption
	models.Orm.Table("host").Find(&dbProjects)

	if len(dbProjects) >= 0 {
		for _, item := range dbProjects {
			project.Value = item.Project
			project.Text = item.Project
			projects = append(projects, project)
		}
	}
	return
}

type ContentHTML struct {
	Content template.HTML
}

type TemplateAPI struct {
	AllCount           template.HTML `gorm:"column:all_count" json:"all_count"`
	AutomatableCount   template.HTML `gorm:"column:automatable_count" json:"automatable_count"`
	UnautomatableCount template.HTML `gorm:"column:unautomatable_count" json:"unautomatable_count"`
	AutoTestCount      template.HTML `gorm:"column:auto_test_count" json:"auto_test_count"`
	UntestCount        template.HTML `gorm:"column:untest_count" json:"untest_count"`
	PassCount          template.HTML `gorm:"column:pass_count" json:"pass_count"`
	FailCount          template.HTML `gorm:"column:fail_count" json:"fail_count"`
	Project            template.HTML `gorm:"column:project" json:"project"`
}

func GetAPISumUp() (sumUps []map[string]types.InfoItem) {
	var sumup map[string]types.InfoItem
	var allApi []TemplateAPI
	models.Orm.Table("api_sum_up").Order("all_count DESC").Limit(10).Find(&allApi)
	for _, item := range allApi {
		sumup = make(map[string]types.InfoItem, 1)
		sumup["关联项目"] = types.InfoItem{Content: item.Project}
		sumup["API总数"] = types.InfoItem{Content: item.AllCount}
		sumup["可自动化数"] = types.InfoItem{Content: item.AutomatableCount}
		sumup["不可自动化数"] = types.InfoItem{Content: item.UnautomatableCount}
		sumup["自动化测试总数"] = types.InfoItem{Content: item.AutoTestCount}
		sumup["未测试总数"] = types.InfoItem{Content: item.UntestCount}
		sumup["通过总数"] = types.InfoItem{Content: item.PassCount}
		sumup["失败总数"] = types.InfoItem{Content: item.FailCount}
		sumUps = append(sumUps, sumup)
	}
	return
}

type Box struct {
	Title  template.HTML
	Color  template.HTML
	Number template.HTML
	Icon   template.HTML
}

func GetBoxData() (boxPlural []Box) {
	var projectCount, apiCaseCount, testCaseCount, dataCaseCount, testAll template.HTML

	models.Orm.Table("host").Count(&projectCount)
	box1 := Box{"环境总数", "white", projectCount, "ion-ios-gear-outline"}
	boxPlural = append(boxPlural, box1)

	models.Orm.Table("api_case").Count(&apiCaseCount)
	box2 := Box{"接口用例总数", "white", apiCaseCount, "ion-ios-gear-outline"}
	boxPlural = append(boxPlural, box2)

	models.Orm.Table("api_test_data").Count(&dataCaseCount)
	box3 := Box{"数据用例总数", "white", dataCaseCount, "ion-ios-gear-outline"}
	boxPlural = append(boxPlural, box3)

	models.Orm.Table("test_case").Count(&testCaseCount)
	box4 := Box{"测试用例总数", "white", testCaseCount, "ion-ios-gear-outline"}
	boxPlural = append(boxPlural, box4)

	models.Orm.Table("api_test_detail").Count(&testAll)
	box5 := Box{"测试总次数", "white", testAll, "ion-ios-gear-outline"}
	boxPlural = append(boxPlural, box5)

	return
}
