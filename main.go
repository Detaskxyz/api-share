package main

import (
	"api-share/http"
	"api-share/websocket"
	"github.com/gin-gonic/gin"
)

func main() {
	bindAddress := "localhost:2303"
	r := gin.Default()
	// ws://127.0.0.1:2303/ping
	r.GET("/ping", websocket.Ping) // WebSocket
	r.POST("/login", http.Login)   // 登录
	r.GET("/get", http.Get)        // 登录
	r.GET("/loginOut", http.LoginOut)
	r.Run(bindAddress)
}
