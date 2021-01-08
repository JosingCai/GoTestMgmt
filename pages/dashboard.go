package pages

import (
	"html/template"
	"testmgmt/biz"

	tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/infobox"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	"github.com/gin-gonic/gin"
)

func GetDashBoardContent(ctx *gin.Context) (types.Panel, error) {

	components := tmpl.Default()
	colComp := components.Col()
	var size = types.SizeMD(2).SM(6).XS(12)
	/**************************
	 * Small Box
	/**************************/

	smallbox1 := smallbox.New().SetColor("white").SetIcon("fa-fast-forward").SetUrl("/admin/info/host").SetTitle("host").SetValue("环境").GetContent()
	smallbox2 := smallbox.New().SetColor("white").SetIcon("fa-fast-forward").SetUrl("/admin/info/api_case").SetTitle("api").SetValue("接口").GetContent()
	smallbox3 := smallbox.New().SetColor("white").SetIcon("fa-fast-forward").SetUrl("/admin/info/api_test_data").SetTitle("data").SetValue("数据").GetContent()
	smallbox4 := smallbox.New().SetColor("white").SetIcon("fa-fast-forward").SetUrl("/admin/info/test_case").SetTitle("case").SetValue("用例").GetContent()
	smallbox5 := smallbox.New().SetColor("white").SetIcon("fa-fast-forward").SetUrl("/admin/info/api_test_detail").SetTitle("case count").SetValue("结果").GetContent()

	col1 := colComp.SetSize(size).SetContent(smallbox1).GetContent()
	col2 := colComp.SetSize(size).SetContent(smallbox2).GetContent()
	col3 := colComp.SetSize(size).SetContent(smallbox3).GetContent()
	col4 := colComp.SetSize(size).SetContent(smallbox4).GetContent()
	col5 := colComp.SetSize(size).SetContent(smallbox5).GetContent()

	row1 := components.Row().SetContent(col1 + col2 + col3 + col4 + col5).GetContent()
	/**************************
	 * Info Box
	/**************************/
	boxPlural := biz.GetBoxData()

	box1 := boxPlural[0]
	box2 := boxPlural[1]
	box3 := boxPlural[2]
	box4 := boxPlural[3]
	box5 := boxPlural[4]
	infobox1 := infobox.New().
		SetText(box1.Title).
		SetColor(box1.Color).
		SetNumber(box1.Number).
		SetIcon(box1.Icon).
		GetContent()

	infobox2 := infobox.New().
		SetText(box2.Title).
		SetColor(box2.Color).
		SetNumber(box2.Number).
		SetIcon(box2.Icon).
		GetContent()

	infobox3 := infobox.New().
		SetText(box3.Title).
		SetColor(box3.Color).
		SetNumber(box3.Number).
		SetIcon(box3.Icon).
		GetContent()

	infobox4 := infobox.New().
		SetText(box4.Title).
		SetColor(box4.Color).
		SetNumber(box4.Number).
		SetIcon(box4.Icon).
		GetContent()

	infobox5 := infobox.New().
		SetText(box5.Title).
		SetColor(box5.Color).
		SetNumber(box5.Number).
		SetIcon(box5.Icon).
		GetContent()

	infoboxCol1 := colComp.SetSize(size).SetContent(infobox1).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
	infoboxCol5 := colComp.SetSize(size).SetContent(infobox5).GetContent()

	row2 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4 + infoboxCol5).GetContent()

	/**************************
	 * Box
	/**************************/
	apiSumUp := biz.GetAPISumUp()
	table := components.Table().SetInfoList(apiSumUp).SetThead(types.Thead{
		{Head: "关联项目"},
		{Head: "API总数"},
		{Head: "可自动化数"},
		{Head: "不可自动化数"},
		{Head: "自动化测试总数"},
		{Head: "未测试总数"},
		{Head: "通过总数"},
		{Head: "失败总数"},
	}).GetContent()

	row3 := components.Box().
		WithHeadBorder().
		SetHeader("接口总览 - top10").
		SetHeadColor("#f7f7f7").
		SetBody(table).
		GetContent()

	tabs := components.Tabs().SetData([]map[string]template.HTML{
		{
			"title": "环境配置说明",
			"content": template.HTML(`<b>示例:</b>
                <p>1. 新建测试环境TEST</p>
                <p>2. 在文件管理Swagger子菜单：导入API文件(TEST.json), 导入测试用例文件（TEST.xmind）</p>
                <p>3. 选中环境，点击Swagger导入API用例，在用例菜单下即可看到对应API数据</p>
                <p>4. 选中环境，点击Xmind导入测试用例，在用例菜单下即可看到对应测试用例数据</p>
                <p>5. 选中环境, 点击接口统计，在报表菜单下即可看到接口统计数据</p>
                <p>6. 选中环境, 点击用例统计，在报表菜单下即可看到用例统计数据</p>
                <p>7. 选中环境, 点击接口总览，在报表菜单下即可看到总个项目的接口统考数据</p>
                `),
		}, {
			"title": "接口用例说明",
			"content": template.HTML(`
                <p>1. 在环境设置abnomal模式，其他相关信息，API进行随机测试</p>
                <p>2. 在测试数据菜单下准备完整的body或urlquery，方便API进行回归测试</p>
                <p>3. 在公用变量定义API需要的各项数据，API进行组合数据测试</p>
              `),
		}, {
			"title": "测试用例说明",
			"content": template.HTML(`
                <p>1. 给测试用例绑定接口用例CaseID</p>
                <p>2. 选择用例，点击回归测试，会执行测试数据下相同CaseID的全部数据</p>
                <p>3. 选择用例，点击随机测试，会根据环境的测试模式，进行随机测试</p>
              `),
		},
	}).GetContent()

	useDanger := components.Box().SetTheme("danger").WithHeadBorder().SetHeader("使用说明").
		SetBody(components.Row().
			SetContent(colComp.SetSize(types.SizeMD(16)).
				SetContent(tabs).
				GetContent()).GetContent()).
		GetContent()
	row5 := components.Row().SetContent(useDanger).GetContent()

	return types.Panel{
		Content:     row1 + row2 + row5 + row3,
		Title:       "Dashboard",
		Description: "",
	}, nil
}
