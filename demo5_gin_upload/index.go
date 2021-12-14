package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	runningDir, _ := os.Getwd()
	count := 0

	router.MaxMultipartMemory = 8 << 20

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "arm",
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		// single file
		count = count + 1
		file, _ := c.FormFile("file")
		extension := filepath.Ext(file.Filename)

		c.SaveUploadedFile(file, fmt.Sprintf("%s/images/%d%s", runningDir, count, extension))
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run("")
}