package biz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testmgmt/models"

	"github.com/GoAdminGroup/go-admin/template/types"
)

var (
	BASEPATH    string
	SERVER_PORT int
)

type Project struct {
	Project string `gorm:"column:project" json:"project"`
}

type Config struct {
	FileBasePath string `json:"file_base_path"`
	ServerPort   int    `json:"server_port"`
}

func GetProjects() (projects []types.FieldOption) {
	var dbProjects []Project
	var project types.FieldOption
	models.Orm.Table("host").Find(&dbProjects)

	if len(dbProjects) >= 0 {
		for _, item := range dbProjects {
			project.Value = item.Project
			project.Text = item.Project
			projects = append(projects, project)
		}
	}
	return
}

func init() {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		info := fmt.Sprintf("Init Config Failed: %s", err)
		panic(info)
	}
	var config Config
	err = json.Unmarshal([]byte(content), &config)
	if err != nil {
		info := fmt.Sprintf("Init Config Failed: %s", err)
		panic(info)
	}
	if len(config.FileBasePath) == 0 {
		panic("Not Found file_base_path")
	}

	BASEPATH = config.FileBasePath
	SERVER_PORT = config.ServerPort
	_, err = os.Stat(BASEPATH)
	if err != nil {
		if os.IsNotExist(err) {
			subErr := os.MkdirAll(BASEPATH, os.ModePerm)
			if subErr != nil {
				LogHandle.Printf("Error: %s", subErr)
			}
		}
	}
	subPaths := [4]string{
		BASEPATH + "/" + "api",
		BASEPATH + "/" + "file",
		BASEPATH + "/" + "test",
		BASEPATH + "/" + "log",
	}

	for _, item := range subPaths {
		_, subErr := os.Stat(item)
		if subErr != nil {
			if os.IsNotExist(err) {
				subSubErr := os.MkdirAll(item, os.ModePerm)
				if subSubErr != nil {
					LogHandle.Printf("Error: %s", subSubErr)
				}
			}
		}
	}

}
