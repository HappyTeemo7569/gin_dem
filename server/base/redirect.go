package base

import (
	"gin_demo/server/define"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect1(c *gin.Context) {
	//重定向到到百度，获取百度对应的数据；
	//重定向状态码：StatusMovedPermanently
	url := "http://www.baidu.com" //重定向的URL
	c.Redirect(http.StatusMovedPermanently, url)
}

func Redirect2(c *gin.Context) {
	c.Request.URL.Path = "/TestRedirect" //重定向的路由
	define.GinEngine.HandleContext(c)
}
func TestRedirect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "响应成功",
	})
}
