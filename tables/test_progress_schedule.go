package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetTestProgressScheduleTable(ctx *context.Context) table.Table {

	testProgressSchedule := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := testProgressSchedule.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutThreeCol)
	info.SetFilterFormHeadWidth(4)
	info.SetFilterFormInputWidth(8)
	info.AddField("Id", "id", db.Int).
		FieldHide()
	info.AddField("任务名称", "task_id", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("计划开始时间", "p_start_time", db.Datetime)
	info.AddField("计划结束时间", "p_finish_time", db.Datetime)
	info.AddField("实际开始时间", "a_start_time", db.Datetime)
	info.AddField("实际结束时间", "a_finish_time", db.Datetime)
	info.AddField("完成进度", "progress", db.Varchar)
	info.AddField("版本分支", "milestone", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("优先级", "priority", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("执行人员", "executor", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("关联产品", "project", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("备注", "remark", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("删除时间", "deleted_at", db.Timestamp).
		FieldHide()

	info.SetTable("test_progress_schedule").SetTitle("测试计划").SetDescription("测试计划")

	formList := testProgressSchedule.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("任务名称", "task_id", db.Varchar, form.Text)
	formList.AddField("计划开始时间", "p_start_time", db.Datetime, form.Datetime)
	formList.AddField("计划结束时间", "p_finish_time", db.Datetime, form.Datetime)
	formList.AddField("实际开始时间", "a_start_time", db.Datetime, form.Datetime)
	formList.AddField("实际结束时间", "a_finish_time", db.Datetime, form.Datetime)
	formList.AddField("完成进度", "progress", db.Varchar, form.Text)
	formList.AddField("版本分支", "milestone", db.Varchar, form.Text)
	formList.AddField("优先级", "priority", db.Varchar, form.Text)
	formList.AddField("执行人员", "executor", db.Varchar, form.Text)
	formList.AddField("关联产品", "project", db.Varchar, form.Text)
	formList.AddField("备注", "remark", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert().FieldDisableWhenCreate()
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("删除时间", "deleted_at", db.Timestamp, form.Datetime).
		FieldHide().FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("test_progress_schedule").SetTitle("测试计划").SetDescription("测试计划")

	return testProgressSchedule
}
