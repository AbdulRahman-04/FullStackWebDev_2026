package public

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/models"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/utils"
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

// otp  generate
func OTPGenerate() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(10000))
	return fmt.Sprintf("%04d", n.Int64())
}

// user signup controller
func UserSignup(c *gin.Context) {
	// struct
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

	// trim spaces
	userInput.Email = strings.TrimSpace(userInput.Email)
	userInput.FullName = strings.TrimSpace(userInput.FullName)
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

	if len(userInput.Password) < 6 || strings.HasPrefix(userInput.Password, " ") || strings.HasSuffix(userInput.Password, " ") {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// duplicate check in db
	var Count int64
	checkUser := utils.PostgresDB.Model(&models.User{}).Where("email = ? OR phone = ?", userInput.Email, userInput.Phone).Count(&Count)
	if checkUser.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if Count > 0 {
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

	// create var nd push into db
	var user models.User

	user.FullName = userInput.FullName
	user.Email = userInput.Email
	user.Password = string(hashPass)
	user.Phone = userInput.Phone
	user.Provider = "email"
	user.EmailVerifyToken = emailToken

	createUser := utils.PostgresDB.Model(&models.User{}).Create(&user)
	if createUser.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}
	// set otp in redis
	err = utils.RedisSetKey("otp_"+userInput.Phone, otp, 2*time.Minute)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid redis err",
		})
		return
	}

	// send email and otp
	utils.EmailQueue <- utils.EmailData{
		To:      userInput.Email,
		Subject: "email verify",
		Text:    "Verify email",
		Html: fmt.Sprintf(
			`<p>Click below to verify:</p>
         <a href="%s/api/public/user/emailverify/%s">Verify Email</a>`, config.AppConfig.URL, emailToken),
	}

	utils.SMSQueue <- utils.SMSData{
		To:   userInput.Phone,
		Body: "otp: " + otp,
	}

	c.JSON(200, gin.H{
		"msg": "User signed up successfully! verify ur email and phone nuber and signin🎉",
	})
}

// email verify api
func EmailVerify(c *gin.Context) {
	token := c.Param("token")

	checkDb := utils.PostgresDB.Model(&models.User{}).Where("email_verify_token = ? AND email_verified = false", token).Updates(map[string]any{
		"email_verify_token": "",
		"email_verified":     true,
	})

	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if checkDb.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid email token",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "email verified✨",
	})
}

// phone verify
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

	// trim spaces
	userInput.Phone = strings.TrimSpace(userInput.Phone)
	userInput.OTP = strings.TrimSpace(userInput.OTP)

	// validations
	if userInput.Phone == "" || len(userInput.OTP) < 4 {
		c.JSON(400, gin.H{
			"msg": "invalid phone or length of otp is invalid",
		})
		return
	}

	// check if phoen exists in db
	var checkPhone models.User
	dbCheck := utils.PostgresDB.Model(&models.User{}).Where("phone = ?", userInput.Phone).First(&checkPhone)
	if dbCheck.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid phone number",
		})
		return
	}

	// get otp from redis
	otp, err := utils.RedisGetKey("otp_" + userInput.Phone)
	if err != nil || otp != userInput.OTP {
		c.JSON(400, gin.H{
			"msg": "invalid or expired otp",
		})
		return
	}

	// update db
	updateDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkPhone.ID).Updates(map[string]any{
		"phone_verified": true,
	})

	if updateDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if updateDb.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid useridnot found",
		})
		return
	}

	// redis del key
	_ = utils.RedisDelKey("otp_" + userInput.Phone)

	c.JSON(200, gin.H{
		"msg": "Phone verified",
	})
}

// user signin api
func UserSignin(c *gin.Context) {
	type UserSignin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userInput UserSignin
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// trim spaces and validations
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

	// check if email exists in db
	var checkEmail models.User
	dbCheck := utils.PostgresDB.Model(&models.User{}).Where("email = ?", userInput.Email).First(&checkEmail)
	if dbCheck.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request,no email found",
		})
		return
	}

	// pass check
	if err := bcrypt.CompareHashAndPassword([]byte(checkEmail.Password), []byte(userInput.Password)); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// ye check bhool gaya! ⚠️
	if checkEmail.EmailVerified != true || checkEmail.PhoneVerified != true {
		c.JSON(400, gin.H{
			"msg": "pls verify ur email and phone then login",
		})
		return
	}

	// token gen
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   checkEmail.ID,
		"role": checkEmail.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "token generation error",
		})
		return
	}

	// refresh token gen
	refreshToken := TokenGenerate(32)

	updateDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?", checkEmail.ID).Updates(map[string]any{
		"refresh_token":  refreshToken,
		"refresh_expiry": time.Now().Add(7 * 24 * time.Hour),
	})

	if updateDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err",
		})
		return
	}

	if updateDb.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "no userid found",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":          "Logged In✅",
		"accessToken":  token,
		"refreshToken": refreshToken,
	})

}

// refresh token api 
func RefreshToken(c*gin.Context) {
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

	// check rt exists in db 
	var checkToken models.User
	checkDb := utils.PostgresDB.Model(&models.User{}).Where("refresh_token = ?", userInput.RefreshToken).First(&checkToken)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token not found",
		})
		return
	}

	// check token expiry 
	if time.Now().After(checkToken.RefreshExpiry) {
		c.JSON(400, gin.H{
			"msg": "invalid token expired",
		})
		return
	}

	// gen new access token 
	newToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": checkToken.ID,
		"role": checkToken.Role,
		"exp":  time.Now().Add(24*time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token generation err",
		})
		return
	}

	refreshToken := TokenGenerate(32)

	updateDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?",checkToken.ID).Updates(map[string]any{
		"refresh_token":refreshToken,
		"refresh_expiry":time.Now().Add(7*24*time.Hour),
	})

	if updateDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if updateDb.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid id, user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"accessToken":  newToken,
		"refreshToken": refreshToken,
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
		To:      userInput.Email,
		Subject: "Team Social",
		Html: fmt.Sprintf(
			`<h3>Hello %s, your temporary password is <b>%s</b></h3>`, checkUser.FullName, tempPass),
	}

	c.JSON(200, gin.H{
		"msg": "Temp password sent on ur email✅",
	})
}
