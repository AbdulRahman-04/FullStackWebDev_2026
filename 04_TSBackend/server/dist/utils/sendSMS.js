"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const twilio_1 = __importDefault(require("twilio"));
const config_1 = __importDefault(require("config"));
const SID = config_1.default.get("SID");
const TOKEN = config_1.default.get("TOKEN");
const PHONE = config_1.default.get("PHONE");
let client = new twilio_1.default.Twilio(SID, TOKEN);
async function SendSMS(smsData) {
    try {
        await client.messages.create({
            body: smsData.body,
            to: smsData.to,
            from: PHONE
        });
        console.log(`sms sent✅`);
    }
    catch (error) {
        console.log(error);
    }
}
exports.default = SendSMS;
