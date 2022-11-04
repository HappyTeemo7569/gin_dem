package server

func InitRouter() {
	GinEngine.GET("/get", GetMsg) //get方法
}
