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

// TokenGenerate
func TokenGenerate(lenght int) string {
	b := make([]byte, lenght)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// otp generate func
func OTPGenerate() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(10000))

	return fmt.Sprintf("%04d", n.Int64())
}

// user signup function
func UserSignup(c *gin.Context) {
	// make struct
	type UserSignup struct {
		FullName string `json:"fullName"`
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

	// trim spaces
	userInput.Email = strings.TrimSpace(userInput.Email)
	userInput.FullName = strings.TrimSpace(userInput.FullName)
	userInput.Phone = strings.TrimSpace(userInput.Phone)

	// validations
	if userInput.FullName == "" || userInput.Email == "" || userInput.Password == "" || userInput.Phone == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request, fill all fields",
		})
		return
	}

	if !strings.Contains(userInput.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	if strings.HasPrefix(userInput.Password, " ") || strings.HasSuffix(userInput.Password, " ") || len(userInput.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// duplicate count
	var count int64
	utils.PostgresDB.Model(&models.User{}).Where("email = ? OR phone = ?", userInput.Email, userInput.Phone).Count(&count)
	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "invalid request, user already exists",
		})
		return
	}

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

	// create var to insert user into db
	var user models.User

	user.FullName = userInput.FullName
	user.Email = userInput.Email
	user.Password = string(hashPass)
	user.Phone = userInput.Phone
	user.EmailVerifyToken = emailToken
	user.Provider = "email"

	// create user in db
	res := utils.PostgresDB.Create(&user)

	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	// set otp in redis
	_ = utils.RedisSetKey("otp_"+userInput.Phone, otp, 2*time.Minute)

	// send email nd otp
	utils.EmailQueue <- utils.EmailData{
		To:      userInput.Email,
		Subject: "email verify",
		Text:    "Verify email",
		Html: fmt.Sprintf(
			`<p>Click below to verify:</p>
         <a href="%s/api/public/user/emailverify/%s">Verify Email</a>`, config.AppConfig.URL, emailToken),
	}

	// sms
	utils.SmsQueue <- utils.SMSData{
		To:   userInput.Phone,
		Body: "otp :" + otp,
	}

	log.Printf("Otp %s", otp)

	c.JSON(200, gin.H{
		"msg": "User signed up successfully!🎉, verify ur email nd pass nd login",
	})

	fmt.Println("Generated User ID:", user.ID)

}

// email verify api
func EmailVerify(c *gin.Context) {
	token := c.Param("token")

	// check token in db nd update db
	res := utils.PostgresDB.Model(&models.User{}).Where("email_verify_token = ? AND email_verified = false", token).Updates(map[string]any{
		"email_verify_token": "",
		"email_verified":     true,
	})

	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid or expired token",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "emailVerified✨🚀",
	})
}

// phone verify token
func PhoneVerify(c *gin.Context) {
	type PhoneVerify struct {
		Phone string `json:"phone"`
		OTP   string `json:"otp"`
	}

	var userInput PhoneVerify
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validation
	if userInput.Phone == "" || len(userInput.OTP) < 4 {
		c.JSON(400, gin.H{
			"msg": "invalid phone or length of otp is invalid",
		})
		return
	}

	// check if phone exists in db
	var checkPhone models.User
	res := utils.PostgresDB.Model(&models.User{}).Where("phone = ?", userInput.Phone).First(&checkPhone)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid phone number",
		})
		return
	}

	// get otp from redis
	otp, err := utils.RedisGetKey("otp_" + userInput.Phone)
	if err != nil || otp != userInput.OTP {
		c.JSON(400, gin.H{
			"msg": "invalid otp",
		})
		return
	}

	// updaate db
	result := utils.PostgresDB.Model(&models.User{}).Where("id = ? ", checkPhone.ID).Updates(map[string]any{
		"phone_verified": true,
	})

	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid id, user phone no  not found",
		})
		return
	}

	// del key from redis
	_ = utils.RedisDelKey("otp_" + userInput.Phone)

	c.JSON(200, gin.H{
		"msg": "Phone number verified✨✅",
	})
}

// signin api
func UserSignin(c *gin.Context) {
	// struct
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

	// validations
	userInput.Email = strings.TrimSpace(userInput.Email)

	if userInput.Email == "" || userInput.Password == "" {
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

	// check db if email exists or not
	var checkUser models.User
	res := utils.PostgresDB.Model(&models.User{}).Where("email = ?", userInput.Email).First(&checkUser)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "no email found⚠️",
		})
		return
	}

	// check pass
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

	// token generate
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   checkUser.ID,
		"role": checkUser.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token gen err",
		})
		return
	}

	refreshToken := TokenGenerate(32)

	// update db for rt nd re
	result := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkUser.ID).Updates(map[string]any{
		"refresh_token":  refreshToken,
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
			"msg": "user id not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"Msg":          "User logged in✅",
		"accessToken":  token,
		"refreshToken": refreshToken,
	})
}

// refresh token api
func RefreshToken(c *gin.Context) {
	type RefreshToken struct {
		RefreshToken string `json:"refreshToken"`
	}

	var userInput RefreshToken
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}
	// check rt exists in db
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

	// new token generate
	newToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   checkToken.ID,
		"role": checkToken.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token gen err",
		})
		return
	}

	// new refresh token gen
	newRefreshToken := TokenGenerate(32)

	// update db
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
			"msg": "invalid id, user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"accessToken":  newToken,
		"refreshToken": newRefreshToken,
	})
}

// forgot password api
func ForgotPassword(c*gin.Context){
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

	// email exists indb
	var checkUser models.User
	res := utils.PostgresDB.Model(&models.User{}).Where("email = ?", userInput.Email).First(&checkUser)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request, email doesn;t exist",
		})
		return
	}

	// create temp password
	tempPass := TokenGenerate(20)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(tempPass), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid bcrypt err",
		})
		return
	}

	// db update 
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
			"msg": "invalid id, user not found",
		})
		return
	}

	// send pass on email 
	utils.EmailQueue <- utils.EmailData{
		To: userInput.Email,
				Subject: "Team Social",
				Html: fmt.Sprintf(
					`<h3>Hello %s, your temporary password is <b>%s</b></h3>`,checkUser.FullName, tempPass),
	}

	c.JSON(200, gin.H{
		"msg": "Temp password sent on ur email✅",
	})
}