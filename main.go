package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu,fixbugs")
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
