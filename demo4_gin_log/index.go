package main

import (
	"net/http"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	runningDir, _ := os.Getwd()
	count := 0

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile
	// r.Use(gin.Logger())

	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s  \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/profile"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "arm",
		})
	})

	r.GET("/error", func(c *gin.Context) {
		count = count + 1
		errlogfile.WriteString(fmt.Sprintf("Error : %d\n", count))
		c.Data(200, "text/html; charset=utf-8", []byte("error"))
	})

	r.GET("/profile", func(c *gin.Context) {
		count = count + 1
		accesslogfile.WriteString(fmt.Sprintf("Count : %d\n", count))
		c.JSON(http.StatusOK, gin.H{
			"message": "profile",
		})
	})

	r.Run()
}
