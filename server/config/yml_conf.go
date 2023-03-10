package config

import (
	"easy_file/server/common"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Yml *common.Conf

func YmlConf() {
	yamlFile, err := os.ReadFile(common.ConfigFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Yml)
	// TODO 如果值为空则设置默认值  如果有环境变量则优先读取
	//配置文件读取顺序 环境变量 > conf.yml > 默认值
	if err != nil {
		fmt.Println(err.Error())
	}
}
