// package private

// import (
// 	"log"
// 	"net/http"
// 	"strings"

// 	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/models"
// 	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
// 	"github.com/gin-gonic/gin"
// )

// var ctx = context.Background()

// func CreatePost(c *gin.Context) {

// 	// safer than MustGet
// 	userIDInterface, exists := c.Get("userId")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
// 		return
// 	}
// 	userID := userIDInterface.(string)

// 	// Struct supports BOTH json + multipart form
// 	type CreatePostInput struct {
// 		Caption  string `json:"caption" form:"caption" binding:"required"`
// 		Song     string `json:"song" form:"song" binding:"required"`
// 		Location string `json:"location" form:"location" binding:"required"`
// 		IsPublic bool   `json:"is_public" form:"is_public" binding:"required"`
// 		ImageUrl string `json:"image_url"` // required only for JSON
// 	}

// 	var input CreatePostInput

// 	// Automatically binds JSON or form-data
// 	if err := c.ShouldBind(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid input"})
// 		return
// 	}

// 	var imageURL string

// 	// If multipart request → upload file
// 	if strings.HasPrefix(c.ContentType(), "multipart/") {

// 		// Ensure file exists
// 		_, err := c.FormFile("file")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"msg": "file required"})
// 			return
// 		}

// 		url, err := utils.UploadFile(c)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"msg": "upload failed"})
// 			return
// 		}
// 		imageURL = url

// 	} else {
// 		// JSON request → image_url required
// 		if input.ImageUrl == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"msg": "image_url required"})
// 			return
// 		}
// 		imageURL = input.ImageUrl
// 	}

// 	post := models.Post{
// 		UserID:   userID,
// 		Caption:  input.Caption,
// 		Song:     input.Song,
// 		ImageUrl: imageURL,
// 		Location: input.Location,
// 		IsPublic: input.IsPublic,
// 	}

// 	if err := utils.PostgresDB.Create(&post).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"})
// 		return
// 	}

// 	// invalidate feed cache (same function)
// 	go invalidateFeedCache(userID)

// 	c.JSON(http.StatusCreated, post)
// }

// func invalidateFeedCache(userID string) {
// 	key := "feed:" + userID
// 	if err := utils.RedisDelKey(key); err != nil {
// 		log.Printf("failed to invalidate feed cache for user %s: %v", userID, err)
// 	}
// }

package private

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/models"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func CreatePost(c *gin.Context) {
	// get user id
	userIDInterface, exists := c.Get("userId")
	if !exists || userIDInterface == "" {
		c.JSON(400, gin.H{
			"msg": "no user id found",
		})
		return
	}

	userId := userIDInterface.(string)

	// struct
	type CreatePost struct {
		Caption  string `json:"caption" form:"caption" binding:"required"`
		Song     string `json:"song" form:"song" binding:"required"`
		ImageUrl string `json:"image_url"`
		Location string `json:"location" binding:"required"`
		IsPublic bool   `json:"is_public" form:"is_public" binding:"required"`
	}

	var postInput CreatePost
	if err := c.ShouldBind(&postInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	var imageUrl string

	if strings.HasPrefix(c.ContentType(), "multipart/") {
		_, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"msg": "file required",
			})
			return
		}

		url, err := utils.UploadFile(c)
		if err != nil {
			c.JSON(400, gin.H{
				"msg": "couldn't upload file",
			})
			return
		}

		imageUrl = url
	}

	// create var nd push in db
	var post models.Post

	post.Caption = postInput.Caption
	post.Song = postInput.Song
	post.IsPublic = postInput.IsPublic
	post.Location = postInput.Location
	post.ImageUrl = imageUrl

	insertDb := utils.PostgresDB.Model(&models.Post{}).Create(&post)
	if insertDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db error",
		})
		return
	}

	// redis cache invalidate
	go invalidateFeedCache(userId)
}

func invalidateFeedCache(userId string) {
	key := "feed:" + userId
	if err := utils.RedisDelKey(key); err != nil {
		log.Printf("couldnt invalidate cache for user %s and err %v", userId, err)
	}
}

// get all posts 
func GetAllPosts(c*gin.Context){
	// make page : page depends on total value in db nd limit 
	// e.g total post : 58 and limit 10 then pages are 6 
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) *limit

	// create unique redis key
	cachedKey := fmt.Sprintf("feed:%d:%d",page,limit)

	// check redis first
	cached, err := utils.RedisGetKey(cachedKey)
	if err == nil && cached != "" {
		var posts []models.Post
		if json.Unmarshal([]byte(cached), &posts) == nil {
			c.JSON(200, gin.H{
				"source":"redis",
				"page": page,
				"limit": limit,
				"data": posts,
			})
			return
		}
	}

	var posts []models.Post
	db := utils.PostgresDB.Order("created_at DESC").Limit(limit).Offset(offset).Find(&posts)
	if db.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err",
		})
		return
	}

	// redis set 
	bytes, err := json.Marshal(posts)
	if err != nil {
		log.Printf("couldnt into json")
		return
	}

	go utils.RedisSetKey(cachedKey, string(bytes), 60*time.Second)
	c.JSON(200, gin.H{
		"source": "db",
		"page": page,
		"limit": limit,
		"data": posts,
	})
}