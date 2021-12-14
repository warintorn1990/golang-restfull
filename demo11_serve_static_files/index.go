package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/image", func(c *gin.Context) {
		c.File("static/test.jpeg")
	})

	r.GET("/html", func(c *gin.Context) {
		c.File("static/index.html")
	})

	r.GET("/download", func(c *gin.Context) {
		c.Header("Content-Description", "Simulation File Download")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+"แมว.jpeg")
		c.Header("Content-Type", "application/octet-stream")

		c.File("static/test.jpeg")
	})

	r.Run(":85")

}