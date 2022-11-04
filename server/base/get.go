package base

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

func GetOtherData(c *gin.Context) {
	url := "http://www.baidu.com"
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable) //应答client
		return
	}
	body := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	//数据写入响应体
	c.DataFromReader(http.StatusOK, contentLength, contentType, body, nil)
}

func Json(c *gin.Context) {
	c.JSON(200, gin.H{
		"html": "<b>hello,Gin框架</b>",
	})
}

func SomeHTML(c *gin.Context) {
	c.PureJSON(http.StatusOK, gin.H{
		"html": "<b>hello,Gin框架</b>",
	})
}

func SomeXML(c *gin.Context) {
	type Message struct {
		Name string
		Msg  string
		Age  int
	}
	info := Message{}
	info.Name = "阿晓"
	info.Msg = "hello"
	info.Age = 23
	c.XML(http.StatusOK, info)
}

func SomeYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"message": "Gin框架的多形式渲染",
		"status":  200,
	})
}
