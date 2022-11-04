package userserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件开始执行=====")
		name := c.PostForm("name")
		ageStr := c.PostForm("age")
		age, err := strconv.Atoi(ageStr) //string转int
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "输入的数据错误，年龄不是整数")
			return
		}
		if age < 0 || age > 100 {
			c.AbortWithStatusJSON(http.StatusBadRequest, "输入的数据错误，年龄数据错误")
			return
		}
		if len(name) < 6 || len(name) > 12 {
			c.AbortWithStatusJSON(http.StatusBadRequest, "用户名只能是6-12位")
			return
		}
		c.Next() //执行后续操作
		fmt.Println(name, age)

	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件2开始执行=====")

		c.Next() //执行后续操作

		return
	}
}
