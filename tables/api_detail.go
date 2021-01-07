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

func GetApiDetailTable(ctx *context.Context) table.Table {

	apiDetail := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := apiDetail.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.AddField("ID", "id", db.Int).
		FieldHide()
	info.AddField("用例ID", "case_id", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("所属模块", "module", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("功能描述", "apiFunction", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("方法", "httpMethod", db.Varchar).FieldWidth(80).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("路径", "path", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Header", "header", db.Longtext).FieldWidth(150)
	info.AddField("Path变量", "pathVariable", db.Longtext).FieldWidth(150)
	info.AddField("Query变量", "queryParameter", db.Longtext).FieldWidth(150)
	info.AddField("Body", "body", db.Longtext).FieldWidth(300)
	info.AddField("Response", "response", db.Longtext).FieldWidth(300)
	info.AddField("关联项目", "project", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldSortable().FieldWidth(120)
	info.AddField("更新时间", "updated_at", db.Timestamp).
		FieldHide()
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.AddButton("测试", icon.Android, action.Ajax("api_batch_run",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunAPI(id, ""); err == nil {
					status = "测试完成"
				} else {
					status = fmt.Sprintf("测试失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("测试", action.Ajax("api_run",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunAPI(id, ""); err == nil {
				status = "测试完成"
			} else {
				status = fmt.Sprintf("测试失败：%s: %s", id, err)
			}
			return true, status, ""
		}))

	projects := biz.GetProjects()
	info.AddSelectBox("关联项目", projects, action.FieldFilter("project"))

	info.AddSelectBox("请求方法", types.FieldOptions{
		{Value: "get", Text: "get"},
		{Value: "post", Text: "post"},
		{Value: "delete", Text: "delete"},
		{Value: "put", Text: "put"},
	}, action.FieldFilter("httpMethod"))

	info.SetTable("api_detail").SetTitle("API详情").SetDescription("API详情")

	formList := apiDetail.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("用例ID", "case_id", db.Varchar, form.Text)
	formList.AddField("所属模块", "module", db.Varchar, form.Text)
	formList.AddField("功能描述", "apiFunction", db.Varchar, form.Text)
	formList.AddField("方法", "httpMethod", db.Varchar, form.Text)
	formList.AddField("路径", "path", db.Varchar, form.Text)
	formList.AddField("Header", "header", db.Text, form.Text)
	formList.AddField("Path变量", "pathVariable", db.Text, form.Text)
	formList.AddField("Query变量", "queryParameter", db.Text, form.Text)
	formList.AddField("Body", "body", db.Text, form.Text)
	formList.AddField("Response", "response", db.Text, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("api_detail").SetTitle("API详情").SetDescription("API详情")

	return apiDetail
}
