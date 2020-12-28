package pages

import (
	"html/template"

	tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/infobox"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	"github.com/GoAdminGroup/themes/sword/components/chart_legend"
	"github.com/gin-gonic/gin"
)

type Box struct {
	Title  template.HTML
	Color  template.HTML
	Number template.HTML
	Icon   template.HTML
}

func GetDashBoardContent(ctx *gin.Context) (types.Panel, error) {

	components := tmpl.Default()
	colComp := components.Col()
	var size = types.SizeMD(3).SM(6).XS(12)
	/**************************
	 * Small Box
	/**************************/

	smallbox1 := smallbox.New().SetColor("yellow").SetIcon("fa-eye").SetUrl("/admin/info/api_case").SetTitle("api").SetValue("接口用例").GetContent()
	smallbox2 := smallbox.New().SetColor("red").SetIcon("fa-eye-slash").SetUrl("/admin/info/test_case").SetTitle("case").SetValue("测试用例").GetContent()
	smallbox3 := smallbox.New().SetColor("green").SetIcon("fa-fast-backward").SetUrl("/admin/info/api_sum_up").SetTitle("api count").SetValue("接口总览").GetContent()
	smallbox4 := smallbox.New().SetColor("purple").SetIcon("fa-fast-forward").SetUrl("/admin/info/testcase_count").SetTitle("case count").SetValue("用例统计").GetContent()

	col1 := colComp.SetSize(size).SetContent(smallbox1).GetContent()
	col2 := colComp.SetSize(size).SetContent(smallbox2).GetContent()
	col3 := colComp.SetSize(size).SetContent(smallbox3).GetContent()
	col4 := colComp.SetSize(size).SetContent(smallbox4).GetContent()

	row1 := components.Row().SetContent(col1 + col2 + col3 + col4).GetContent()
	/**************************
	 * Info Box
	/**************************/
	box1 := Box{"CPU TRAFFIC", "aqua", "100", "ion-ios-gear-outline"}
	box2 := Box{"Likes", "red", "1030.00<small>$</small>", icon.GooglePlus}
	box3 := Box{"Sales", "green", "760", "ion-ios-cart-outline"}
	box4 := Box{"New Members", "yellow", "2,349", "ion-ios-people-outline"}
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
		SetText(box4.Title).
		SetColor(box4.Color).
		SetNumber(box4.Number).
		SetIcon(box4.Icon).
		GetContent()
	infobox6 := infobox.New().
		SetText(box4.Title).
		SetColor(box4.Color).
		SetNumber(box4.Number).
		SetIcon(box4.Icon).
		GetContent()
	infobox7 := infobox.New().
		SetText(box4.Title).
		SetColor(box4.Color).
		SetNumber(box4.Number).
		SetIcon(box4.Icon).
		GetContent()
	infobox8 := infobox.New().
		SetText(box4.Title).
		SetColor(box4.Color).
		SetNumber(box4.Number).
		SetIcon(box4.Icon).
		GetContent()

	infoboxCol1 := colComp.SetSize(size).SetContent(infobox1).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
	infoboxCol5 := colComp.SetSize(size).SetContent(infobox5).GetContent()
	infoboxCol6 := colComp.SetSize(size).SetContent(infobox6).GetContent()
	infoboxCol7 := colComp.SetSize(size).SetContent(infobox7).GetContent()
	infoboxCol8 := colComp.SetSize(size).SetContent(infobox8).GetContent()
	row2 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4 + infoboxCol5 + infoboxCol6 + infoboxCol7 + infoboxCol8).GetContent()

	/**************************
	 * Box
	/**************************/

	table := components.Table().SetInfoList([]map[string]types.InfoItem{
		{
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		},
	}).SetThead(types.Thead{
		{Head: "Order ID"},
		{Head: "Item"},
		{Head: "Status"},
		{Head: "Popularity"},
		{Head: "Popularity"},
		{Head: "Popularity"},
		{Head: "Popularity"},
	}).GetContent()

	row3 := components.Box().
		WithHeadBorder().
		SetHeader("Latest Orders").
		SetHeadColor("#f7f7f7").
		SetBody(table).
		GetContent()

	/**************************
	 * Pie Chart
	/**************************/

	pie := chartjs.Pie().
		SetHeight(170).
		SetLabels([]string{"Navigator", "Opera", "Safari", "FireFox", "IE", "Chrome"}).
		SetID("pieChart").
		AddDataSet("Chrome").
		DSData([]float64{100, 300, 600, 400, 500, 700}).
		DSBackgroundColor([]chartjs.Color{
			"rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)", "rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)",
		}).
		GetContent()

	legend := chart_legend.New().SetData([]map[string]string{
		{
			"label": " Chrome",
			"color": "red",
		}, {
			"label": " IE",
			"color": "Green",
		}, {
			"label": " FireFox",
			"color": "yellow",
		}, {
			"label": " Sarafri",
			"color": "blue",
		}, {
			"label": " Opera",
			"color": "light-blue",
		}, {
			"label": " Navigator",
			"color": "gray",
		},
	}).GetContent()

	boxDanger := components.Box().SetTheme("danger").WithHeadBorder().SetHeader("Project").
		SetBody(components.Row().
			SetContent(colComp.SetSize(types.SizeMD(4)).
				SetContent(pie).
				GetContent() + colComp.SetSize(types.SizeMD(4)).
				SetContent(legend).
				GetContent()).GetContent()).
		GetContent()
	row4 := components.Row().SetContent(boxDanger).GetContent()
	// row6 := components.Row().SetContent(boxcol).GetContent()

	tabs := components.Tabs().SetData([]map[string]template.HTML{
		{
			"title": "环境配置说明",
			"content": template.HTML(`<b>使用示例：:</b>
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

	row5 := components.Row().SetContent(tabs).GetContent()

	return types.Panel{
		Content:     row1 + row3 + row2 + row4 + row5,
		Title:       "Dashboard",
		Description: "",
	}, nil
}
