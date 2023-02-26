package api

import (
	"easy_file/src/config"
	"easy_file/src/database"
	"easy_file/src/redis"
	"easy_file/src/structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Article(r *gin.Engine) {
	r.POST("/redisDemo", func(c *gin.Context) {
		kv := config.KV{}
		err := c.ShouldBind(&kv)

		res, err := redis.RedisDemo(kv)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"newRes": res,
		})
	})

	r.POST("/userCreate", func(c *gin.Context) {
		user := structs.User{}
		err := c.ShouldBind(&user)
		fmt.Println(user)
		err = database.UserCreate(user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.POST("/batchCreate", func(c *gin.Context) {
		var users []structs.User
		err := c.ShouldBind(&users)
		fmt.Println(users)
		err = database.BatchCreate(users)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.POST("/articleCreate", func(c *gin.Context) {
		user := structs.User{}
		err := c.ShouldBind(&user)
		fmt.Println(user)
		err = database.UserCreate(user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.GET("/getUser", func(c *gin.Context) {
		user := structs.User{
			Name: c.Param("name"),
		}
		err := c.ShouldBind(&user)
		database.SelectDemo(&user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.GET("/getUsers", func(c *gin.Context) {
		var users []structs.User
		name := c.Query("name")
		err := database.SelectDemo2(&users, name)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.PUT("/updateUsers", func(c *gin.Context) {

		user := structs.User{}
		err := c.ShouldBind(&user)
		err = database.UpdateDemo(&user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.DELETE("/deleteUsers", func(c *gin.Context) {

		user := structs.User{}
		err := c.ShouldBind(&user)
		err = database.DeleteDemo(&user)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}
