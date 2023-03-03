package common

import (
	"github.com/sirupsen/logrus"
)

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
	PublicNot       = 0
	PublicYse       = 1
	PublicDefault   = 0

	RunMode    = "debug"
	Port       = "8090"
	LogLevel   = logrus.InfoLevel
	ApiLogPath = "./log/"
	ApiLogFile = "apiLog.log"
	SysLogPath = "./log/"
	SysLogFile = "sysLog.log"
	ConfigFile = "../config/conf.yml"

	UserPrefix      = "/sys/user/"
	FilePrefix      = "/sys/file/"
	FilePathDefault = "./file/"

	UserStatusOffline = 0
	UserStatusOnline  = 1
	UserStatusDefault = 0

	RootKey       = "/sys/user/root"
	RootUsername  = "root"
	RootPassword  = "password"
	RootRole      = 7
	RootCreatedAt = ""
)
