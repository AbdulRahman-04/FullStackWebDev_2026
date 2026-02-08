package utils

import "log"

var EmailQueue = make(chan EmailData, 50)

func StartEmailWorker(){
	go func() {
		for job := range EmailQueue {
			err := SendEmail(job);
			if err != nil {
				log.Printf("err %s", err)
			}
		}
	}()
}

var SmsQueue = make(chan SMSData, 50)

func StartSMSWorker() {
	go func() {
		for job := range SmsQueue {
			err := SendSMS(job)
			if err != nil {
				log.Printf("err %s", err)
			}
		}
	}()
}