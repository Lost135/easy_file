package main

import (
	"easy_file/src/common"
	"easy_file/src/config"
	"easy_file/src/http"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	//初始化yml配置文件
	config.YmlConf()
	//初始化系统日志
	config.SysLogToFile()
	//初始化数据库连接
	config.EtcdDb()
	//创建路由
	gin.DefaultWriter = colorable.NewColorableStdout()
	//TODO 优化if语句
	mode := common.RunMode
	if config.Yml.RunMode != "" {
		mode = config.Yml.RunMode
	}
	gin.SetMode(mode)
	r := gin.New()
	//初始化接口日志
	r.Use(config.ApiLogToFile())
	//调用api
	http.Apis(r)
	//启动
	config.CatchInfo("server started success")
	port := common.Port
	if config.Yml.Port != "" {
		port = config.Yml.Port
	}
	err := r.Run(":" + port)
	config.CatchErr(err)
}
