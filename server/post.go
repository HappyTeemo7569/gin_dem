package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostMsg(c *gin.Context) {
	//name:=c.Query("name")//获取URL中的数据
	name := c.DefaultPostForm("name", "Gin") //获取body中的数据（form形式）
	fmt.Println(name)
	form, b := c.GetPostForm("name") //获取body中的数据（form形式）
	fmt.Println(form, b)
	c.JSON(http.StatusOK, "欢迎您："+name)
}
