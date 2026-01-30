// package utils

// import (
// 	"fmt"

// 	"gopkg.in/gomail.v2"
// 	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
// )

// type EmailData struct {
// 	To string
// 	From string
// 	Html string
// 	Subject string
// 	Text string
// }

// func SendEmail(data EmailData) error {
// 	user := config.AppConfig.Email
// 	pass := config.AppConfig.Pass

// 	s := gomail.NewMessage()
// 	s.SetAddressHeader("From", user, "Team Insta")
// 	s.SetHeader("To", data.To)
// 	s.SetHeader("Subject", data.Subject)
// 	s.SetBody("Text/plain", data.Text)
// 	s.AddAlternative("text/html", data.Html)

// 	t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

// 	if err := t.DialAndSend(s); err != nil {
// 		return err // ❗ yahin galti thi
// 	}

// 	fmt.Println("email sent✅")
// 	return nil
// }

package utils

import (
	"log"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
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
 
	// get user and pass
	user := config.AppConfig.Email
	pass := config.AppConfig.Pass

	// create sender 
	s := gomail.NewMessage()

	s.SetAddressHeader("From", user, "Team social")
	s.SetHeader("To", emailData.To)
	s.SetHeader("Subject", emailData.Subject)
	s.SetBody("text/plain", emailData.Text)
	s.AddAlternative("text/html", emailData.Html)

	// create transporter
	t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

	if err := t.DialAndSend(s); err  != nil {
		return  err
	}

	log.Printf("email sent✅")
	return  nil

}