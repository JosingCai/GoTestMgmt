package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCaseTestCountTable(ctx *context.Context) table.Table {

	caseTestCount := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := caseTestCount.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("用例ID", "case_id", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("功能描述", "API_function", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Run次数", "run_times", db.Int).FieldFilterable()
	info.AddField("测试次数", "test_times", db.Int)
	info.AddField("通过次数", "pass_times", db.Int)
	info.AddField("失败次数", "fail_times", db.Int)
	info.AddField("未测试次数", "untest_times", db.Int)
	info.AddField("测试结果", "test_result", db.Char).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("失败原因", "fail_reason", db.Longtext).FieldWidth(300).
		FieldHide()
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp).
		FieldHide()
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()
	info.AddSelectBox("关联项目", types.FieldOptions{
		{Value: "BOOT3X", Text: "BOOT3X"},
		{Value: "POWER", Text: "POWER"},
		{Value: "REPORT", Text: "REPORT"},
		{Value: "ACT2", Text: "ACT2"},
		{Value: "FLOW", Text: "FLOW"},
		{Value: "PORTAL", Text: "PORTAL"},
		{Value: "CmpSaas", Text: "CmpSaas"},
		{Value: "Catalog", Text: "Catalog"},
		{Value: "Discovery", Text: "Discovery"},
		{Value: "YDDUC", Text: "YDDUC"},
		{Value: "XMDB", Text: "XMDB"},
		{Value: "CmpCore", Text: "CmpCore"},
	}, action.FieldFilter("project"))

	info.SetTable("case_test_count").SetTitle("用例统计").SetDescription("用例统计")

	formList := caseTestCount.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("用例ID", "case_id", db.Varchar, form.Text)
	formList.AddField("功能描述", "API_function", db.Varchar, form.Text)
	formList.AddField("Run次数", "run_times", db.Int, form.Number)
	formList.AddField("测试次数", "test_times", db.Int, form.Number)
	formList.AddField("通过次数", "pass_times", db.Int, form.Number)
	formList.AddField("失败次数", "fail_times", db.Int, form.Number)
	formList.AddField("未测试次数", "untest_times", db.Int, form.Number)
	formList.AddField("测试结果", "test_result", db.Char, form.Text)
	formList.AddField("失败原因", "fail_reason", db.Longtext, form.RichText)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("case_test_count").SetTitle("用例统计").SetDescription("用例统计")

	return caseTestCount
}
