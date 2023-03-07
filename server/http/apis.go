package http

import (
	"easy_file/server/common"
	"easy_file/server/config"
	"easy_file/server/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Apis(route *gin.Engine) {
	group := route.Group(common.ApiPrefix).Group(common.Version)

	group.POST("/login", func(context *gin.Context) {
		user := common.User{}
		err := context.ShouldBind(&user)
		if err != nil {
			config.CatchErr(err)
			context.JSON(http.StatusForbidden, gin.H{
				"message": err,
			})
			return
		}

		res, err := db.LoginUser(&user)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"token":    res,
			"username": user.Username,
		})
	})
}
