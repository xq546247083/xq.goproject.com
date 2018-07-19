package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载默认配置
	router := gin.Default()
	router.Use(Append())
	initHtmlLoad(router)
	initApiLoad(router)

	router.Run(":8080")
}

// 测试中间件
func Append() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Write([]byte("hello"))
		c.Next()
	}
}

// 初始化html加载
func initHtmlLoad(router *gin.Engine) {
	// 加载hmtl静态文件
	router.LoadHTMLGlob("html/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "1.html", gin.H{
			"title": "GIN: 测试加载HTML模板",
		})
	})
}

// 初始化Api加载
func initApiLoad(router *gin.Engine) {
	router.GET("/welcome", func(c *gin.Context) {
		name := c.Query("name")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", name, lastname)
	})
}
