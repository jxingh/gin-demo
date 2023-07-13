package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func _string(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func _json(c *gin.Context) {
	// json 响应结构体
	// type UserInfo struct {
	// 	UserName string `json:"user_name"`
	// 	Age      int    `json:"age"`
	// 	Password string `json:"-"` // 忽略转换为json，也可以将属性名改为小写
	// }
	// user := UserInfo{"luck", 22, "123456"}
	// c.JSON(http.StatusOK, user)

	// json 响应map
	// userMap := map[string]string{
	// 	"user_name": "luck",
	// 	"age":       "22",
	// }
	// c.JSON(http.StatusOK, userMap)

	// 直接响应json
	c.JSON(http.StatusOK, gin.H{"user_name": "luck", "age": 22})
}

func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "luck", "message": "hey", "status": http.StatusOK})
}

// func _yaml(c *gin.Context) {

// }

// func _html(c *gin.Context) {

// }

func main() {
	router := gin.Default()
	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.Run(":8080")
}
