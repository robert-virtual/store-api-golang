package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func openDB() {
	config := mysql.Config{
		User:   os.Getenv("DBUSER"), // obtener variable de entorno
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "ec2-3-141-200-64.us-east-2.compute.amazonaws.com",
		DBName: "gostore",
	}
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Conectado a base de datos")
}
