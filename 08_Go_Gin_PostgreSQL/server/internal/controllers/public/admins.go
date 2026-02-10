package public

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/models"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// admin signup api
func AdminSignUp(c*gin.Context){
	type AdminSignUp struct {
		FullName string `json:"fullName"`
		Email string `json:"email"`
		Password string `json:"password"`
		Phone string `json:"phone"`
	}

	var adminInput AdminSignUp
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations 
	adminInput.FullName = strings.TrimSpace(adminInput.FullName)
	adminInput.Email = strings.TrimSpace(adminInput.Email)
	adminInput.Phone = strings.TrimSpace(adminInput.Phone)

	if adminInput.FullName == "" || adminInput.Email == "" || adminInput.Password == "" || adminInput.Phone == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request, fill all fields",
		})
		return
	}

	if !strings.Contains(adminInput.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if strings.HasPrefix(adminInput.Password, " ") || strings.HasSuffix(adminInput.Password, " ") || len(adminInput.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// duplicate count
	var count int64
	utils.PostgresDB.Model(&models.Admin{}).Where("email = ? OR phone = ?", adminInput.Email, adminInput.Phone).Count(&count)
	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "user already exist",
		})
		return
	}

	// hash pass 
	hashPass, err := bcrypt.GenerateFromPassword([]byte(adminInput.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid bcrypt err",
		})
		return
	}

	emailToken := TokenGenerate(32)
	otp := OTPGenerate()


	// create admin var nd insert into db
	var Admin models.Admin

	Admin.FullName = adminInput.FullName
	Admin.Email = adminInput.Email
	Admin.Password = string(hashPass)
	Admin.Phone = adminInput.Phone
	Admin.EmailVerifyToken = emailToken
	Admin.Provider = "email"

	// insert into db 
	insert := utils.PostgresDB.Model(&models.Admin{}).Create(&Admin)
	if insert.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	// redis set otp 
	_ = utils.RedisSetKey("otp_"+adminInput.Phone, otp, 2*time.Minute)

	utils.EmailQueue <- utils.EmailData{
		To:      adminInput.Email,
		Subject: "email verify",
		Text:    "Verify email",
		Html: fmt.Sprintf(
			`<p>Click below to verify:</p>
         <a href="%s/api/public/admin/emailverify/%s">Verify Email</a>`,config.AppConfig.URL, emailToken),
	}

	utils.SmsQueue <- utils.SMSData{
		To: adminInput.Phone,
		Body: "Otp: "+otp,
	}

	log.Printf("OTP: %s", otp)

	c.JSON(200, gin.H{
		"msg": "Admin signed up successfully!✅",
	})
}

// email verify api 
func AdminEmailVerify(c*gin.Context){
	token := c.Param("token")

	// db update 
	dbUpdate := utils.PostgresDB.Model(&models.Admin{}).Where("email_verify_token = ? AND email_verified = false", token).Updates(map[string]any{
		"email_verify_token": "",
		"email_verified":true,
	})

	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid or expired token",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "email verified✨🚀",
	})
}

// phone verify api 
func AdminPhoneVerify(c*gin.Context){
	type AdminPhoneVerify struct {
		Phone string `json:"phone"`
		OTP string `json:"otp"`
	}

	var adminInput AdminPhoneVerify
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validation
	if adminInput.Phone == "" || len(adminInput.OTP) != 4 {
		c.JSON(400, gin.H{
			"msg": "invalid phone number",
		})
		return
	}

	// check if phone exist or not 
	var checkPhone models.Admin
	admin := utils.PostgresDB.Model(&models.Admin{}).Where("phone = ?",adminInput.Phone).First(&checkPhone)
	if admin.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid mobile no",
		})
		return
	}

	// get otp from redis 
	otp, err := utils.RedisGetKey("otp_"+adminInput.Phone)
	if err != nil || otp != adminInput.OTP {
		c.JSON(400, gin.H{
			"msg": "invalid otp",
		})
		return
	}

	// db update 
	dbUpdate := utils.PostgresDB.Model(&models.Admin{}).Where("id = ?", checkPhone.ID).Updates(map[string]any{
		"phone_verified":true,
	})

	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid id, user not found",
		})
		return
	}

	// del redist key 
	_ = utils.RedisDelKey("otp_"+adminInput.Phone)

	c.JSON(200, gin.H{
		"msg": "Phone Number verified✅",
	})
}

// admin sign in 
func AdminSignIn(c*gin.Context){
	type AdminSignIn struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var adminInput AdminSignIn
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations 
	adminInput.Email = strings.TrimSpace(adminInput.Email)

	if adminInput.Email == "" || adminInput.Password == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request, fill all fields",
		})
		return
	}

	if !strings.Contains(adminInput.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if strings.HasPrefix(adminInput.Password, " ") || strings.HasSuffix(adminInput.Password, " ") || len(adminInput.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// check email exist in db or not 
	var checkAdmin models.Admin
	res := utils.PostgresDB.Model(&models.Admin{}).Where("email = ?", adminInput.Email).First(&checkAdmin)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request, admin doesnt exist",
		})
		return
	}

	// pass check
	if err := bcrypt.CompareHashAndPassword([]byte(checkAdmin.Password), []byte(adminInput.Password)); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// check if email nd phone verified
	if checkAdmin.EmailVerified != true || checkAdmin.PhoneVerified != true {
		c.JSON(400, gin.H{
			"msg": "invalid request, pls verify email or phone first",
		})
		return
	}

	// token gen
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": checkAdmin.ID,
		"role": checkAdmin.Role,
		"exp": time.Now().Add(24*time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token gen err",
		})
		return
	}


	// refresh token gen
	refreshToken := TokenGenerate(32)

	dbUpdate := utils.PostgresDB.Model(&models.Admin{}).Where("id = ?", checkAdmin.ID).Updates(map[string]any{
		"refresh_token": refreshToken,
		"refresh_expiry": time.Now().Add(7*24*time.Hour),
	})

	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid id, user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"accessToken": token,
		"refreshToken": refreshToken,
	})

}

// refresh token
func AdminRefreshToken (c*gin.Context){
	type AdminRefreshToken struct {
		RefreshToken string `json:"refreshToken"`
	}

	var adminInput AdminRefreshToken
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// check if rt exist or not 
	var checkToken models.Admin
	res := utils.PostgresDB.Model(&models.Admin{}).Where("refresh_token = ?", adminInput.RefreshToken).First(&checkToken)
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

	// gen new acces token
	newToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": checkToken.ID,
		"role": checkToken.Role,
		"exp": time.Now().Add(24*time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token gen err",
		})
		return
	}

	// new refresh token 
	newRefreshToken := TokenGenerate(32)

	dbUpdate := utils.PostgresDB.Model(&models.Admin{}).Where("id = ?",checkToken.ID).Updates(map[string]any{
		"refresh_token": newRefreshToken,
		"refresh_expiry": time.Now().Add(7*24*time.Hour),
	})

	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid id, admin not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"accessToken": newToken,
		"refreshToken": newRefreshToken,
	})
}

// forgot password api 
func AdminForgotPassword(c*gin.Context){
	type AdminForgotPassword struct {
		Email string `json:"email"`
	}

	var adminInput AdminForgotPassword
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validatons
	adminInput.Email = strings.TrimSpace(adminInput.Email)

	if adminInput.Email == "" || !strings.Contains(adminInput.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	// check email exists
	var checkEmail models.Admin
	ress := utils.PostgresDB.Model(&models.Admin{}).Where("email = ?",adminInput.Email).First(&checkEmail)
	if ress.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid admin, email doesnt exist ",
		})
		return
	}

	// generate temp pass
	tempPass := TokenGenerate(32)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(tempPass), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid bcrypt err",
		})
		return
	}

	// db update 
	dbUpdate := utils.PostgresDB.Model(&models.Admin{}).Where("id = ?", checkEmail.ID).Updates(map[string]any{
		"password": string(hashPass),
	})

	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid id, admin not found",
		})
		return
	}

	utils.EmailQueue <- utils.EmailData{
		To: adminInput.Email,
				Subject: "Team Social",
				Html: fmt.Sprintf(
					`<h3>Hello %s, your temporary password is <b>%s</b></h3>`,checkEmail.FullName, tempPass),
	}

	c.JSON(200, gin.H{
		"msg": "Temp email sent successfully!✅",
	})
}