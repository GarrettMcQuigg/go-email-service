package internal

type Configuration struct {
	SenderEmail string `mapstructure:"senderEmail"`
	SenderPassword string `mapstructure:"senderPassword"`
	SmtpServer string `mapstructure:"smtpServer"`
}