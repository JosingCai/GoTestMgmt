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

func GetApiTestDataTable(ctx *context.Context) table.Table {

	apiTestData := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := apiTestData.GetInfo().HideFilterArea()
	info.SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("数据描述", "data_desc", db.Varchar).
		FieldFilterable().
		FieldEditAble().FieldWidth(150)
	info.AddField("关联项目", "project", db.Varchar).FieldWidth(120)
	info.AddField("API描述", "apiFunction", db.Varchar).FieldWidth(150)
	info.AddField("关联API", "case_id", db.Varchar).
		FieldFilterable().FieldWidth(150)
	info.AddField("所属模块", "module", db.Varchar).
		FieldFilterable().FieldWidth(150)
	info.AddField("UrlQuery", "urlQuery", db.Longtext).
		FieldFilterable().FieldWidth(300)
	info.AddField("Body", "body", db.Longtext).FieldWidth(300)
	info.AddField("预期结果", "expected_result", db.Varchar).
		FieldFilterable().
		FieldEditAble().FieldWidth(120)
	info.AddField("实际结果", "actual_result", db.Varchar).
		FieldFilterable().FieldWidth(120)
	info.AddField("测试结果", "result", db.Varchar).
		FieldFilterable().FieldWidth(120)
	info.AddField("失败原因", "fail_reason", db.Longtext).
		FieldFilterable().FieldWidth(200)
	info.AddField("返回信息", "response", db.Longtext).FieldWidth(120)
	info.AddField("创建时间", "created_at", db.Timestamp).FieldWidth(120).
		FieldHide()
	info.AddField("更新时间", "updated_at", db.Timestamp).FieldWidth(120)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.AddButton("测试", icon.Android, action.Ajax("testdata_batch_run",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunData(id); err == nil {
					status = "测试完成"
				} else {
					status = fmt.Sprintf("测试失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("测试", action.Ajax("testdata_run",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunData(id); err == nil {
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
		{Value: "pass", Text: "pass"},
		{Value: "fail", Text: "fail"},
	}, action.FieldFilter("result"))

	info.SetTable("api_test_data").SetTitle("测试数据").SetDescription("测试数据")
	// projects := biz.GetProject()
	formList := apiTestData.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("数据描述", "data_desc", db.Varchar, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("API描述", "apiFunction", db.Varchar, form.Text)
	formList.AddField("关联API", "case_id", db.Varchar, form.Text)
	formList.AddField("所属模块", "module", db.Varchar, form.Text)
	formList.AddField("UrlQuery", "urlQuery", db.Longtext, form.Text)
	formList.AddField("Body", "body", db.Longtext, form.Text)
	formList.AddField("预期结果", "expected_result", db.Varchar, form.Text)
	formList.AddField("实际结果", "actual_result", db.Varchar, form.Text)
	formList.AddField("测试结果", "result", db.Varchar, form.Text)
	formList.AddField("失败原因", "fail_reason", db.Longtext, form.Text)
	formList.AddField("返回信息", "response", db.Longtext, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("api_test_data").SetTitle("测试数据").SetDescription("测试数据")

	return apiTestData
}
