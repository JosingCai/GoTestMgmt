package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetIssueTagCountTable(ctx *context.Context) table.Table {

	issueTagCount := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := issueTagCount.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("Id", "id", db.Int)
	info.AddField("版本分支", "milestone", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("标签", "tag", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("总数", "all_count", db.Int)
	info.AddField("Closed状态数", "open_count", db.Int)
	info.AddField("Closed_count", "closed_count", db.Int)
	info.AddField("产品", "product", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联产品", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("备注", "remark", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.SetTable("issue_tag_count").SetTitle("标签统计").SetDescription("标签统计")

	formList := issueTagCount.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("版本分支", "milestone", db.Varchar, form.Text)
	formList.AddField("标签", "tag", db.Varchar, form.Text)
	formList.AddField("总数", "all_count", db.Int, form.Number)
	formList.AddField("Open状态数", "open_count", db.Int, form.Number)
	formList.AddField("Closed状态数", "closed_count", db.Int, form.Number)
	formList.AddField("产品", "product", db.Varchar, form.Text)
	formList.AddField("关联产品", "project", db.Varchar, form.Text)
	formList.AddField("备注", "remark", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("issue_tag_count").SetTitle("标签统计").SetDescription("标签统计")

	return issueTagCount
}
