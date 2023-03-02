package common

import "github.com/sirupsen/logrus"

const (
	DeletedNot      = 0
	DeletedYes      = 1
	DeletedDefault  = 0
	BucketDefault   = "default"
	FileRoleNone    = 0
	FileRoleOpera   = 1
	FileRoleWrite   = 2
	FileRoleRead    = 4
	PasswordDefault = "password"

	RunMode    = "debug"
	Port       = "8090"
	LogLevel   = logrus.InfoLevel
	ApiLogPath = "./log/"
	ApiLogFile = "apiLog.log"
	SysLogPath = "./log/"
	SysLogFile = "sysLog.log"
	ReadFile   = "../config/conf.yml"

	UserPrefix = "/sys/user/"
	FilePrefix = "/sys/file/"
)
