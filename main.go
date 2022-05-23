package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectToDB()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	products := router.Group("/products")
	{
		products.GET("", getProduct)
		products.POST("", postProduct)
	}
	users := router.Group("")
	{
		users.GET("/users", getUsers)
		// estas dos no son accedibles desde el navegador xq el navegar hace peticiones get
		users.POST("/login", login)
		users.POST("/register", register)
	}
	router.Run()
}
