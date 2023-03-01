package structs

import "github.com/sirupsen/logrus"

type Conf struct {
	RunMode string `yaml:"runMode"`
	Port    string `yaml:"port"`
	Etcd    struct {
		Addr string `yaml:"address"`
		Name string `yaml:"name"`
		Pass string `yaml:"password"`
	}
	Mysql string `yaml:"mysql"`
	Redis struct {
		Addr   string `yaml:"address"`
		Passwd string `yaml:"password"`
		Db     int    `yaml:"db"`
	}
	Mongodb struct {
		Addr     string `yaml:"addr"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Log struct {
		Level      logrus.Level `yaml:"level"`
		ApiLogPath string       `yaml:"apiLogPath"`
		ApiLogFile string       `yaml:"apiLogFile"`
		SysLogPath string       `yaml:"sysLogPath"`
		SysLogFile string       `yaml:"sysLogFile"`
	}
}
