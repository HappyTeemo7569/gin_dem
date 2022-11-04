package userserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserServer() http.Handler {
	r1 := gin.Default()
	r1.Use(UserMiddleware(), AuthMiddleware())
	r1.POST("/UserServer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "用户服务器程序1的相应",
		},
		)
	})
	r1.POST("/login", Login)
	return r1
}
