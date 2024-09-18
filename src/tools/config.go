package tools

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Db struct {
		LocalMyDb     string `yaml:"localMyDb"`
		LocalCkDbUrl  string `yaml:"localCkDbUrl"`
		AmazonCenter  string `yaml:"amazonCenter"`
		ProductSystem string `yaml:"productSystem"`
		WalmartCenter string `yaml:"walmartCenter"`
	}
}

var initPath = "resources/config.yaml"

// var initPath = "/Volumes/SolfD/GolangDemo/demogo/resources/config.yaml"
func ConfigDbConnUrl(path string) map[string]string {
	if path == "" {
		path = initPath
	}
	// 读取yml文件 yaml path = config.yaml
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	// 新建map
	dbMap := make(map[string]string)
	dbMap["localMyDb"] = config.Db.LocalMyDb
	dbMap["localCkDbUrl"] = config.Db.LocalCkDbUrl
	dbMap["amazonCenter"] = config.Db.AmazonCenter
	dbMap["productSystem"] = config.Db.ProductSystem
	dbMap["walmartCenter"] = config.Db.WalmartCenter
	return dbMap

}
