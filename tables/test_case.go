package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetTestCaseTable(ctx *context.Context) table.Table {

	testCase := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := testCase.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("Id", "id", db.Int)
	info.AddField("用例编号", "case_number", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例名称", "case_name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例类型", "case_type", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("优先级", "priority", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("预置条件", "pre_condition", db.Longtext)
	info.AddField("测试范围", "test_range", db.Longtext)
	info.AddField("测试步骤", "test_steps", db.Longtext)
	info.AddField("预期结果", "expect_result", db.Longtext)
	info.AddField("是否自动化", "auto", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("功能开发者", "fun_developer", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例设计者", "case_designer", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例执行者", "case_executor", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("测试时间", "test_time", db.Datetime)
	info.AddField("测试结果", "test_result", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("用例模块", "module", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联AP", "case_id", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("备注", "remark", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

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
	formList.AddField("测试时间", "test_time", db.Datetime, form.Datetime)
	formList.AddField("测试结果", "test_result", db.Varchar, form.Text)
	formList.AddField("用例模块", "module", db.Varchar, form.Text)
	formList.AddField("关联AP", "case_id", db.Varchar, form.Text)
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
