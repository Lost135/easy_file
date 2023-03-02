package common

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Conf struct {
	RunMode string `yaml:"runMode"`
	Port    string `yaml:"port"`
	Etcd    struct {
		Addr string `yaml:"address"`
		Name string `yaml:"name"`
		Pass string `yaml:"password"`
	}
	RootPassword string `yaml:"rootPassword"`
	Log          struct {
		Level      logrus.Level `yaml:"level"`
		ApiLogPath string       `yaml:"apiLogPath"`
		ApiLogFile string       `yaml:"apiLogFile"`
		SysLogPath string       `yaml:"sysLogPath"`
		SysLogFile string       `yaml:"sysLogFile"`
	}
}

type User struct {
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Bucket    string     `json:"bucket"`
	CreatedAt *time.Time `json:"createdAt"`
	DelFlag   int8       `json:"delFlag"`
}

type File struct {
	Filename  string     `json:"filename"`
	Path      string     `json:"path"`
	Bucket    string     `json:"bucket"`
	CreatedAt *time.Time `json:"createdAt"`
	DelFlag   int8       `json:"delFlag"`
}
