package sys

import (
	"easy_file/src/config"
	"easy_file/src/database"
	"easy_file/src/structs"
	"easy_file/src/utils"
	"github.com/gin-gonic/gin"
)

func UserSys(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		kv := config.KV{}
		err := c.ShouldBind(&kv)
		res, err := database.RedisDemo(&kv)
		utils.CommonReturn(c, err, res)
	})

	r.POST("/register", func(c *gin.Context) {
		user := structs.User{}
		err := c.ShouldBind(&user)
		count, res, err := database.DbSql("mixUser", "user").IfExist(user)
		if count == 0 && err == nil {
			utils.CommonReturn(c, "err", res)
		}
		utils.CommonReturn(c, err, res)
	})

	r.POST("/logout", func(c *gin.Context) {
		kv := config.KV{}
		err := c.ShouldBind(&kv)
		res, err := database.RedisDemo(&kv)
		utils.CommonReturn(c, err, res)
	})
}
