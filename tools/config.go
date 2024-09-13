package tools

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Db struct {
		LocalMyDb     string `yaml:"localMyDb"`
		LocalCkDbUrl  string `yaml:"localCkDbUrl"`
		AmazonCenter  string `yaml:"amazonCenter"`
		ProductSystem string `yaml:"productSystem"`
	}
}

var initPath = "D:\\demogo\\resources\\config.yaml"

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
	return dbMap

}
