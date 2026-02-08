package utils

import (
	"log"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSData struct {
	To   string
	Body string
}

func SendSMS(smsData SMSData) error {
	// create client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AppConfig.SID,
		Password: config.AppConfig.TOKEN,
	})

	// send sms
	_, err := client.Api.CreateMessage(&openapi.CreateMessageParams{
		From: &config.AppConfig.PHONE,
		To:   &smsData.To,
		Body: &smsData.Body,
	})

	if err != nil {
		log.Printf("couldn't send sms")
		return err
	}

	log.Printf("Sms sent✅")
	return nil
}
