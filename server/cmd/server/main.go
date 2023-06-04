package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"server/internal"
)

func main() {
	router := gin.Default()

	config := internal.Configuration{
		SenderEmail: "",
		SenderPassword: "",
		SmtpServer: "",
	}

	viperRef := viper.New()
	viperRef.SetConfigName("config")
	viperRef.SetConfigType("yml")
	viperRef.AddConfigPath(".")
	viperRef.AddConfigPath("../../")

	viperRef.AutomaticEnv()

	if err := viperRef.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viperRef.Unmarshal(&config); err != nil {
		panic(err)
	}

	router.Use(func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	})

	router.POST("/send-email", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "successfully sent email"})
	})
	
	router.Run()
}