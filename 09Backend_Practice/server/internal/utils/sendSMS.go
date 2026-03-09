package utils

import (
	"log"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSData struct {
	To string
	Body string
}

func SendSMS(smsData SMSData) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AppConfig.SID,
		Password: config.AppConfig.TOKEN,
	})

	// send sms 
	_, err := client.Api.CreateMessage(&openapi.CreateMessageParams{
		From: &config.AppConfig.PHONE,
		To: &smsData.To,
		Body: &smsData.Body,
	})

	if err != nil {
		log.Printf("err %v", err)
	}

	log.Printf("sms sent✅")
	return  nil
}