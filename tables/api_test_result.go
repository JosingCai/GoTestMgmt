package tables

import (
	"testmgmt/biz"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetApiTestResultTable(ctx *context.Context) table.Table {

	apiTestResult := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := apiTestResult.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.SetSortField("updated_at")
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("用例ID", "case_id", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("请求参数", "requestVars", db.Longtext).FieldWidth(400)
	info.AddField("测试结果", "result", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("提供依赖变量", "outVars", db.Longtext).FieldWidth(120)
	info.AddField("关联项目", "project", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldHide()
	info.AddField("更新时间", "updated_at", db.Timestamp).FieldWidth(120)
	info.AddField("创建时间", "deleted_at", db.Timestamp).
		FieldHide()
	projects := biz.GetProjects()
	info.AddSelectBox("关联项目", projects, action.FieldFilter("project"))

	info.AddSelectBox("测试结果", types.FieldOptions{
		{Value: "success", Text: "success"},
		{Value: "fail", Text: "fail"},
	}, action.FieldFilter("result"))

	info.SetTable("api_test_result").SetTitle("测试结果").SetDescription("测试结果")

	formList := apiTestResult.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("用例ID", "case_id", db.Varchar, form.Text)
	formList.AddField("请求参数", "requestVars", db.Longtext, form.Text)
	formList.AddField("测试结果", "result", db.Varchar, form.Text)
	formList.AddField("提供依赖变量", "outVars", db.Longtext, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("api_test_result").SetTitle("测试结果").SetDescription("测试结果")

	return apiTestResult
}
