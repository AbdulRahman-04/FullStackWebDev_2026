package utils

import "log"

var EmailQueue = make(chan EmailData, 50)

func StartEmailWorker() {
	go func() {
		for job := range EmailQueue {
			err := SendEmail(job)
			if err != nil {
				log.Printf("Couldnn't send email %v", err)
			}
		}
	}()
}

var SMSQueue = make(chan SMSData, 50)

func StartSMSWorker() {
	go func() {
		for job := range SMSQueue {
			err := SendSMS(job)
			if err != nil {
				log.Printf("coudln't send sms %v", err)
			}
		}
	}()
}