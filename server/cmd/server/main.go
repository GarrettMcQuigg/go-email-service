package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"server/internal"
)

func main() {
	router := gin.Default()

	v := viper.New()
	v.SetConfigFile("config.yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
        return
	}

	serverURL := v.GetString("server.host")
	serverPort := v.GetString("server.port")
	fmt.Println("host server URL:", serverURL + ":" + serverPort)

	err := internal.SendEmail()
	if err != nil {
		fmt.Println("Failed to send email:", err)
        return
	}

	router.Run(":8080")
}