// package utils

// import "log"

// // ---------------- EMAIL WORKER ----------------

// var EmailQueue = make(chan EmailData, 50)

// // StartEmailWorker starts 1 background worker for emails
// func StartEmailWorker() {
// 	go func() {
// 		for job := range EmailQueue {
// 			if err := SendEmail(job); err != nil {
// 				log.Println("email error:", err)
// 			}
// 		}
// 	}()
// }

// // ---------------- SMS WORKER ----------------

// var SMSQueue = make(chan SMSdata, 50)

// // StartSMSWorker starts 1 background worker for sms
// func StartSMSWorker() {
// 	go func() {
// 		for job := range SMSQueue {
// 			if err := SendSMS(job); err != nil {
// 				log.Println("sms error:", err)
// 			}
// 		}
// 	}()
// }

package utils

import "log"

var EmailQueue = make(chan EmailData, 50)

func StartEmailWorker(){
	go func() {
		for job := range EmailQueue {
			err := SendEmail(job);
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

var SMSQueue = make(chan SMSdata, 50)

func StartSMSWorker(){
	go func() {
		for job := range SMSQueue {
			err := SendSMS(job)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}