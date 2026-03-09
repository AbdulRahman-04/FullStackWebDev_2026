package utils

import (
	"log"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	From string
	To string
	Subject string
    Text string
	Html string
}

func SendEmail(emailData EmailData) error {
	user := config.AppConfig.EMAIL
	pass := config.AppConfig.PASS

	s := gomail.NewMessage()

	s.SetAddressHeader("From", user, "Team Follow")
	s.SetHeader("To", emailData.To)
	s.SetHeader("Subject", emailData.Subject)
	s.SetBody("text/plain", emailData.Text)
	s.AddAlternative("text/html", emailData.Html)

	t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

	if err := t.DialAndSend(s); err != nil {
		log.Printf("Couldn't send email")
	}

	log.Printf("Email Sent✅")
    return  nil
}