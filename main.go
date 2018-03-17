package main

import (
	"github.com/branLiang/gin-playground/cmd"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化命令行工具集
	cmd.Execute()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
