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

func GetApiSumUpTable(ctx *context.Context) table.Table {

	apiSumUp := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := apiSumUp.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("API总数", "all_count", db.Int)
	info.AddField("可自动化数", "automatable_count", db.Int)
	info.AddField("不可自动化数", "unautomatable_count", db.Int)
	info.AddField("自动化测试总数", "auto_test_count", db.Int)
	info.AddField("未测试总数", "untest_count", db.Int)
	info.AddField("通过总数", "pass_count", db.Int)
	info.AddField("失败总数", "fail_count", db.Int)
	info.AddField("自动化率", "auto_per", db.Double)
	info.AddField("通过率", "pass_per", db.Double)
	info.AddField("失败率", "fail_per", db.Double)
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp).
		FieldHide()
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()
	projects := biz.GetProjects()
	info.AddSelectBox("关联项目", projects, action.FieldFilter("project"))

	info.SetTable("api_sum_up").SetTitle("API统计").SetDescription("API统计")

	formList := apiSumUp.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("API总数", "all_count", db.Int, form.Number)
	formList.AddField("可自动化数", "automatable_count", db.Int, form.Number)
	formList.AddField("不可自动化数", "unautomatable_count", db.Int, form.Number)
	formList.AddField("自动化测试总数", "auto_test_count", db.Int, form.Number)
	formList.AddField("未测试总数", "untest_count", db.Int, form.Number)
	formList.AddField("通过总数", "pass_count", db.Int, form.Number)
	formList.AddField("失败总数", "fail_count", db.Int, form.Number)
	formList.AddField("自动化率", "auto_per", db.Double, form.Text)
	formList.AddField("通过率", "pass_per", db.Double, form.Text)
	formList.AddField("失败率", "fail_per", db.Double, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("api_sum_up").SetTitle("API统计").SetDescription("API统计")

	return apiSumUp
}
