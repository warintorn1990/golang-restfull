package main

import (
	"net/http"
	"os"
	"fmt"
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
	runningDir, _ := os.Getwd()

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "arm",
		})
	})

	r.GET("/book/:from/:to", handleRequest)

	r.Run()
}
