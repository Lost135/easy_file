package config

import (
	"easy_file/src/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

var Syslog *logrus.Logger

func MkDir(path string) {
	//创建目录
	flag := true
	fi, err := os.Stat(path)
	if err != nil {
		flag = os.IsExist(err) //err!=nil,使用os.IsExist()判断为false,说明文件或文件夹不存在
	} else {
		flag = fi.IsDir() //err==nil,说明文件或文件夹存在
	}

	if flag {
		fmt.Println("Log path exist")
	} else {
		fmt.Println("Log path not exist, creating...")
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = os.Chmod(path, 0777) //通过chmod重新赋权限
	if err != nil {
		fmt.Println("err", err)
	}

}

func CWFile(file string) *os.File {
	//写入文件
	src, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("err", err)
	}
	return src
}

func ApiLogToFile() gin.HandlerFunc {

	//日志文件
	apiLogPath := common.ApiLogPath
	if Yml.Log.ApiLogPath != "" {
		apiLogPath = Yml.Log.ApiLogPath
	}

	apiLogFile := common.ApiLogFile
	if Yml.Log.ApiLogFile != "" {
		apiLogFile = Yml.Log.ApiLogFile
	}
	file := path.Join(apiLogPath, apiLogFile)
	fi := CWFile(file)

	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = fi
	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func SysLogToFile() {
	//TODO 程序运行时日志文件被占用，无法进行其他操作
	sysLogPath := common.SysLogPath
	if Yml.Log.ApiLogPath != "" {
		sysLogPath = Yml.Log.SysLogPath
	}

	sysLogFile := common.SysLogFile
	if Yml.Log.SysLogFile != "" {
		sysLogFile = Yml.Log.SysLogFile
	}
	MkDir(sysLogPath)
	//日志文件
	file := path.Join(sysLogPath, sysLogFile)
	//写入文件
	fi := CWFile(file)

	//实例化
	Syslog = logrus.New()
	//设置输出
	Syslog.Out = fi
	//设置日志级别
	if strconv.Itoa(int(Yml.Log.Level)) != "" {
		Syslog.SetLevel(Yml.Log.Level)
	} else {
		Syslog.SetLevel(common.LogLevel)
	}
	//设置日志格式
	Syslog.SetFormatter(&logrus.JSONFormatter{})
}

func SysLog() *logrus.Entry {
	return Syslog.WithFields(logrus.Fields{})
}

func CatchErr(err interface{}) {
	if err != nil {
		SysLog().Errorf("%s", err)
		return
	}
}

func CatchInfo(err interface{}) {
	if err != nil {
		SysLog().Infof("%s", err)
		return
	}
}

func CatchWarn(err interface{}) {
	if err != nil {
		SysLog().Warnf("%s", err)
		return
	}
}

func CatchFatal(err interface{}) {
	if err != nil {
		SysLog().Fatalf("%s", err)
		return
	}
}
