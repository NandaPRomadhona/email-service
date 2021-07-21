package main

import (
	"fmt"
	"github.com/email-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//ganti env jadi shell
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	fmt.Println("======Send Email Service====")

	route := gin.Default()
	route.POST("/user-validation", services.UserValidationEmailHandler)
	route.Run("0.0.0.0:9999")

}