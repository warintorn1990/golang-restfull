package main

import (
	"demo8/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login", api.Login)
		authenAPI.GET("/register", api.Register)		
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("/list", api.ListProduct)
		stockAPI.GET("/create", api.CreateProduct)
	}

	router.Run("")
}