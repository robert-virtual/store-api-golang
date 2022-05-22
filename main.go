package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	openDB()
	router := gin.Default()
	products := router.Group("/products")
	{
		products.GET("/", getProduct)
		products.POST("/", postProduct)
	}
	users := router.Group("/")
	{
		users.GET("/users", getUsers)
		users.POST("/login", login)
		users.POST("/register", register)
	}
	router.Run()
}
