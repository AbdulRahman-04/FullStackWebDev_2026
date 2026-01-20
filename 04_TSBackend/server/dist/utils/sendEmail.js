"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const nodemailer_1 = __importDefault(require("nodemailer"));
const config_1 = __importDefault(require("config"));
const USER = config_1.default.get("USER");
const PASS = config_1.default.get("PASS");
async function SendEmail(emailData) {
    try {
        let transporter = nodemailer_1.default.createTransport({
            host: "smtp.gmail.com",
            port: 465,
            secure: true,
            auth: {
                user: USER,
                pass: PASS
            }
        });
        let sender = await transporter.sendMail({
            from: emailData.from,
            to: emailData.to,
            subject: emailData.subject,
            text: emailData.text,
            html: emailData.html
        });
        console.log("email sent✅", `${emailData.to}: ${sender.messageId}`);
    }
    catch (error) {
        console.log(error);
    }
}
exports.default = SendEmail;
