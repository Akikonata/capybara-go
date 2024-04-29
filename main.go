package main

import (
	"capybara-go/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type CapybaraBehavior struct {
	Emotion  string `json:"emotion"`  // 水豚的情绪
	Movement int    `json:"movement"` // 水豚的移动方向
	Action   string `json:"action"`   // 水豚的动作
}

func main() {
	port := config.GlobalConfig.Server.Port
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run(fmt.Sprintf(":%d", port))
}
