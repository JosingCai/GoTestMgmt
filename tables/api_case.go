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

func GetApiCaseTable(ctx *context.Context) table.Table {

	apiCase := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := apiCase.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("ID", "id", db.Int)
	info.AddField("用例ID", "case_id", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("所属模块", "module", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("执行次数", "runNum", db.Int).FieldFilterable()
	info.AddField("前置用例", "beforeCase", db.Varchar)
	info.AddField("后置用例", "afterCase", db.Varchar)
	info.AddField("提供依赖变量", "outVars", db.Varchar)
	info.AddField("验证返回变量", "chkVars", db.Varchar)
	info.AddField("依赖用例", "param_def", db.Varchar)
	info.AddField("API信息", "raw", db.Longtext)
	info.AddField("关联项目", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}) //.FieldJoin(types.Join{
	// 	Table:     "host",    // 连表的表名
	// 	Field:     "project", // 要连表的字段
	// 	JoinField: "project", // 连表的表的字段

	// })
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	// info.AddButton("测试", icon.Save, action.PopUp("/admin/api/run", "运行"))
	info.AddButton("测试", icon.Android, action.Ajax("apicase_batch_run",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			idStr := ctx.FormValue("ids")
			var status string
			ids := strings.Split(idStr, ",")
			for _, id := range ids {
				if err := biz.RunApiCase(id, ""); err == nil {
					status = "测试完成"
				} else {
					status = fmt.Sprintf("测试失败：%s: %s", id, err)
					return false, status, ""
				}
			}
			return true, status, ""
		}))

	info.AddActionButton("测试", action.Ajax("apicase_run",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			id := ctx.FormValue("id")
			var status string
			if err := biz.RunApiCase(id, ""); err == nil {
				status = "测试完成"
			} else {
				status = fmt.Sprintf("测试失败：%s: %s", id, err)
			}
			return true, status, ""
		}))

	info.SetTable("api_case").SetTitle("API用例").SetDescription("API用例")

	formList := apiCase.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("用例ID", "case_id", db.Varchar, form.Text)
	formList.AddField("所属模块", "module", db.Varchar, form.Text)
	formList.AddField("执行次数", "runNum", db.Int, form.Number).FieldDefault("1")
	formList.AddField("前置用例", "beforeCase", db.Varchar, form.Text)
	formList.AddField("后置用例", "afterCase", db.Varchar, form.Text)
	formList.AddField("提供依赖变量", "outVars", db.Varchar, form.Text)
	formList.AddField("验证返回变量", "chkVars", db.Varchar, form.Text)
	formList.AddField("依赖用例", "param_def", db.Varchar, form.Text)
	formList.AddField("API信息", "raw", db.Longtext, form.Text)
	formList.AddField("关联项目", "project", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("api_case").SetTitle("API用例").SetDescription("API用例")

	return apiCase
}
