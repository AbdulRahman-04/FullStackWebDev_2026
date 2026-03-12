package public

import (
	"fmt"
	"strings"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/models"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Admin signup api
func AdminSignup(c *gin.Context) {
	type AdminSignup struct {
		AdminName string `json:"admin_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Phone     string `json:"phone"`
	}

	var adminInput AdminSignup
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	adminInput.AdminName = strings.TrimSpace(adminInput.AdminName)
	adminInput.Email = strings.TrimSpace(adminInput.Email)
	adminInput.Phone = strings.TrimSpace(adminInput.Phone)

	if adminInput.AdminName == "" || adminInput.Email == "" || adminInput.Password == "" || adminInput.Phone == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request,fill all fields",
		})
		return
	}

	if !strings.Contains(adminInput.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(adminInput.Password) < 6 || strings.HasPrefix(adminInput.Password, " ") || strings.HasSuffix(adminInput.Password, " ") {
		c.JSON(400, gin.H{
			"msg": "invalid password type",
		})
		return
	}

	var count int64
	checkDb := utils.PostgresDB.Model(&models.Admin{}).Where("email = ?", adminInput.Email).Count(&count)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "user already exists",
		})
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(adminInput.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid bcrypt er",
		})
		return
	}

	emailToken := TokenGenerate(32)
	OTP := OTPGenerate()

	var admin models.Admin
	admin.AdminName = adminInput.AdminName
	admin.Email = adminInput.Email
	admin.Password = string(hashPass)
	admin.Phone = adminInput.Phone
	admin.Provider = "email"
	admin.Email_Verify_Token = emailToken

	pushDb := utils.PostgresDB.Model(&models.Admin{}).Create(&admin)
	if pushDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	// set otp in redis for 2 mins
	err = utils.RedisSetKey("otp_"+adminInput.Phone, OTP, 2*time.Minute)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid redis err",
		})
		return
	}

	// send email nd sms
	utils.EmailQueue <- utils.EmailData{
		To:      adminInput.Email,
		Subject: "Email Verification",
		Text:    "Verify email",
		Html: fmt.Sprintf(
			`<p>Click below to verify:</p>
         <a href="%s/api/public/user/emailverify/%s">Verify Email</a>`, config.AppConfig.URL, emailToken),
	}

	utils.SMSQueue <- utils.SMSData{
		To:   adminInput.Phone,
		Body: "otp: " + OTP,
	}

	c.JSON(200, gin.H{
		"msg": "admin signed up successfully!✅",
	})
}

// email verify
func AdminEmailVerify(c *gin.Context) {
	token := c.Param("token")

	// check token in db and update
	checkDb := utils.PostgresDB.Model(&models.Admin{}).Where("email_verify_token = ? AND email_verified = false", token).Updates(map[string]any{
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
		"msg": "email verified🎉",
	})
}

// phone verify
func AdminPhoneVerify(c *gin.Context) {
	type AdminPhoneVerify struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}

	var adminInput AdminPhoneVerify
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	adminInput.Phone = strings.TrimSpace(adminInput.Phone)
	adminInput.Otp = strings.TrimSpace(adminInput.Otp)

	if adminInput.Phone == "" || adminInput.Otp == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request, fill all fields",
		})
		return
	}

	// check if phone exists in db
	var checkPhone models.Admin
	checkDb := utils.PostgresDB.Model(&models.Admin{}).Where("phone = ?", adminInput.Phone).First(&checkPhone)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid err, phone number not found",
		})
		return
	}

	// get otp from redis for this user number
	otp, err := utils.RedisGetKey("otp_" + adminInput.Phone)
	if err != nil || otp != adminInput.Otp {
		c.JSON(400, gin.H{
			"msg": "invalid or expired otp",
		})
		return
	}

	// db update
	update := utils.PostgresDB.Model(&models.Admin{}).Where("id = ?", checkPhone.ID).Updates(map[string]any{
		"phone_verified": true,
	})

	if update.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if update.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid userid not found",
		})
		return
	}

	// redis del key
	_ = utils.RedisDelKey("otp_" + adminInput.Phone)

	c.JSON(200, gin.H{
		"msg": "phone number verified✅",
	})
}

// sign in
func AdminSignin(c *gin.Context) {
	type AdminSignin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var adminInput AdminSignin
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	adminInput.Email = strings.TrimSpace(adminInput.Email)
	adminInput.Password = strings.TrimSpace(adminInput.Password)

	if adminInput.Email == "" || adminInput.Password == "" {
		c.JSON(400, gin.H{
			"msg": "invalid req, fill all field",
		})
		return
	}

	if !strings.Contains(adminInput.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(adminInput.Password) < 6 || strings.HasPrefix(adminInput.Password, " ") || strings.HasSuffix(adminInput.Password, " ") {
		c.JSON(400, gin.H{
			"msg": "invalid password type",
		})
		return
	}

	// check if email exists in db
	var checkEmail models.Admin
	checkDb := utils.PostgresDB.Model(&models.Admin{}).Where("email = ?", adminInput.Email).First(&checkEmail)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid req, no email found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(checkEmail.Password), []byte(adminInput.Password)); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// email/phone verified check nahi kiya!
	if !checkEmail.EmailVerified || !checkEmail.PhoneVerified {
		c.JSON(400, gin.H{"msg": "pls verify email and phone first"})
		return
	}
	// ye bcrypt check ke baad, token gen se pehle lagao

	// generate accessstoken and refresh token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   checkEmail.ID,
		"role": checkEmail.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token generation err",
		})
		return
	}

	refreshToken := TokenGenerate(32)

	updateDb := utils.PostgresDB.Model(&models.Admin{}).Where("id = ?", checkEmail.ID).Updates(map[string]any{
		"refresh_token":  refreshToken,
		"refresh_expiry": time.Now().Add(7 * 24 * time.Hour),
	})

	if updateDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if updateDb.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid err, userid not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":          "logged in",
		"accessToken":  token,
		"refreshToken": refreshToken,
	})

}

// refresh token api
func AdminRefreshToken(c *gin.Context) {
	type AdminRefreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}

	var adminInput AdminRefreshToken
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	adminInput.RefreshToken = strings.TrimSpace(adminInput.RefreshToken)

	if adminInput.RefreshToken == "" {
		c.JSON(400, gin.H{
			"msg": "invalid request, fill all fields",
		})
		return
	}

	// check if rt exists in db and its validity
	var checkToken models.Admin
	checkkDb := utils.PostgresDB.Model(&models.Admin{}).Where("refresh_token = ?", adminInput.RefreshToken).First(&checkToken)
	if checkkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid refresh token not found",
		})
		return
	}

	if time.Now().After(checkToken.RefreshExpiry) {
		c.JSON(400, gin.H{
			"msg": "refresh token expired",
		})
		return
	}

	// generate new access token nd refresh token
	newToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   checkToken.ID,
		"role": checkToken.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token err",
		})
		return
	}

	newRefreshToken := TokenGenerate(32)

	dbUpdate := utils.PostgresDB.Model(&models.Admin{}).Where("id = ?", checkToken.ID).Updates(map[string]any{
		"refresh_token":  newRefreshToken,
		"refresh_expiry": time.Now().Add(7 * 24 * time.Hour),
	})

	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid user id not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"accessToken":  newToken,
		"refreshToken": newRefreshToken,
	})
}

// forgot pass api
func AdminForgotPass(c *gin.Context) {
	type AdminForgotPass struct {
		Email string `json:"email"`
	}

	var adminInput AdminForgotPass
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	adminInput.Email = strings.TrimSpace(adminInput.Email)

	if adminInput.Email == "" || !strings.Contains(adminInput.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email type",
		})
		return
	}

	// check email in db
	var checkEmail models.Admin
	checkDb := utils.PostgresDB.Model(&models.Admin{}).Where("email = ?", adminInput.Email).First(&checkEmail)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid email not found",
		})
		return
	}

	newPass := TokenGenerate(10)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(newPass), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid bcrypt err",
		})
		return
	}

	// update db
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
			"msg": "invalid user id not found",
		})
		return
	}

	// send pass on email

	utils.EmailQueue <- utils.EmailData{
		To:      adminInput.Email,
		Subject: "Team Social",
		Html: fmt.Sprintf(
			`<h3>Hello %s, your temporary password is <b>%s</b></h3>`, checkEmail.AdminName, newPass),
	}

	c.JSON(200, gin.H{
		"msg": "temp pass sent to ur email✅",
	})
}
