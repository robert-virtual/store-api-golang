package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getProduct(ctx *gin.Context) {
	prods, err := findProducts()
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, prods)
}

func postProduct(ctx *gin.Context) {
	var product product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(400, err)
		fmt.Println(err)
		return // para que no continue con el codigo
	}
	_, error := createProduct(product)
	if error != nil {
		ctx.JSON(500, error)
		return
	}

	ctx.JSON(201, gin.H{
		"msg":     "producto creado",
		"product": product,
	})
}
