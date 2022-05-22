package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

// con fines de desarooolo y para porder visualizar los usuarios vamos a crear esta ruta pero la deshabiliatermos posteriormente
func getUsers(ctx *gin.Context) {
	users, err := findUsers()

	if err != nil {
		ctx.JSON(500, gin.H{
			"msg": err,
		})
		return // para que no continue con el codigo
	}
	ctx.JSON(200, users)

}
func login(ctx *gin.Context) {
	var userRecived user
	err := ctx.BindJSON(&userRecived) // solo necesitamos email y passsword
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": err,
		})
		return // para que no continue con el codigo
	}
	user, error := findUser(userRecived.Email)
	if error != nil {
		ctx.JSON(401, error)
		fmt.Println(error)
		return
	}
	// desencriptar la password y cmpararla con la password enviada por el usuario
	ok, err := argon2.VerifyEncoded([]byte(userRecived.Password), []byte(user.Password))
	if !ok {
		ctx.JSON(401, error)
		return
	}
	// enviar cookies o un json web token ya que los datos son correctos
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.Id,
		"nbf": time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(201, gin.H{
		"user":  user,
		"token": tokenString,
	})

}

//post user
func register(ctx *gin.Context) {
	var user user
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err)
		return // para que no continue con el codigo
	}
	_, error := createUser(user)
	if error != nil {

		ctx.JSON(400, error)
		return
	}
	ctx.JSON(201, user)
}
