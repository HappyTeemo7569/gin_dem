package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMsg(c *gin.Context) {
	name := c.Query("name")
	//返回字符串类型数据
	//c.String(http.StatusOK,"欢迎您：%s",name)
	//返回json数据
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, //状态
		"msg":  "返回信息",        //描述信息
		"data": "欢迎您：" + name, //数据
	})
}
