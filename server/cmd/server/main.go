package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"server/internal"

	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()

	v := viper.New()
	v.SetConfigFile("C:/Users/gmati/OneDrive/Desktop/go-email-service/server/config.yml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
	}

	senderEmail := v.GetString("server.senderEmail")
	recipient := "garrettmcquigg@gmail.com"

	err := internal.SendEmail(senderEmail, recipient)
	if err != nil {
		fmt.Println("Failed to send email:", err)
        return
	}

	router.Run(":8080")
}