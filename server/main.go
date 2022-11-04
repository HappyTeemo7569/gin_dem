package server

import "github.com/gin-gonic/gin"

var GinEngine *gin.Engine

func init() {
	GinEngine = gin.Default()
}

func Run() {
	InitRouter()
	GinEngine.Run(":9090") //如果不指定IP地址、端口号，默认为本机IP地址、8080端口
}
