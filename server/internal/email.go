package internal

import (
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Recipient   string `json:"recipient" binding:"required"`
	Subject     string `json:"subject" binding:"required"`
	Body        string `json:"body" binding:"required"`
}

const (
	smtpPort = 587
)

func ComposeEmail(ctx *gin.Context) {
	var requestBody RequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conf, exists := ctx.Get("config")
	if !exists {
		ctx.JSON(400, "error, bad request")
		return
	}

	config, ok := conf.(Configuration)
	if !ok {
		ctx.JSON(400, "error, bad request")
		return
	}

	message := "From: " + config.SenderEmail + "\n" +
		"To: " + requestBody.Recipient + "\n" +
		"Subject: " + requestBody.Subject + "\n" +
		"MIME-Version: 1.0" + "\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"" + "\n" +
		"\n" +
		requestBody.Body

	auth := smtp.PlainAuth("", config.SenderEmail, config.SenderPassword, config.SmtpServer)

	err := smtp.SendMail(config.SmtpServer+":"+strconv.Itoa(smtpPort), auth, config.SenderEmail, []string{requestBody.Recipient}, []byte(message))
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}
