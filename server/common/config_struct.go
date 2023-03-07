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
	Username  string `json:"username"`
	Password  string `json:"password"`
	Bucket    `json:"bucket"`
	Status    int8   `json:"status"`
	CreatedAt string `json:"createdAt"`
	Deleted   int8   `json:"deleted"`
}

type Bucket struct {
	Name string `json:"name"`
	Role int8   `json:"role"`
}

type File struct {
	Filename  string `json:"filename"`
	Path      string `json:"path"`
	Bucket    string `json:"bucket"`
	CreatedAt string `json:"createdAt"`
	Deleted   int8   `json:"deleted"`
	Public    int8   `json:"public"`
}

type Auth struct {
	Token string `json:"token"`
}

type Claims struct {
	Exp        time.Time `json:"exp"`
	Authorized string    `json:"authorized"`
	UserId     string    `json:"userId"`
	Username   string    `json:"username"`
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Bucket   `json:"bucket"`
	Status   int8   `json:"status"`
	Token    string `json:"token"`
}
