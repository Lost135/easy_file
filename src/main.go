package main

import (
	"easy_file/src/api"
	"easy_file/src/config"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {

	//初始化yml配置文件
	config.YmlConf()
	//初始化系统日志
	config.SysLogToFile()
	//初始化数据库连接
	//config.DBClient()
	config.MongoClient()
	config.RedisClient()
	//创建路由
	gin.DefaultWriter = colorable.NewColorableStdout()
	mode := config.Yml.RunMode
	gin.SetMode(mode)
	r := gin.New()
	//初始化接口日志
	r.Use(config.ApiLogToFile())
	//调用api
	api.SysApi(r)
	api.Apis(r)
	//启动
	config.CatchInfo("server started success")
	err := r.Run(config.Yml.Port)
	config.CatchErr(err)
}
