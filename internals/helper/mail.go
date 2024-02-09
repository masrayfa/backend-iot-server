package helper

import (
	"github.com/masrayfa/configs"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	config := configs.GetConfig()

	// send email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.Mail.SenderName)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(config.Mail.SMTPHost, config.Mail.SMTPPort, config.Mail.AuthenticationMail, config.Mail.AuthenticationPassword)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}