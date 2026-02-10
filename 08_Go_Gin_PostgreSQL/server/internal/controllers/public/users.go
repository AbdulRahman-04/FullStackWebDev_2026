// package public

// import (
// 	"crypto/rand"
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"strings"
// 	"time"

// 	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
// 	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/models"
// 	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"
// )

// // token generate function
// func TokenGenerate(length int) string {
// 	b := make([]byte, length)
// 	_, _ = rand.Read(b)
// 	return hex.EncodeToString(b)
// }

// // otp generate function
// func OtpGenerate() string {
// 	n, _ := rand.Int(rand.Reader, big.NewInt(10000))
// 	return fmt.Sprintf("%04d", n.Int64())
// }

// // signup api
// func UserSignup(c *gin.Context) {
// 	// create struct to store/map userinput values inside variable
// 	type UserSignUp struct {
// 		Email    string `json:"email"`
// 		FullName string `json:"fullName"`
// 		Password string `json:"password"`
// 		Phone    string `json:"phone"`
// 	}

// 	var userInput UserSignUp
// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "ivalid request",
// 		})
// 		return
// 	}

// 	// trim spaces
// 	userInput.Email = strings.TrimSpace(userInput.Email)
// 	userInput.Phone = strings.TrimSpace(userInput.Phone)
// 	userInput.FullName = strings.TrimSpace(userInput.FullName)

// 	// validations
// 	if userInput.Email == "" || userInput.FullName == "" || userInput.Password == "" || userInput.Phone == "" {
// 		c.JSON(400, gin.H{
// 			"msg": "fill all fields",
// 		})
// 		return
// 	}

// 	if !strings.Contains(userInput.Email, "@") || len(userInput.Password) < 6 {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid email or pass length less than 6",
// 		})
// 		return
// 	}

// 	if strings.HasPrefix(userInput.Password, " ") || strings.HasSuffix(userInput.Password, " ") {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid pass",
// 		})
// 		return
// 	}

// 	// duplicate count
// 	var count int64
// 	utils.PostgresDB.Model(&models.User{}).Where("email = ? OR phone = ?", userInput.Email, userInput.Phone).Count(&count)
// 	if count > 0 {
// 		c.JSON(400, gin.H{
// 			"msg": "user already exists",
// 		})
// 		return
// 	}

// 	// sql query
// 	// SELECT COUNT(*) FROM users WHERE email = userInput.email OR phone = userInput.phone

// 	// hash pass
// 	hashPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "password hash err",
// 		})
// 		return
// 	}

// 	emailToken := TokenGenerate(32)
// 	otp := OtpGenerate()

// 	// create var to insert values in  db
// 	var user models.User

// 	user.FullName = userInput.FullName
// 	user.Email = userInput.Email
// 	user.Password = string(hashPass)
// 	user.Phone = userInput.Phone
// 	user.EmailVerifyToken = emailToken
// 	user.Provider = "email"

// 	// insert into db
// CreateUser := utils.PostgresDB.Create(&user)
// if CreateUser.Error != nil {
//       c.JSON(400, gin.H{
// 		"msg": "no user found",
// 	})
// 	return
// }

// 	// sql query
// 	// INSERT INTO users (
// 	// 	full_name,
// 	// 	email,
// 	// 	password,
// 	// 	phone,
// 	// 	provider,
// 	// 	role,
// 	// 	email_verified,
// 	// 	phone_verified,
// 	// 	created_at,
// 	// 	updated_at
// 	// )
// 	// VALUES (
// 	// 	'Dev Rahman',
// 	// 	'devrxhmann@gmail.com',
// 	// 	'$2a$10$hash...',
// 	// 	'+918186978069',
// 	// 	'email',
// 	// 	'user',
// 	// 	false,
// 	// 	false,
// 	// 	NOW(),
// 	// 	NOW()
// 	// );

// 	// redis set otp
// 	_ = utils.RedisSetKey("otp_"+userInput.Phone, otp, 2*time.Minute)

// 	// send email
// 	utils.EmailQueue <- utils.EmailData{
// 		To:      userInput.Email,
// 		Subject: "email verify",
// 		Text:    "verify email",
// 		Html: fmt.Sprintf(
// 			`<p>Click below to verify:</p>
//          <a href="%s/api/public/user/emailverify/%s">Verify Email</a>`, config.AppConfig.URL, emailToken),
// 	}

// 	// send sms
// 	utils.SmsQueue <- utils.SMSData{
// 		To:   userInput.Phone,
// 		Body: "otp: " + otp,
// 	}

// 	fmt.Printf("OTP: %s", otp)

// 	// response
// 	c.JSON(200, gin.H{
// 		"msg": "user signed up successfully!✨",
// 	})

// }

// // email verify api
// func EmailVerify(c *gin.Context) {
// 	token := c.Param("token")

// 	// compare nd set db
// 	res := utils.PostgresDB.Model(&models.User{}).Where("email_verify_token = ? AND email_verified = false", token).Updates(map[string]any{
// 		"email_verify_token": "",
// 		"email_verified":     true,
// 	})

// sql
// Update users SET email_verify_token = '',email_verified = true WHERE email_verify_token ="param token" AND email_verified = false

// if res.Error != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "db err",
// 		})
// 		return
// 	}

// 	if res.RowsAffected == 0 {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid or expired token",
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"msg": "email verified✅",
// 	})
// }

// // phone verify
// func PhoneVerify(c *gin.Context) {
// 	// struct
// 	type PhoneVerify struct {
// 		Phone string `json:"phone"`
// 		Otp   string `json:"otp"`
// 	}

// 	var userInput PhoneVerify
// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid request",
// 		})
// 		return
// 	}

// 	if userInput.Phone == "" || len(userInput.Otp) < 4 {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid phone or length of otp is invalid",
// 		})
// 		return
// 	}

// 	// get redis key
// 	otp, err := utils.RedisGetKey("otp_" + userInput.Phone)
// 	if err != nil || otp != userInput.Otp {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid or expired otp",
// 		})
// 		return
// 	}

// 	// update db
// 	res := utils.PostgresDB.Model(&models.User{}).Where("phone = ?", userInput.Phone).Updates(map[string]any{
// 		"phone_verified": true,
// 	})

// 	if res.Error != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "db error",
// 		})
// 		return
// 	}

// 	if res.RowsAffected == 0 {
// 		c.JSON(400, gin.H{
// 			"msg": "no phone no found in db",
// 		})
// 		return
// 	}

// 	// sql query
// 	// UPDATE users SET phone_verified = true WHERE phone = "user input wala phone"

// 	// redis del key
// 	_ = utils.RedisDelKey("otp_" + userInput.Phone)

// 	c.JSON(200, gin.H{
// 		"msg": "Phone Verified✅",
// 	})
// }

// // signin api
// func UserSignin(c *gin.Context) {
// 	// make struct
// 	type UserSignIn struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	var userInput UserSignIn
// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid request",
// 		})
// 		return
// 	}

// 	// validations
// 	userInput.Email = strings.TrimSpace(userInput.Email)

// 	if userInput.Email == "" || userInput.Password == "" {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid request, pls fill all fields",
// 		})
// 		return
// 	}

// 	if strings.HasPrefix(userInput.Password, " ") || strings.HasSuffix(userInput.Password, " ") {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid password, must not have spaces before nd after",
// 		})
// 		return
// 	}

// 	// check db if user exists or not
// 	var checkUser models.User
// 	if err := utils.PostgresDB.Model(&models.User{}).Where("email = ?", userInput.Email).First(&checkUser).Error; err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "user doesn't exist",
// 		})
// 		return
// 	}

// 	// check pass
// 	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(userInput.Password)); err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid password",
// 		})
// 		return
// 	}

// 	// check if email nd phone verified
// 	if checkUser.EmailVerified != true || checkUser.PhoneVerified != true {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid request, pls verify email or phone first",
// 		})
// 		return
// 	}

// 	// token generate
// 	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":   checkUser.ID,
// 		"role": checkUser.Role,
// 		"exp":  time.Now().Add(24 * time.Hour).Unix(),
// 	}).SignedString([]byte(config.AppConfig.JWT_KEY))

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "token generate error",
// 		})
// 		return
// 	}

// 	// generate refresh token
// 	refreshToken := TokenGenerate(32)

// 	// db update
// 	res := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkUser.ID).Updates(map[string]any{
// 		"refresh_token":  refreshToken,
// 		"refresh_expiry": time.Now().Add(24 * 7 * time.Hour),
// 	})

// 	if res.Error != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "db update failed",
// 		})
// 		return
// 	}

// 	if res.RowsAffected == 0 {
// 		c.JSON(400, gin.H{
// 			"msg": "no user id found to update",
// 		})
// 		return
// 	}

// 	// sql query
// 	// Update users SET refresh_token = "generatedtoken",refresh_expiry=7 days timestamp WHERE id = checkUser.id

// 	c.JSON(200, gin.H{
// 		"msg":          "user logged in✅",
// 		"AccessToken":  token,
// 		"refreshToken": refreshToken,
// 	})

// }

// // refresh api
// func RefreshToken(c *gin.Context) {
// 	// make struct
// 	type RefreshToken struct {
// 		RefreshToken string `json:"refreshToken"`
// 	}

// 	var userInput RefreshToken
// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "ivalid request",
// 		})
// 		return
// 	}

// 	// check if this refresh token exist ind bor not
// 	var checkToken models.User
// 	if err := utils.PostgresDB.Model(&models.User{}).Where("refresh_token = ?", userInput.RefreshToken).First(&checkToken).Error; err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid refresh token",
// 		})
// 		return
// 	}

// 	// check if current token is expires
// 	if time.Now().After(checkToken.RefreshExpiry) {
// 		c.JSON(400, gin.H{
// 			"msg": "refersh token expired",
// 		})
// 		return
// 	}

// 	// create new access token
// 	newToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":   checkToken.ID,
// 		"role": checkToken.Role,
// 		"exp":  time.Now().Add(24 * time.Hour).Unix(),
// 	}).SignedString([]byte(config.AppConfig.JWT_KEY))

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "ivalid token err",
// 		})
// 		return
// 	}

// 	// new refresh token gen
// 	newRefreshToken := TokenGenerate(32)

// 	// update db refresh token
// 	res := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkToken.ID).Updates(map[string]any{
// 		"refresh_token":  newRefreshToken,
// 		"refresh_expiry": time.Now().Add(7 * 24 * time.Hour),
// 	})

// 	if res.Error != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "ivalid db err",
// 		})
// 		return
// 	}

// 	if res.RowsAffected == 0 {
// 		c.JSON(400, gin.H{
// 			"msg": "ivalid request, no id found in db",
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"accessToken":  newToken,
// 		"refreshToken": newRefreshToken,
// 	})
// }

// // forgot password api
// func ForgotPassword(c *gin.Context) {
// 	// make struct
// 	type ForgotPassword struct {
// 		Email string `json:"email"`
// 	}

// 	var userInput ForgotPassword
// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid request",
// 		})
// 		return
// 	}

// 	userInput.Email = strings.TrimSpace(userInput.Email)

// 	// check if user exists in db
// 	var checkUser models.User
// 	if err := utils.PostgresDB.Model(&models.User{}).Where("email = ?", userInput.Email).First(&checkUser).Error; err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "no email found",
// 		})
// 		return
// 	}

// 	// generate new password
// 	tempPass := TokenGenerate(12)

// 	// hash pass
// 	hashPass, err := bcrypt.GenerateFromPassword([]byte(tempPass), 10)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid bcrypt error",
// 		})
// 		return
// 	}

// 	// db update password
// 	res := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkUser.ID).Updates(map[string]any{
// 		"password": string(hashPass),
// 	})

// 	if res.Error != nil {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid db error",
// 		})
// 		return
// 	}

// 	if res.RowsAffected == 0 {
// 		c.JSON(400, gin.H{
// 			"msg": "invalid, no user found",
// 		})
// 		return
// 	}

// 	// send temp pass on mail
// 	utils.EmailQueue <- utils.EmailData{
// 		To:      userInput.Email,
// 		Subject: "Team Social",
// 		Html: fmt.Sprintf(
// 			`<h3>Hello %s, your temporary password is <b>%s</b></h3>`, checkUser.FullName, tempPass),
// 	}

// 	c.JSON(200, gin.H{
// 		"msg": "Temporary password sent to ur email✅",
// 	})

// }

package public

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/models"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// token generate
func TokenGenerate(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// otp generate
func OTPGenerate() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(10000))
	return fmt.Sprintf("%04d", n.Int64())
}

// signup api
func UserSignup(c *gin.Context) {
	// make struct
	type UserSignup struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
	}

	var userInput UserSignup
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// trim
	userInput.FullName = strings.TrimSpace(userInput.FullName)
	userInput.Email = strings.TrimSpace(userInput.Email)
	userInput.Phone = strings.TrimSpace(userInput.Phone)

	// validations
	if userInput.Email == "" || userInput.FullName == "" || userInput.Password == "" || userInput.Phone == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request, fill all fields",
		})
		return
	}

	if !strings.Contains(userInput.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if strings.HasPrefix(userInput.Password, " ") || strings.HasSuffix(userInput.Password, " ") || len(userInput.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// duplicate checck in db
	var count int64
	utils.PostgresDB.Model(&models.User{}).Where("email = ? OR phone = ? ", userInput.Email, userInput.Phone).Count(&count)
	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "user already exists",
		})
		return
	}

	// sql
	// SELECT COUNT(*) FROM users WHERE email = "userinp email" OR phone = "userinp phone"

	// hash pass
	hashPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid bcrypt err",
		})
		return
	}
	emailToken := TokenGenerate(32)
	otp := OTPGenerate()

	// create variable to insert values in db
	var user models.User

	user.FullName = userInput.FullName
	user.Email = userInput.Email
	user.Password = string(hashPass)
	user.Phone = userInput.Phone
	user.EmailVerifyToken = emailToken
	user.Provider = "email"

	createUser := utils.PostgresDB.Model(&models.User{}).Create(&user)

	if createUser.Error != nil {
		c.JSON(400, gin.H{
			"msg": "no user found",
		})
		return
	}

	// sql
	// INSERT into users (fields) VALUES ()

	// set redis key for otp
	_ = utils.RedisSetKey("otp_"+userInput.Phone, otp, 2*time.Minute)

	// send email
	utils.EmailQueue <- utils.EmailData{
		To:      userInput.Email,
		Subject: "Team Social",
		Text:    "Verify email",
		Html: fmt.Sprintf(
			`<p>Click below to verify:</p>
         <a href="%s/api/public/user/emailverify/%s">Verify Email</a>`, config.AppConfig.URL, emailToken),
	}

	// sms worker
	utils.SmsQueue <- utils.SMSData{
		To:   userInput.Phone,
		Body: "OTP: " + otp,
	}

	log.Printf("Otp %s", otp)

	c.JSON(200, gin.H{
		"msg": "User signed up successfully🎉, verify ur email nd phone and then login🚀",
	})
}

// email verify api
func EmailVerify(c *gin.Context) {
	// take token from param
	token := c.Param("token")

	// find email token nd update db
	res := utils.PostgresDB.Model(&models.User{}).Where("email_verify_token = ? AND email_verified = false", token).Updates(map[string]any{
		"email_verify_token": "",
		"email_verified":     true,
	})

	// sql
	// Update users SET email_verify_token = '',email_verified = true WHERE email_verify_token ="param token" AND email_verified = false

	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err",
		})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid or expired email tokn",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "Email Verified✨✅",
	})
}

// phone verify api
func PhoneVerify(c *gin.Context) {
	// make struct
	type PhoneVerify struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}

	var userInput PhoneVerify
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validation
	if userInput.Phone == "" || len(userInput.Otp) < 4 {
		c.JSON(400, gin.H{
			"msg": "invalid phone or length of otp is invalid",
		})
		return
	}

	// verify if phone no exists in db
	var checkUser models.User
	res := utils.PostgresDB.Model(&models.User{}).Where("phone = ?", userInput.Phone).First(&checkUser)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid phone number",
		})
		return
	}

	// get otp from redis
	otp, err := utils.RedisGetKey("otp_" + userInput.Phone)
	if err != nil || otp != userInput.Otp {
		c.JSON(400, gin.H{
			"msg": "invalid or expired otp",
		})
		return
	}

	// update db for phone_verified
	res = utils.PostgresDB.Model(&models.User{}).Where("id = ? ", checkUser.ID).Updates(map[string]any{
		"phone_verified": true,
	})

	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid user id, not found ",
		})
		return
	}

	// del otp from redis
	_ = utils.RedisDelKey("otp_" + userInput.Phone)

	c.JSON(200, gin.H{
		"msg": "User Phone number verified✨🚀",
	})
}

// sign in api
func UserSignin(c *gin.Context) {
	// type struct
	type UserSignIn struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userInput UserSignIn
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// 	// validations
	userInput.Email = strings.TrimSpace(userInput.Email)

	if userInput.Email == "" || userInput.Password == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request, pls fill all fields",
		})
		return
	}

	if strings.HasPrefix(userInput.Password, " ") || strings.HasSuffix(userInput.Password, " ") {
		c.JSON(400, gin.H{
			"msg": "invalid password, must not have spaces before nd after",
		})
		return
	}

	// check if email exists in db
	var checkUser models.User
	ress := utils.PostgresDB.Model(&models.User{}).Where("email = ?", userInput.Email).First(&checkUser)
	if ress.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request, email doesn't exist",
		})
		return
	}

	// password check
	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(userInput.Password)); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// check if email nd phone verified
	if checkUser.EmailVerified != true || checkUser.PhoneVerified != true {
		c.JSON(400, gin.H{
			"msg": "invalid request, pls verify email or phone first",
		})
		return
	}

	// access token generate
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   checkUser.ID,
		"role": checkUser.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token generation err",
		})
		return
	}

	// generate refresh token
	refreshToken := TokenGenerate(32)

	// update db
	res := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkUser.ID).Updates(map[string]any{
		"refresh_token":  refreshToken,
		"refresh_expiry": time.Now().Add(7 * 24 * time.Hour),
	})

	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid id, user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":          "Logged in successful!✅",
		"accessToken":  token,
		"refreshToken": refreshToken,
	})

}

// refresh api

func RefreshToken(c *gin.Context) {
	type RefreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}

	var userInput RefreshToken
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// check if refresh token exists and its validity
	var checkToken models.User
	res := utils.PostgresDB.Model(&models.User{}).Where("refresh_token = ?", userInput.RefreshToken).First(&checkToken)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid refresh token",
		})
		return
	}

	if time.Now().After(checkToken.RefreshExpiry) {
		c.JSON(400, gin.H{
			"msg": "expired refresh token",
		})
		return
	}

	// gen new token
	newToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   checkToken.ID,
		"role": checkToken.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token generation err",
		})
		return
	}

	// new refres token
	newRefreshToken := TokenGenerate(32)

	// db update
	result := utils.PostgresDB.Model(&models.User{}).Where("id = ? ", checkToken.ID).Updates(map[string]any{
		"refresh_token":  newRefreshToken,
		"refresh_expiry": time.Now().Add(7 * 24 * time.Hour),
	})

	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid user id",
		})
		return
	}

	c.JSON(200, gin.H{
		"accesstoken": newToken,
		"refrehToken": newRefreshToken,
	})
}

// forgot password api
func ForgotPassword(c *gin.Context) {
	type ForgotPassword struct {
		Email string `json:"email"`
	}

	var userInput ForgotPassword
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// check if user exists or not
	var checkUser models.User
	res := utils.PostgresDB.Model(&models.User{}).Where("email = ?", userInput.Email).First(&checkUser)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	// generate random pass
	tempass := TokenGenerate(12)

	// hash pass
	hashPass, err := bcrypt.GenerateFromPassword([]byte(tempass), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid bcrypt err",
		})
		return
	}

	// update db
	result := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkUser.ID).Updates(map[string]any{
		"password": string(hashPass),
	})

	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid user id",
		})
		return
	}

	// send pass to email
	utils.EmailQueue <- utils.EmailData{
		To: userInput.Email,
				Subject: "Team Social",
				Html: fmt.Sprintf(
					`<h3>Hello %s, your temporary password is <b>%s</b></h3>`, checkUser.FullName, tempass),
	}

	c.JSON(200, gin.H{
		"msg": "Temp password sent to email✅",
	})

}
