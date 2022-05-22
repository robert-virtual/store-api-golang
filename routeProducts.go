package main

import "github.com/gin-gonic/gin"

func getProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "lista de productos",
	})
}

func postProduct(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"msg": "crear producto",
	})
}
