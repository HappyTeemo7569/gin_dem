package videosever

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideoServer() http.Handler {
	r1 := gin.Default()
	r1.Use(VideoMiddleware())
	r1.GET("/VideoServer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "视频服务器程序1的相应",
		},
		)
	})
	return r1
}
