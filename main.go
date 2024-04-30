package main

import (
	"capybara-go/config"
	"capybara-go/prompt"
	"capybara-go/wenxin"
	"fmt"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type ChatRequest struct {
	Message string `json:"message"`
}

func chat(c *gin.Context) {
	var chatContent ChatRequest
	if err := c.ShouldBindJSON(&chatContent); err == nil {
		query := chatContent.Message
		result := wenxin.Chat(prompt.BuildCapyPrompt(query))
		c.JSON(200, gin.H{"status": "ok", "data": prompt.GetJSONObj(result)})
	} else {
		// 请求体解析失败，处理错误
		c.JSON(400, gin.H{"error": err.Error()})
	}
}

type CapybaraBehavior struct {
	Emotion  string `json:"emotion"`  // 水豚的表情
	Movement int    `json:"movement"` // 水豚的位移
	Action   string `json:"action"`   // 水豚的动作
}

func main() {
	port := config.GlobalConfig.Server.Port
	r := gin.Default()
	r.GET("/ping", pong)
	r.POST("/chat", chat)
	r.Run(fmt.Sprintf(":%d", port))
}
