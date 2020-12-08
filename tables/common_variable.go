package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCommonVariableTable(ctx *context.Context) table.Table {

	commonVariable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := commonVariable.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("Id", "id", db.Int)
	info.AddField("名称", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("值", "value", db.Longtext).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.SetTable("common_variable").SetTitle("公共变量").SetDescription("公共变量")

	formList := commonVariable.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("名称", "name", db.Varchar, form.Text)
	formList.AddField("值", "value", db.Longtext, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("common_variable").SetTitle("公共变量").SetDescription("公共变量")

	return commonVariable
}
