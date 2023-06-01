package email

import (
	"net/smtp"
	"strconv"
)

type Email struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"message"`
	TemplateRef      string `json:"tempalteRef"`
}

const (
	senderEmail 	= "garrettmcquigg@gmail.com"
	recipientEmail  = "garrettmcquigg@gmail.com"
	senderPassword 	= "aatwtqcapsdfvveh"
	smtpServer 		= "smtp.gmail.com"
	smtpPort 		= 587
)

func composeEmail() string {
	recipient := "garrettmcquigg@gmail.com"
	subject := "Test Email"
	body := "This is the email body."

	message := "From: " + senderEmail + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0" + "\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"" + "\n" +
		"\n" +
		body

	return message
}

func sendEmail() error {
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	message := composeEmail()

	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
