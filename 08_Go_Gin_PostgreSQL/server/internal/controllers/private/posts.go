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

// create post api
func CreatePost(c *gin.Context) {
	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(400, gin.H{
			"msg": "unauthorised",
		})
		return
	}

	userId := userIDInterface.(string)

	// make struct
	type CreatePost struct {
		Caption  string `json:"caption" form:"caption" binding:"required"`
		Song     string `json:"song" form:"song" binding:"required"`
		ImageUrl string `json:"image_url"`
		Location string `json:"location" form:"location" binding:"required"`
		IsPublic bool   `json:"is_public" form:"is_public" binding:"required"`
	}

	var postInput CreatePost
	if err := c.ShouldBind(&postInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	var image_url string
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
				"msg": "file upload failed",
			})
			return
		}

		image_url = url
	} else {
		if postInput.ImageUrl == "" {
			c.JSON(400, gin.H{
				"msg": "image_url required",
			})
			return
		}

		image_url = postInput.ImageUrl
	}

	// create post
	var post models.Post
	post.Caption = postInput.Caption
	post.Song = postInput.Song
	post.ImageUrl = image_url
	post.Location = postInput.Location
	post.IsPublic = postInput.IsPublic
	post.UserID = userId   // 🔥 THIS WAS MISSING

	db := utils.PostgresDB.Model(&models.Post{}).Create(&post)
	if db.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	// invalidate redis cache - ye user k purane posts del krdo redis se
	go invalidateFeedCache(userId)

	c.JSON(200, gin.H{
		"msg": "post created✅",
	})
}

func invalidateFeedCache(userId string) {
	key := "post:" + userId
	if err := utils.RedisDelKey(key); err != nil {
		log.Printf("%s %v", userId, err)
	}
}

// get all posts
func GetAllPosts(c *gin.Context) {
	// page and limit and offset
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// off set
	offset := (page - 1) * limit

	cachedKey := fmt.Sprintf("posts:feed:%d:%d", page, limit)

	// check redis
	cachedData, err := utils.RedisGetKey(cachedKey)
	if err == nil && cachedData != "" {
		var posts []models.Post
		if json.Unmarshal([]byte(cachedData), &posts) == nil {
			c.JSON(200, gin.H{
				"src":   "redis",
				"page":  page,
				"limit": limit,
				"data": posts,
			})
			return
		}
	}

	// if key isnt there in redis then check db
	var posts []models.Post
	db := utils.PostgresDB.Model(&models.Post{}).Limit(limit).Offset(offset).Find(&posts)
	if db.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err",
		})
		return
	}

	// set redis key
	bytes, _ := json.Marshal(posts)
	go utils.RedisSetKey(cachedKey, string(bytes), 5*time.Hour)

	c.JSON(200, gin.H{
		"src":   "db",
		"page":  page,
		"limit": limit,
		"data":  posts,
	})
}

// get one api
func GetOnePost(c *gin.Context) {

	postId := c.Param("postId")

	cacheKey := "post:" + postId

	// redis check
	cachedData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cachedData != "" {
		var post models.Post
		if json.Unmarshal([]byte(cachedData), &post) == nil {
			c.JSON(200, gin.H{
				"src":  "redis",
				"post": post,
			})
			return
		}
	}

	// db check
	var post models.Post
	db := utils.PostgresDB.Model(&models.Post{}).First(&post)
	if db.Error != nil {
		c.JSON(400, gin.H{
			"msg": "err",
		})
		return
	}

	// set in redis
	bytes, _ := json.Marshal(post)
	go utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)

	c.JSON(200, gin.H{
		"src":  "db",
		"post": post,
	})

}


func UpdatePost(c *gin.Context) {

	userID := c.MustGet("userId").(string)

	// convert id to int (DB me bigint hai)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": "invalid post id"})
		return
	}

	// 🔹 Check post exists
	var post models.Post
	if err := utils.PostgresDB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"msg": "post not found"})
		return
	}

	// 🔹 Ownership check
	if post.UserID != userID {
		c.JSON(403, gin.H{"msg": "forbidden"})
		return
	}

	// 🔹 Input struct
	type UpdatePostInput struct {
		Caption  *string `json:"caption"`
		ImageURL *string `json:"image_url"`
		Location *string `json:"location"`
		IsPublic *bool   `json:"is_public"`
	}

	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"msg": "invalid json"})
		return
	}

	updates := map[string]any{
		"updated_at": time.Now(),
	}

	if input.Caption != nil {
		updates["caption"] = *input.Caption
	}
	if input.ImageURL != nil {
		updates["image_url"] = *input.ImageURL
	}
	if input.Location != nil {
		updates["location"] = *input.Location
	}
	if input.IsPublic != nil {
		updates["is_public"] = *input.IsPublic
	}

	if len(updates) == 1 {
		c.JSON(400, gin.H{"msg": "nothing to update"})
		return
	}

	// 🔹 Update only this row
	if err := utils.PostgresDB.Model(&post).Updates(updates).Error; err != nil {
		c.JSON(500, gin.H{"msg": "update failed"})
		return
	}

	// 🔹 Cache invalidation
	go func() {
		utils.RedisDelKey("post:" + strconv.Itoa(id))
		invalidateFeedCache(userID)
	}()

	utils.PostgresDB.First(&post, id)

	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {

	userID := c.MustGet("userId").(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": "invalid post id"})
		return
	}

	var post models.Post
	if err := utils.PostgresDB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"msg": "post not found"})
		return
	}

	if post.UserID != userID {
		c.JSON(403, gin.H{"msg": "forbidden"})
		return
	}

	if err := utils.PostgresDB.Delete(&post).Error; err != nil {
		c.JSON(500, gin.H{"msg": "delete failed"})
		return
	}

	go func() {
		utils.RedisDelKey("post:" + strconv.Itoa(id))
		invalidateFeedCache(userID)
	}()

	c.JSON(200, gin.H{"msg": "deleted"})
}
