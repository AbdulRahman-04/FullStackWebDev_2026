package utils

import (
	"log"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	From    string
	To      string
	Subject string
	Text    string
	Html    string
}

func SendEmail(emailData EmailData) error {
	// get email nd app pass
	user := config.AppConfig.EMAIL
	pass := config.AppConfig.PASS

	// create sender
	s := gomail.NewMessage()

	s.SetAddressHeader("From", user, "team social")
	s.SetHeader("To", emailData.To)
	s.SetHeader("Subject", emailData.Subject)
	s.SetBody("text/plain", emailData.Text)
	s.AddAlternative("text/html", emailData.Html)

	// create dialer
	t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

	// send email
	if err := t.DialAndSend(s); err != nil {
		log.Printf("Couldn't send email")
		return  err
	}

	log.Printf("email sent✅")

	return nil
}
