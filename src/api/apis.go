package api

import (
	"easy_file/src/api/sys"
	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	r.Group("/api")
	sys.Article(r)
}

func SysApi(r *gin.Engine) {
	r.Group("/")
	sys.UserSys(r)
}
