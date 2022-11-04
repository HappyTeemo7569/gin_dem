package server

import (
	"fmt"
	"gin_demo/server/base"
	"gin_demo/server/define"
	"gin_demo/server/userserver"
	v1 "gin_demo/server/v1"
	"gin_demo/server/videosever"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter() http.Handler {
	define.GinEngine = gin.Default()
	v1.InitRouterV1()
	base.InitRouterBase()
	return define.GinEngine
}

func RunPlus() {
	//服务器0：
	server00 := &http.Server{
		Addr:         ":9090",
		Handler:      InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//服务器1：
	server01 := &http.Server{
		Addr:         ":9091",
		Handler:      userserver.UserServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//服务器2：
	server02 := &http.Server{
		Addr:         ":9092",
		Handler:      videosever.VideoServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//开启服务
	define.GroupServer.Go(func() error { //开启服务器程序0
		return server00.ListenAndServe()
	})
	define.GroupServer.Go(func() error { //开启服务器程序1
		return server01.ListenAndServe()
	})
	define.GroupServer.Go(func() error { //开启服务器程序2
		return server02.ListenAndServe()
	})
	if err := define.GroupServer.Wait(); err != nil {
		fmt.Println("执行失败：", err)
	}
}
