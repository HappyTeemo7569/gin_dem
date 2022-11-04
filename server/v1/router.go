package v1

import (
	"gin_demo/server/define"
)

func InitRouterV1() {
	v1 := define.GinEngine.Group("/v1")

	r := v1.Group("/user") //路由分组（2级路径）
	r.GET("/login", login) // 响应请求:/v1/user/login

	r2 := r.Group("/showInfo")    //路由分组（3级路径）
	r2.GET("/abstract", abstract) //响应请求：v1/user/showInfo/abstract
	r2.GET("/detail", detail)     //响应请求：v1/user/showInfo/detail
}
