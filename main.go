package main

import (
	"log"
	"module/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	r := gin.Default()
	r.GET("/qris", handler.QrisHandler)
	r.Run(":8000")
}
