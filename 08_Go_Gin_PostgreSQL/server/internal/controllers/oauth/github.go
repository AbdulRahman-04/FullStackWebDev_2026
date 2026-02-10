package oauth

import (
	"encoding/json"
	"fmt"
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
func GithubLoginUser(c *gin.Context) {
	c.Redirect(
		http.StatusTemporaryRedirect,
		utils.GithubOauthUser().AuthCodeURL("github_user"),
	)
}

func GithubLoginAdmin(c *gin.Context) {
	c.Redirect(
		http.StatusTemporaryRedirect,
		utils.GithubOauthAdmin().AuthCodeURL("github_admin"),
	)
}

// func GithubCallbackUser(c *gin.Context) {
// 	githubCallback(c, "user")
// }

// func GithubCallbackAdmin(c *gin.Context) {
// 	githubCallback(c, "admin")
// }

func GithubCallback(c *gin.Context, role string) {

	var oauthCfg *oauth2.Config
	if role == "user" {
		oauthCfg = utils.GithubOauthUser()
	} else {
		oauthCfg = utils.GithubOauthAdmin()
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(400, gin.H{"error": "missing code"})
		return
	}

	token, err := oauthCfg.Exchange(c, code)
	if err != nil {
		c.JSON(500, gin.H{"error": "token exchange failed"})
		return
	}

	client := oauthCfg.Client(c, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(500, gin.H{"error": "failed fetching github user"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var gh map[string]any
	_ = json.Unmarshal(body, &gh)

	email := fmt.Sprintf("%v", gh["email"])
	name := fmt.Sprintf("%v", gh["login"])
	avatar := fmt.Sprintf("%v", gh["avatar_url"])

	db := utils.PostgresDB

	// ============================
	// ADMIN FLOW
	// ============================
	if role == "admin" {
		var admin models.Admin
		err := db.Where("email = ?", email).First(&admin).Error

		if err != nil {
			admin = models.Admin{
				FullName:         name,
				Email:            email,
				Provider:         "github",
				Role:             "admin",

				// 👇 explicit defaults (OPTION 2)
				EmailVerified:    true,
				PhoneVerified:    false,
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
			"provider": "github",
			"role":     "admin",
			"token":    jwtToken,
			"profile": gin.H{
				"name":   name,
				"email":  email,
				"avatar": avatar,
			},
		})
		return
	}

	// ============================
	// USER FLOW
	// ============================
	var user models.User
	err = db.Where("email = ?", email).First(&user).Error

	if err != nil {
		user = models.User{
			FullName:        name,
			Email:           email,
			Provider:        "github",
			Role:            "user",
			ProfilePicture:  avatar,

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
		"provider": "github",
		"role":     "user",
		"token":    jwtToken,
		"profile": gin.H{
			"name":   name,
			"email":  email,
			"avatar": avatar,
		},
	})
}
