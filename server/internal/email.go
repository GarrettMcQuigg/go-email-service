package internal

import (
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Recipient   string `json:"recipient"`
	Subject     string `json:"subject"`
	Body        string `json:"message"`
	TemplateRef string `json:"tempalteRef"`
}

const (
	smtpPort = 587
)

var config Configuration

func ComposeEmail(ctx *gin.Context) {
	// con, exists := ctx.Get("config")
	// if !exists {
	// 	ctx.JSON(400, "error, bad request")
	// }
	var requestBody RequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	emailResponse := RequestBody{
		Recipient: requestBody.Recipient,
		Subject:   requestBody.Subject,
		Body:      requestBody.Body,
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

	ctx.JSON(http.StatusOK, gin.H{"email": emailResponse, "success": "success"})
}
