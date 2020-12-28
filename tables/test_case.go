package tables

import (
	"fmt"
	"strings"
	"testmgmt/biz"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetTestCaseTable(ctx *context.Context) table.Table {

	testCase := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := testCase.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("用例编号", "case_number", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例名称", "case_name", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例类型", "case_type", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("优先级", "priority", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("预置条件", "pre_condition", db.Longtext).FieldWidth(120)
	info.AddField("测试范围", "test_range", db.Longtext).FieldWidth(120)
	info.AddField("测试步骤", "test_steps", db.Longtext).FieldWidth(200)
	info.AddField("预期结果", "expect_result", db.Longtext).FieldWidth(200)
	info.AddField("自动化", "auto", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("功能开发者", "fun_developer", db.Varchar).
		FieldHide().FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例设计者", "case_designer", db.Varchar).
		FieldHide().FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例执行者", "case_executor", db.Varchar).
		FieldHide().FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("测试时间", "test_time", db.Varchar).FieldWidth(120).FieldEditAble(editType.Text)
	info.AddField("测试结果", "test_result", db.Varchar).
		FieldWidth(120).
		FieldEditAble(editType.Text).FieldEditOptions(types.FieldOptions{
		{Value: "pass", Text: "pass"},
		{Value: "fail", Text: "fail"},
		{Value: "untest", Text: "untest"},
		{Value: "deprecated", Text: "deprecated"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "pass", Text: "pass"},
		{Value: "fail", Text: "fail"},
		{Value: "untest", Text: "untest"},
		{Value: "deprecated", Text: "deprecated"},
	})

	info.AddField("用例模块", "module", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联AP", "case_id", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联项目", "project", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("备注", "remark", db.Varchar).
		FieldHide().FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp).FieldWidth(120)
	info.AddField("更新时间", "updated_at", db.Timestamp).
		FieldHide()
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.AddButton("随机测试", icon.Android, action.Ajax("random_batch_test",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunRandom(id); err == nil {
					status = "测试完成"
				} else {
					status = fmt.Sprintf("测试失败：%s", err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("随机测试", action.Ajax("random_test",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunRandom(id); err == nil {
				status = "测试完成"
			} else {
				status = fmt.Sprintf("测试失败：%s", err)
			}

			return true, status, ""
		}))

	info.AddButton("回归测试", icon.Android, action.Ajax("regress_batch_test",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunRegress(id); err == nil {
					status = "测试完成"
				} else {
					status = fmt.Sprintf("测试失败：%s", err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("回归测试", action.Ajax("regress_test",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunRegress(id); err == nil {
				status = "测试完成"
			} else {
				status = fmt.Sprintf("测试失败：%s", err)
			}

			return true, status, ""
		}))

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

	info.AddSelectBox("测试结果", types.FieldOptions{
		{Value: "pass", Text: "pass"},
		{Value: "fail", Text: "fail"},
		{Value: "untest", Text: "untest"},
		{Value: "deprecated", Text: "deprecated"},
	}, action.FieldFilter("test_result"))

	info.SetTable("test_case").SetTitle("测试用例").SetDescription("测试用例")

	formList := testCase.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("用例编号", "case_number", db.Varchar, form.Text)
	formList.AddField("用例名称", "case_name", db.Varchar, form.Text)
	formList.AddField("用例类型", "case_type", db.Varchar, form.Text)
	formList.AddField("优先级", "priority", db.Varchar, form.Text)
	formList.AddField("预置条件", "pre_condition", db.Longtext, form.RichText)
	formList.AddField("测试范围", "test_range", db.Longtext, form.RichText)
	formList.AddField("测试步骤", "test_steps", db.Longtext, form.RichText)
	formList.AddField("预期结果", "expect_result", db.Longtext, form.RichText)
	formList.AddField("是否自动化", "auto", db.Varchar, form.Text)
	formList.AddField("功能开发者", "fun_developer", db.Varchar, form.Text)
	formList.AddField("用例设计者", "case_designer", db.Varchar, form.Text)
	formList.AddField("用例执行者", "case_executor", db.Varchar, form.Text)
	formList.AddField("测试时间", "test_time", db.Varchar, form.Date)
	formList.AddField("测试结果", "test_result", db.Varchar, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "pass", Value: "pass"},
			{Text: "fail", Value: "fail"},
			{Text: "untest", Value: "untest"},
			{Text: "deprecated", Value: "deprecated"},
		}).FieldDefault("untest")

	formList.AddField("用例模块", "module", db.Varchar, form.Text)
	formList.AddField("关联API", "case_id", db.Varchar, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("备注", "remark", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("test_case").SetTitle("测试用例").SetDescription("测试用例")

	return testCase
}
