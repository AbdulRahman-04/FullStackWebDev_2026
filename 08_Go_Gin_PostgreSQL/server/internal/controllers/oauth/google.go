package oauth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/models"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

// ---------------------------
// LOGIN URLS
// ---------------------------
func GoogleLoginUser(c *gin.Context) {
	c.Redirect(
		http.StatusTemporaryRedirect,
		utils.GoogleOauthUser().AuthCodeURL("google_user"),
	)
}

func GoogleLoginAdmin(c *gin.Context) {
	c.Redirect(
		http.StatusTemporaryRedirect,
		utils.GoogleOauthAdmin().AuthCodeURL("google_admin"),
	)
}

// func GoogleCallbackUser(c *gin.Context) {
// 	googleCallback(c, "user")
// }

// func GoogleCallbackAdmin(c *gin.Context) {
// 	googleCallback(c, "admin")
// }

func GoogleCallback(c *gin.Context, role string) {

	var oauthCfg *oauth2.Config
	if role == "user" {
		oauthCfg = utils.GoogleOauthUser()
	} else {
		oauthCfg = utils.GoogleOauthAdmin()
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(400, gin.H{"error": "missing code"})
		return
	}

	token, err := oauthCfg.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(500, gin.H{"error": "token exchange failed"})
		return
	}

	client := oauthCfg.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(500, gin.H{"error": "failed fetching google user"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var gu struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	_ = json.Unmarshal(body, &gu)

	db := utils.PostgresDB

	// ============================
	// ADMIN FLOW
	// ============================
	if role == "admin" {
		var admin models.Admin
		err := db.Where("email = ?", gu.Email).First(&admin).Error

		if err != nil {
			admin = models.Admin{
				FullName:       gu.Name,
				Email:          gu.Email,
				Provider:       "google",
				Role:           "admin",
				EmailVerified:  true,
				PhoneVerified:  false,
				EmailVerifyToken: "",
				PhoneVerifyToken: "",
				RefreshToken:     "",
			}
			db.Create(&admin)
		}

		claims := jwt.MapClaims{
			"id":   admin.ID,
			"role": "admin",
			"exp":  time.Now().Add(24 * time.Hour).Unix(),
		}
		tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		jwtToken, _ := tokenJWT.SignedString([]byte(config.AppConfig.JWT_KEY))

		c.JSON(200, gin.H{
			"success":  true,
			"provider": "google",
			"role":     "admin",
			"token":    jwtToken,
			"profile": gin.H{
				"name":   gu.Name,
				"email":  gu.Email,
				"avatar": gu.Picture,
			},
		})
		return
	}

	// ============================
	// USER FLOW
	// ============================
	var user models.User
	err = db.Where("email = ?", gu.Email).First(&user).Error

	if err != nil {
		user = models.User{
			FullName:        gu.Name,
			Email:           gu.Email,
			Provider:        "google",
			Role:            "user",
			ProfilePicture:  gu.Picture,

			// 👇 explicit defaults (OPTION 2)
			EmailVerified:   true,
			PhoneVerified:   false,
			EmailVerifyToken: "",
			PhoneVerifyToken: "",
			RefreshToken:     "",
		}
		db.Create(&user)
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": "user",
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := tokenJWT.SignedString([]byte(config.AppConfig.JWT_KEY))

	c.JSON(200, gin.H{
		"success":  true,
		"provider": "google",
		"role":     "user",
		"token":    jwtToken,
		"profile": gin.H{
			"name":   gu.Name,
			"email":  gu.Email,
			"avatar": gu.Picture,
		},
	})
}