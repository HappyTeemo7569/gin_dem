package userserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserServer() http.Handler {
	r1 := gin.Default()
	r1.Use(UserMiddleware())
	r1.GET("/UserServer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "用户服务器程序1的相应",
		},
		)
	})
	return r1
}
