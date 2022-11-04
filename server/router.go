package server

func InitRouter() {
	GinEngine.GET("/get", GetMsg)    //get方法
	GinEngine.POST("/post", PostMsg) //post方法

	//一般重定向:重定向到外部网络
	GinEngine.GET("/redirect1", Redirect1)
	//路由重定向：重定向到具体的路由
	GinEngine.GET("/redirect2", Redirect2)
	//路由：127.0.0.1:9090/TestRedirect
	GinEngine.GET("/TestRedirect", TestRedirect)
}
