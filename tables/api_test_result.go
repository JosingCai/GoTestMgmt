package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetApiTestResultTable(ctx *context.Context) table.Table {

	apiTestResult := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := apiTestResult.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("用例ID", "case_id", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("请求参数", "requestVars", db.Longtext)
	info.AddField("测试结果", "result", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("提供依赖变量", "outVars", db.Longtext)
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("创建时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.SetTable("api_test_result").SetTitle("测试结果").SetDescription("测试结果")

	formList := apiTestResult.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("用例ID", "case_id", db.Varchar, form.Text)
	formList.AddField("请求参数", "requestVars", db.Longtext, form.RichText)
	formList.AddField("测试结果", "result", db.Varchar, form.Text)
	formList.AddField("提供依赖变量", "outVars", db.Longtext, form.RichText)
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
