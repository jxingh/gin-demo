package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认路由
	router := gin.Default()
	// 绑定路由规则何路由函数，访问/index 的路由，将有对应的函数去处理
	router.GET("/index", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	// 启动监听，gin 把web 服务运行在本机的0.0.0.0:8080 端口上
	router.Run(":8080")
	// 用原生http 服务的方式，router.Run() 本质就是http.ListenAndServe() 的封装
	// http.ListenAndServe(":8080", router)
}
