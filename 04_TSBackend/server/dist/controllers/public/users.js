"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
const bcrypt_1 = __importDefault(require("bcrypt"));
const sendEmail_1 = __importDefault(require("../../utils/sendEmail"));
// import SendSMS from "../../utils/sendSMS"
const config_1 = __importDefault(require("config"));
const users_1 = require("../../models/users");
// make router to create apis 
const router = express_1.default.Router();
const URL = config_1.default.get("URL");
const KEY = config_1.default.get("JWT_KEY");
const USER = config_1.default.get("USER");
// signup api
router.post("/signup", async (req, res) => {
    try {
        // take data from req.body
        const { userName, email, password, phone, age } = req.body;
        if (!userName || !email || !password || !phone || !age) {
            res.status(400).json({ msg: "pls fill all fields" });
            return;
        }
        // duplicate check 
        let userExist = await users_1.userModel.findOne({ email });
        if (userExist) {
            res.status(200).json({ msg: "user already exists, pls signin" });
            return;
        }
        // pass hash 
        const passHash = await bcrypt_1.default.hash(password, 10);
        // random token generate for email nd sms
        let emailToken = Math.random().toString(36).substring(2);
        // let phoneToken = Math.random().toString(36).substring(2)
        // create new objext
        const newUser = {
            userName,
            email,
            password: passHash,
            phone,
            age,
            userVerifyToken: {
                emailVerifyToken: emailToken,
                // phoneVerifyToken: phoneToken
            }
        };
        // save user in db 
        await users_1.userModel.create(newUser);
        // email verification link
        const emailData = {
            from: USER,
            to: email,
            subject: "Verification Link",
            text: `${URL}/api/public/emailverify/${emailToken}`
        };
        (0, sendEmail_1.default)(emailData);
        // 8. Verification ke liye SMS data banate hain aur sendSMS function call karte hain.
        // const smsData = {
        //     body: `📲 Team Todo: Dear user, verify your phone by clicking the link: ${URL}/api/public/phoneverify/${phoneToken}. 
        //     If you didn't request this, ignore the message.`,
        //     to: phone
        // };
        // sendSMS(smsData);
        console.log(`${URL}/api/public/emailverify/${emailToken}`);
        // console.log(`${URL}/api/public/phoneverify/${phoneToken}`);
        res.status(200).json({ msg: "User signed up successfully! verify email and phone✨" });
    }
    catch (error) {
        res.status(401).json({ msg: error });
        console.log(error);
    }
});
// email verify api 
router.get("/emailverify/:token", async (req, res) => {
    try {
        const mytoken = req.params.token;
        if (typeof mytoken !== "string") {
            res.status(400).json({ msg: "invalid token format" });
            return;
        }
        const user = await users_1.userModel.findOne({
            "userVerifyToken.emailVerifyToken": mytoken,
        });
        if (!user) {
            res.status(400).json({ msg: "invalid or expired token" });
            return;
        }
        // ✅ CHECK FIRST
        if (user.userVerified.emailVerified) {
            res.status(200).json({ msg: "email already verified." });
            return;
        }
        // ✅ THEN UPDATE
        user.userVerified.emailVerified = true;
        user.userVerifyToken.emailVerifyToken = null;
        await user.save();
        res.status(200).json({ msg: "Email verified successfully! ✨" });
    }
    catch (error) {
        console.log(error);
        res.status(500).json({ msg: "something went wrong" });
    }
});
// // phone verify api 
// router.get("/phoneverify/:token", async (req: Request, res: Response): Promise<void>=>{
//     try {
//         const token =  req.params.token
//         if (typeof token != "string"){
//             res.status(400).json({msg: "invalid token format"})
//             return
//         }
//         let user = await userModel.findOne({"userVerifyToken.phoneVerifyToken": token})
//         if(!user){
//             res.status(400).json({msg: "invalid token"})
//             return
//         }
//         user.userVerified.phoneVerified = true;
//         user.userVerifyToken.phoneVerifyToken = null
//         if(user.userVerified.phoneVerified){
//             res.status(200).json({msg: "phone verified"})
//             return
//         }
//         res.status(200).json({msg: "phone verified✅"})
//     } catch (error) {
//         console.log(error);
//         res.status(500).json({msg: error})
//     }
// })
// signin api
router.post("/signin", async (req, res) => {
    try {
        const { email, password } = req.body;
        if (!email || !password) {
            res.status(400).json({ msg: "pls fill all fields" });
            return;
        }
        let checkUser = await users_1.userModel.findOne({ email });
        if (!checkUser) {
            res.status(400).json({ msg: "no user found" });
            return;
        }
        // compare password 
        let pass = await bcrypt_1.default.compare(password, checkUser.password);
        if (!pass) {
            res.status(400).json({ msg: "invalid password" });
            return;
        }
        // jwt gtoken gen
        let token = jsonwebtoken_1.default.sign({ id: checkUser.id }, KEY, { expiresIn: "40d" });
        res.status(200).json({ msg: "Logged in✅", token });
    }
    catch (error) {
        res.status(500).json({ msg: error });
        console.log(error);
    }
});
// forgot pass api 
router.post("/forgotpass", async (req, res) => {
    try {
        const { email } = req.body;
        if (!email) {
            res.status(400).json({ msg: "no email given" });
            return;
        }
        let checkUser = await users_1.userModel.findOne({ email });
        if (!checkUser) {
            res.status(400).json({ msg: "no user found" });
            return;
        }
        // gen new pass 
        let newPass = Math.random().toString(36).substring(2);
        console.log(newPass);
        //   send pass on email
        const emailData = {
            from: USER,
            subject: "New Password",
            to: email,
            html: `<p>Your new password is: <strong>${newPass}</strong></p>`
        };
        (0, sendEmail_1.default)(emailData);
        let hashPass = await bcrypt_1.default.hash(newPass, 10);
        checkUser.password = hashPass;
        await checkUser.save();
        res.status(200).json({ msg: "temp password sent to email successfully!✅" });
    }
    catch (error) {
        res.status(500).json({ msg: error });
        console.log(error);
    }
});
exports.default = router;
