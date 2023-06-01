package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"server/internal"
)

type Email struct {
	Recipient 		string `json:"recipient"`
	Subject   		string `json:"subject"`
	Body      		string `json:"message"`
	TemplateRef     string `json:"tempalteRef"`
}

const (
	senderEmail 	= "garrettmcquigg@gmail.com"
	recipientEmail  = "garrettmcquigg@gmail.com"
	senderPassword 	= "aatwtqcapsdfvveh"
	smtpServer 		= "smtp.gmail.com"
	smtpPort 		= 587
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