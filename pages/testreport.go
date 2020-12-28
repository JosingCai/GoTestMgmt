package pages

import (
	tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

// type LocalBox struct {
// 	Title  template.HTML
// 	Color  template.HTML
// 	Number template.HTML
// 	Icon   template.HTML
// }

func GetTestReport(ctx *gin.Context) (types.Panel, error) {
	components := tmpl.Default()
	colComp := components.Col()

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
	}).GetContent()

	boxInfo := components.Box().
		WithHeadBorder().
		SetHeader("Latest Orders").
		SetHeadColor("#f7f7f7").
		SetBody(table).
		GetContent()

	tableCol := colComp.SetSize(types.SizeMD(8)).SetContent(boxInfo).GetContent()

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
	row2 := components.Row().SetContent(pie).GetContent()

	return types.Panel{
		Content:     row2 + tableCol,
		Title:       "TestReport",
		Description: "",
	}, nil
}
