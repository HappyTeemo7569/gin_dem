package videosever

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VideoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件开始执行=====")

		videoIdStr := c.Query("videoId")
		videoId, err := strconv.Atoi(videoIdStr) //string转int
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "输入的数据错误，通话Id不是整数")
			return
		}

		if videoId <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, "输入的数据错误，通话Id数据错误")
			return
		}

		c.Next() //执行后续操作
		fmt.Println(videoId)

	}
}
