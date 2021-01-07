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

func GetTestcaseCountTable(ctx *context.Context) table.Table {

	testcaseCount := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := testcaseCount.GetInfo().HideFilterArea()

	info.SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.SetSortField("created_at")

	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("模块名称", "module", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例总数", "allcase", db.Varchar)
	info.AddField("未合入/未测试", "uninclude_untest", db.Int)
	info.AddField("PASS", "pass", db.Int)
	info.AddField("FAIL", "fail", db.Int)
	info.AddField("PASS率", "pass_per", db.Double)
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("备注", "remark", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldSortable()
	info.AddField("更新时间", "updated_at", db.Timestamp).
		FieldHide()
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	projects := biz.GetProjects()
	info.AddSelectBox("关联项目", projects, action.FieldFilter("project"))

	info.SetTable("testcase_count").SetTitle("用例统计").SetDescription("用例统计")

	formList := testcaseCount.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("模块名称", "module", db.Varchar, form.Text)
	formList.AddField("用例总数", "allcase", db.Varchar, form.Text)
	formList.AddField("未合入/未测试", "uninclude_untest", db.Int, form.Number)
	formList.AddField("PASS", "pass", db.Int, form.Number)
	formList.AddField("FAIL", "fail", db.Int, form.Number)
	formList.AddField("PASS率", "pass_per", db.Double, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("备注", "remark", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide()

	formList.SetTable("testcase_count").SetTitle("用例统计").SetDescription("用例统计")

	return testcaseCount
}
