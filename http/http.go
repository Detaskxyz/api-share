package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录
type LoginReq struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func Login(c *gin.Context) {
	var login LoginReq
	_ = c.ShouldBindJSON(&login)
	if login.UserName == "admin" && login.PassWord == "123456" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNzk2NjBjNTItNWUwNS00Mzc5LWE3NWItNjQxYWUyZmYzOTJlIiwiSUQiOjEsIlVzZXJuYW1lIjoibGlhbmdqaWVzIiwiTmlja05hbWUiOiIiLCJBdXRob3JpdHlJZCI6IiIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2Njg0MzU5MTgsImlzcyI6InFtdFBsdXMiLCJuYmYiOjE2Njc4MzAxMTh9.9Tiats2vGBTgAwF65zZY9am1bGvgBle3iC9NuqsIJG4", "msg": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 7, "msg": "登录失败"})
	}
}

// 获取文章
func Get(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{"code": 7, "msg": "未登录"})
		return
	}
	if token == "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNzk2NjBjNTItNWUwNS00Mzc5LWE3NWItNjQxYWUyZmYzOTJlIiwiSUQiOjEsIlVzZXJuYW1lIjoibGlhbmdqaWVzIiwiTmlja05hbWUiOiIiLCJBdXRob3JpdHlJZCI6IiIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2Njg0MzU5MTgsImlzcyI6InFtdFBsdXMiLCJuYmYiOjE2Njc4MzAxMTh9.9Tiats2vGBTgAwF65zZY9am1bGvgBle3iC9NuqsIJG4" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "数据", "msg": "获取成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 7, "msg": "认证信息失效"})
	}
}

// 退出登录
func LoginOut(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{"code": 7, "msg": "未登录"})
		return
	}
	if token == "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNzk2NjBjNTItNWUwNS00Mzc5LWE3NWItNjQxYWUyZmYzOTJlIiwiSUQiOjEsIlVzZXJuYW1lIjoibGlhbmdqaWVzIiwiTmlja05hbWUiOiIiLCJBdXRob3JpdHlJZCI6IiIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2Njg0MzU5MTgsImlzcyI6InFtdFBsdXMiLCJuYmYiOjE2Njc4MzAxMTh9.9Tiats2vGBTgAwF65zZY9am1bGvgBle3iC9NuqsIJG4" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "退出成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 7, "msg": "退出失败"})
	}
}
