package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"server/internal"
)

func main() {
	router := gin.Default()

	err := internal.SendEmail()
	if err != nil {
		fmt.Println("Failed to send email:", err)
        return
	}

	router.Run(":8080")
}