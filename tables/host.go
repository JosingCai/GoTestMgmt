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

func GetHostTable(ctx *context.Context) table.Table {

	host := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := host.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("项目名称", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("IP", "ip", db.Char).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("协议", "protocol", db.Char).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("鉴权", "auth", db.Char).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "yes" {
			return "yes"
		}
		if model.Value == "no" {
			return "no"
		}
		return "unknown"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "yes", Text: "是"},
		{Value: "no", Text: "否"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "yes", Text: "是"},
		{Value: "no", Text: "否"},
	})

	info.AddField("路由前缀", "prepath", db.Varchar).FieldFilterable()
	info.AddField("多线程", "threading", db.Char).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "yes" {
			return "yes"
		}
		if model.Value == "no" {
			return "no"
		}
		return "unknown"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "yes", Text: "是"},
		{Value: "no", Text: "否"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "yes", Text: "是"},
		{Value: "no", Text: "否"},
	})
	info.AddField("用户体系", "usermode", db.Varchar)
	info.AddField("数据库连接", "dbconfig", db.Varchar)
	info.AddField("Token", "token", db.Text).FieldHide()
	info.AddField("测试模式", "testmode", db.Varchar)
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("更新时间", "updated_at", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()
	info.AddButton("Swagger导入API用例", icon.Android, action.Ajax("autogeneration_batch",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.GetSwagger(id); err == nil {
					status = "导入完成"
				} else {
					status = fmt.Sprintf("导入失败：%s", err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("Swagger导入API用例", action.Ajax("autogeneration",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.GetSwagger(id); err == nil {
				status = "导入完成"
			} else {
				status = fmt.Sprintf("导入失败：%s", err)
			}

			return true, status, ""
		}))

	info.AddButton("XMind导入测试用例", icon.Android, action.Ajax("testcase_xmind_batch_import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.GetJSON(id); err == nil {
					status = "导入完成"
				} else {
					status = fmt.Sprintf("导入失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("XMind导入测试用例", action.Ajax("testcase_xmind_import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.GetJSON(id); err == nil {
				status = "导入完成"
			} else {
				status = fmt.Sprintf("导入失败：%s: %s", id, err)
			}
			return true, status, ""
		}))

	info.AddButton("Excel导入测试用例", icon.Android, action.Ajax("testcase_excel_batch_import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunApiCase(id, ""); err == nil {
					status = "导入完成"
				} else {
					status = fmt.Sprintf("导入失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("Excel导入测试用例", action.Ajax("testcase_excel_import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunApiCase(id, ""); err == nil {
				status = "导入完成"
			} else {
				status = fmt.Sprintf("导入失败：%s: %s", id, err)
			}
			return true, status, ""
		}))

	info.AddButton("Excel导入测试计划", icon.Android, action.Ajax("testprogress_batch_import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunApiCase(id, ""); err == nil {
					status = "导入完成"
				} else {
					status = fmt.Sprintf("导入失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("Excel导入测试计划", action.Ajax("testprogress_import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunApiCase(id, ""); err == nil {
				status = "导入完成"
			} else {
				status = fmt.Sprintf("导入失败：%s: %s", id, err)
			}
			return true, status, ""
		}))

	info.AddButton("接口总览", icon.Android, action.Ajax("api_batch_sum_up",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.CountAllAPI(id); err == nil {
					status = "统计完成"
				} else {
					status = fmt.Sprintf("统计失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("接口总览", action.Ajax("api_sum_up",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.CountAllAPI(id); err == nil {
				status = "统计完成"
			} else {
				status = fmt.Sprintf("统计失败：%s: %s", id, err)
			}
			return true, status, ""
		}))
	info.AddButton("接口统计", icon.Android, action.Ajax("apis_batch_sum_up",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.CountAPIs(id); err == nil {
					status = "统计完成"
				} else {
					status = fmt.Sprintf("统计失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("接口统计", action.Ajax("apis_sum_up",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.CountAPIs(id); err == nil {
				status = "统计完成"
			} else {
				status = fmt.Sprintf("统计失败：%s: %s", id, err)
			}
			return true, status, ""
		}))

	info.SetTable("host").SetTitle("测试主机").SetDescription("测试主机")

	formList := host.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("项目名称", "project", db.Varchar, form.Text)
	formList.AddField("IP", "ip", db.Char, form.Ip)
	formList.AddField("协议", "protocol", db.Char, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "http", Value: "http"},
			{Text: "https", Value: "https"},
		}).FieldDefault("http")
	formList.AddField("鉴权", "auth", db.Char, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "是", Value: "yes"},
			{Text: "否", Value: "no"},
		}).FieldDefault("yes")
	formList.AddField("路由前缀", "prepath", db.Varchar, form.Text)
	formList.AddField("多线程", "threading", db.Char, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "是", Value: "yes"},
			{Text: "否", Value: "no"},
		}).FieldDefault("yes")

	formList.AddField("用户体系", "usermode", db.Varchar, form.Radio).FieldOptions(types.FieldOptions{
		{Text: "uam", Value: "uam"},
		{Text: "uc", Value: "uc"},
		{Text: "native", Value: "native"},
	}).FieldDefault("uam")
	formList.AddField("数据库连接", "dbconfig", db.Varchar, form.Text)
	formList.AddField("Token", "token", db.Longtext, form.Text)
	formList.AddField("测试模式", "testmode", db.Varchar, form.Radio).FieldOptions(types.FieldOptions{
		{Text: "normal", Value: "normal"},
		{Text: "abnormal", Value: "abnormal"},
		{Text: "all", Value: "all"},
	}).FieldDefault("abnormal")
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("host").SetTitle("测试主机").SetDescription("测试主机")

	return host
}
