package internal

import (
	"net/smtp"
	"strconv"
)

type RequestBody struct {
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

func composeEmail() string {
	reqBody := RequestBody{
		Recipient: recipientEmail,
		Subject: "Email sent using Golang SMTP",
		Body: "Test Email.",
		TemplateRef: "temp1",
	}

	message := "From: " + senderEmail + "\n" +
		"To: " + recipientEmail + "\n" +
		"Subject: " + reqBody.Subject + "\n" +
		"MIME-Version: 1.0" + "\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"" + "\n" +
		"\n" +
		reqBody.Body

	return message
}

func SendEmail() error {
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	message := composeEmail()

	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
