package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleRequest(c *gin.Context) {
	from, to := c.Param("from"), c.Param("to")
	c.JSON(http.StatusOK, gin.H{"result": "ok", "from": from, "to": to})
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "arm",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
	})

	r.GET("/book/:from/:to", handleRequest)

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.Username == "admin" && form.Password == "1234" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(200, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(401, gin.H{"status": "unable to bind"})
		}
	})

	r.Run()
}
