package http

import (
	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	r.Group("/api")
	//sys.Article(r)
}

func SysApi(r *gin.Engine) {
	r.Group("/")
	//sys.UserSys(r)
}
