package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProductGitlabTable(ctx *context.Context) table.Table {

	productGitlab := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := productGitlab.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("URL", "repo", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("产品", "product", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("RSS_TOKEN", "rss_token", db.Varchar)
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("版本分支", "milestone", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("备注", "remark", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.AddButton("同步", icon.Android, action.Ajax("issue_batch_sync",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			var status string
			status = "功能暂不支持"
			// idStr := ctx.FormValue("ids")
			// ids := strings.Split(idStr, ",")
			// for _, id := range ids {
			// 	if err := biz.GetSwagger(id); err == nil {
			// 		status = "导入完成"
			// 	} else {
			// 		status = fmt.Sprintf("导入失败：%s", err)
			// 		return false, status, ""
			// 	}
			// }
			return true, status, ""
		}))

	info.AddActionButton("同步", action.Ajax("issue_sync",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			var status string
			status = "功能暂不支持"
			// id := ctx.FormValue("id")
			// if err := biz.GetSwagger(id); err == nil {
			// 	status = "导入完成"
			// } else {
			// 	status = fmt.Sprintf("导入失败：%s", err)
			// }

			return true, status, ""
		}))

	info.AddButton("统计", icon.Android, action.Ajax("issue_batch_count",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			var status string
			status = "功能暂不支持"
			// idStr := ctx.FormValue("ids")
			// ids := strings.Split(idStr, ",")
			// for _, id := range ids {
			// 	if err := biz.GetSwagger(id); err == nil {
			// 		status = "导入完成"
			// 	} else {
			// 		status = fmt.Sprintf("导入失败：%s", err)
			// 		return false, status, ""
			// 	}
			// }
			return true, status, ""
		}))

	info.AddActionButton("统计", action.Ajax("issue_count",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			var status string
			status = "功能暂不支持"
			// id := ctx.FormValue("id")
			// if err := biz.GetSwagger(id); err == nil {
			// 	status = "导入完成"
			// } else {
			// 	status = fmt.Sprintf("导入失败：%s", err)
			// }

			return true, status, ""
		}))

	info.SetTable("product_gitlab").SetTitle("Gitlab列表").SetDescription("Gitlab列表")

	formList := productGitlab.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("URL", "repo", db.Varchar, form.Text)
	formList.AddField("产品", "product", db.Varchar, form.Text)
	formList.AddField("RSS_TOKEN", "rss_token", db.Varchar, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("版本分支", "milestone", db.Varchar, form.Text)
	formList.AddField("备注", "remark", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("product_gitlab").SetTitle("Gitlab列表").SetDescription("Gitlab列表")

	return productGitlab
}
