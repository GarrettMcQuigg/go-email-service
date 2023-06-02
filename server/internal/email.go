package internal

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/spf13/viper"
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

func composeEmail() (string, string) {
	v := viper.New()
	v.SetConfigFile("C:/Users/gmati/OneDrive/Desktop/go-email-service/server/config.yml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
	}

	senderEmail := v.GetString("server.senderEmail")

	reqBody := RequestBody{
		Recipient:   "garrettmcquigg@gmail.com",
		Subject:     "Email sent using Golang SMTP",
		Body:        "Test Email.",
		TemplateRef: "temp1",
	}

	recipient := reqBody.Recipient

	message := "From: " + senderEmail + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + reqBody.Subject + "\n" +
		"MIME-Version: 1.0" + "\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"" + "\n" +
		"\n" +
		reqBody.Body

	return message, recipient
}

func SendEmail(senderEmail string, recipient string) error {
	v := viper.New()
	v.SetConfigFile("C:/Users/gmati/OneDrive/Desktop/go-email-service/server/config.yml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
	}

	senderPassword := v.GetString("server.senderPassword")
	smtpServer := v.GetString("server.smtpServer")

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	message, recipient := composeEmail()

	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, senderEmail, []string{recipient}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
