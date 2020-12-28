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

func GetApiTestDetailTable(ctx *context.Context) table.Table {

	apiTestDetail := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := apiTestDetail.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.SetSortField("created_at")
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("用例ID", "case_id", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("功能描述", "APIFunction", db.Varchar).FieldWidth(150).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("URL", "url", db.Varchar).FieldWidth(200).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Body", "body", db.Longtext).FieldWidth(200)
	info.AddField("Response", "response", db.Longtext).FieldWidth(200)
	info.AddField("失败原因", "fail_reason", db.Longtext).FieldWidth(120)
	info.AddField("测试结果", "test_result", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联项目", "project", db.Varchar).FieldWidth(120).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp).FieldWidth(120)
	info.AddField("更新时间", "updated_at", db.Timestamp).FieldWidth(120).
		FieldHide()
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.AddButton("再来一次", icon.Android, action.Ajax("apitest_batch_again",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunAgain(id); err == nil {
					status = "测试完成"
				} else {
					status = fmt.Sprintf("测试失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("再来一次", action.Ajax("apitest_again",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunAgain(id); err == nil {
				status = "测试完成"
			} else {
				status = fmt.Sprintf("测试失败：%s: %s", id, err)
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
		{Value: "success", Text: "success"},
		{Value: "fail", Text: "fail"},
	}, action.FieldFilter("test_result"))

	info.SetTable("api_test_detail").SetTitle("结果详情").SetDescription("结果详情")

	formList := apiTestDetail.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("用例ID", "case_id", db.Varchar, form.Text)
	formList.AddField("功能描述", "APIFunction", db.Varchar, form.Text)
	formList.AddField("URL", "url", db.Varchar, form.Text)
	formList.AddField("Body", "body", db.Longtext, form.Text)
	formList.AddField("Response", "response", db.Longtext, form.RichText)
	formList.AddField("失败原因", "fail_reason", db.Longtext, form.RichText)
	formList.AddField("测试结果", "test_result", db.Varchar, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("api_test_detail").SetTitle("结果详情").SetDescription("结果详情")

	return apiTestDetail
}
