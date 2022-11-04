package server

func InitRouter() {
	GinEngine.GET("/get", GetMsg)    //get方法
	GinEngine.POST("/post", PostMsg) //post方法
}
