package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/gin-contrib/pprof"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
	ada "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	"github.com/GoAdminGroup/go-admin/plugins"
	_ "github.com/GoAdminGroup/themes/sword" // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"

	"testmgmt/biz"
	"testmgmt/models"
	"testmgmt/pages"
	"testmgmt/tables"

	"github.com/GoAdminGroup/filemanager"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	pprof.Register(r)

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "/uploads")

	// you can custom your pages like:
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/dashboard")
	})

	r.GET("/admin", ada.Content(func(ctx *gin.Context) (panel types.Panel, e error) {
		if config.GetTheme() == "adminlte" {
			return pages.GetDashBoardContent(ctx)
		} else {
			return pages.GetDashBoardContent(ctx)
		}
	}))
	r.GET("/admin/dashboard", ada.Content(pages.GetDashBoardContent))
	r.POST("/admin/smoke", biz.StartSmoke)

	plug, _ := plugins.FindByName("filemanager")
	plug.(*filemanager.FileManager).SetPathValidator(func(path string) error {
		if !strings.Contains(path, "testmgmt") {
			return errors.New("没有权限")
		}
		return nil
	})

	models.Init(eng.MysqlConnection())
	listen := fmt.Sprintf(":%d", biz.SERVER_PORT)
	_ = r.Run(listen)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
