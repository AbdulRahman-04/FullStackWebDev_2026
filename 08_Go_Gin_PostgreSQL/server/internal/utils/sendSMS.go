// package utils

// import (

// 	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
// 	"github.com/twilio/twilio-go"

// 	openapi "github.com/twilio/twilio-go/rest/api/v2010"
// )

// type SMSData struct {
// 	From string
// 	To string
// 	Body string
// }

// func SendSMS(data SMSData) error {
// 	// create client
// 	client := twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Username: config.AppConfig.SID,
// 		Password: config.AppConfig.Token,
// 	})

// 	// get data ready
// 	_, err := client.Api.CreateMessage(&openapi.CreateMessageParams{
// 		From: &config.AppConfig.Phone,
// 		To: &data.To,
// 		Body: &data.Body,
// 	})

// 	if err != nil {
// 		return  err
// 	}

// 	return nil
// }

package utils

import (
	"log"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSdata struct {
	From string
	To string
	Body string
}

func SendSMS(smsData SMSdata) error {

	// create client 
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AppConfig.SID,
		Password: config.AppConfig.Token,
	})

	_, err := client.Api.CreateMessage(&openapi.CreateMessageParams{
		From:  &config.AppConfig.Phone,
		To: &smsData.To,
		Body: &smsData.Body,
	})

	if err != nil {
		return  err
	}

	log.Printf("sms sent✅")
	return  nil
}