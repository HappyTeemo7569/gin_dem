package base

import (
	"gin_demo/server/define"
)

func InitRouterBase() {

	define.GinEngine.GET("/get", GetMsg)    //get方法
	define.GinEngine.POST("/post", PostMsg) //post方法

	//一般重定向:重定向到外部网络
	define.GinEngine.GET("/redirect1", Redirect1)
	//路由重定向：重定向到具体的路由
	define.GinEngine.GET("/redirect2", Redirect2)
	//路由：127.0.0.1:9090/TestRedirect
	define.GinEngine.GET("/TestRedirect", TestRedirect)

	define.GinEngine.GET("/GetOtherData", GetOtherData)

	define.GinEngine.GET("/Json", Json)
	define.GinEngine.GET("/SomeHTML", SomeHTML)
	define.GinEngine.GET("/SomeXML", SomeXML)
	define.GinEngine.GET("/SomeYAML", SomeYAML)
}
