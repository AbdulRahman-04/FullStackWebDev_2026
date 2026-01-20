import twilio from "twilio"
import config from "config"


const SID: string = config.get<string>("SID")
const TOKEN: string = config.get<string>("TOKEN")
const PHONE: string = config.get<string>("PHONE")

interface SmsData {
    body : string,
    to: string
}

let client = new twilio.Twilio(SID, TOKEN)


async function SendSMS(smsData: SmsData) {
    try {
 
        await client.messages.create({
            body: smsData.body,
            to: smsData.to,
            from: PHONE
        })

        console.log(`sms sent✅`);
        

        
    } catch (error) {
        console.log(error);
        
    }
    
}

export default SendSMS;