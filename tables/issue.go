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
)

func GetIssueTable(ctx *context.Context) table.Table {

	issue := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := issue.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("Id", "id", db.Int)
	info.AddField("产品", "product", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("版本分支", "milestone", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("No.", "issue_id", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("类型", "issue_type", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("名称", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("提交人", "author", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("解决人", "assignees", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("回归人", "examiner", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("更新时间", "updated", db.Varchar)
	info.AddField("回归结果", "result", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("是否ReOpen", "reopen", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("标签", "tag", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联产品", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("备注", "remark", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.AddButton("同步", icon.Android, action.Ajax("issue_batch_sync",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.GetSwagger(id); err == nil {
					status = "同步完成"
				} else {
					status = fmt.Sprintf("同步失败：%s", err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("同步", action.Ajax("issue_sync",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.GetSwagger(id); err == nil {
				status = "同步完成"
			} else {
				status = fmt.Sprintf("同步失败：%s", err)
			}

			return true, status, ""
		}))
	info.SetTable("issue").SetTitle("Issue列表").SetDescription("Issue列表")

	formList := issue.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("产品", "product", db.Varchar, form.Text)
	formList.AddField("版本分支", "milestone", db.Varchar, form.Text)
	formList.AddField("No.", "issue_id", db.Varchar, form.Text)
	formList.AddField("类型", "issue_type", db.Varchar, form.Text)
	formList.AddField("名称", "name", db.Varchar, form.Text)
	formList.AddField("提交人", "author", db.Varchar, form.Text)
	formList.AddField("解决人", "assignees", db.Varchar, form.Text)
	formList.AddField("回归人", "examiner", db.Varchar, form.Text)
	formList.AddField("更新时间", "updated", db.Varchar, form.Text)
	formList.AddField("回归结果", "result", db.Varchar, form.Text)
	formList.AddField("是否ReOpen", "reopen", db.Varchar, form.Text)
	formList.AddField("标签", "tag", db.Varchar, form.Text)
	formList.AddField("关联产品", "project", db.Varchar, form.Text)
	formList.AddField("备注", "remark", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("issue").SetTitle("Issue列表").SetDescription("Issue列表")

	return issue
}
