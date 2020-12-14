package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", getFromWeb())
	r.POST("/post", postFromApp("hoge"))
	r.Run()
}

func postFromApp(arg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"arg": arg,
		})
	}
}

// Webが呼ぶやつ
func getFromWeb() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
