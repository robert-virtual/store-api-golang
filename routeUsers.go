package main

import "github.com/gin-gonic/gin"

// con fines de desarooolo y para porder visualizar los usuarios vamos a crear esta ruta pero la deshabiliatermos posteriormente
func getUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "getUsers",
	})

}
func login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "login",
	})

}

//post user
func register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "crear user",
	})
}
