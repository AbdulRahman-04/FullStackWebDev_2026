package private

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/models"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/utils"
	"github.com/gin-gonic/gin"
)

// create post
func CreatePost(c *gin.Context) {
	// get user id
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(400, gin.H{
			"msg": "unauthorised",
		})
		return
	}

	userId, ok := userIdInterface.(string)
	if !ok {
		c.JSON(400, gin.H{
			"msg": "invalid userId conversion",
		})
		return
	}

	// struct to take data in json
	type CreatePost struct {
		Caption  string `json:"caption" form:"caption"`
		Song     string `json:"song" form:"song"`
		ImageUrl string `json:"image_url"`
		Location string `json:"location"`
		IsPublic bool   `json:"is_public"`
	}

	var postInput CreatePost
	if err := c.ShouldBind(&postInput); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	var ImageUrl string
	if strings.HasPrefix(c.ContentType(), "multipart/form-data") {
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

		ImageUrl = url
	} else {
		if postInput.ImageUrl == "" {
			c.JSON(400, gin.H{
				"msg": "image_url required",
			})
			return
		}

		ImageUrl = postInput.ImageUrl
	}

	// push in db as a row
	var post models.Post
	post.UserID = userId
	post.Caption = postInput.Caption
	post.Song = postInput.Song
	post.ImageUrl = ImageUrl
	post.Location = postInput.Location
	post.IsPublic = postInput.IsPublic

	createPost := utils.PostgresDB.Model(&models.Post{}).Create(&post)
	if createPost.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	// feed invalidate cache for the user
	invalidateFeedCache()

	c.JSON(200, gin.H{
		"msg": "post created✅",
	})
}

func invalidateFeedCache() {
	utils.RedisDelKey("post:feed:1:10")
	utils.RedisDelKey("post:feed:2:10")
	utils.RedisDelKey("post:feed:3:10")
	utils.RedisDelKey("post:feed:4:10")
	utils.RedisDelKey("post:feed:5:10")
}

// get all posts
func GetAllPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 50 {
		limit = 50
	}

	// create offset
	offset := (page - 1) * limit
	cacheKey := fmt.Sprintf("post:feed:%d:%d", page, limit)

	// check redis first
	cachedData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cachedData != "" {
		var posts []models.Post
		if json.Unmarshal([]byte(cachedData), &posts) == nil {
			c.JSON(200, gin.H{
				"src":   "redis",
				"page":  page,
				"limit": limit,
				"posts": posts,
			})
			return
		}
	}

	// check db if redis fails
	var posts []models.Post
	checkDb := utils.PostgresDB.Model(&models.Post{}).Order("created_at DESC").Limit(limit).Offset(offset).Find(&posts)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "no data found in db",
		})
		return
	}

	// set cache key with db data in redis
	bytes, _ := json.Marshal(posts)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)

	if err != nil {
		log.Printf("redis err %v:", err)
	}

	c.JSON(200, gin.H{
		"src":   "db",
		"page":  page,
		"limit": limit,
		"posts": posts,
	})
}

// get one post
func GetOnePost(c *gin.Context) {
	// get post id from param
	postId := c.Param("postId")

	cacheKey := "post:" + postId

	// check redis first
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

	// get from db if not in redis
	var post models.Post
	getPost := utils.PostgresDB.Model(&models.Post{}).Where("id = ?", postId).First(&post)
	if getPost.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err, no post found",
		})
		return
	}

	// set cachekey nd db data into redis
	bytes, _ := json.Marshal(&post)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)

	if err != nil {
		log.Printf("redis err %v:", err)
	}
	c.JSON(200, gin.H{
		"src":  "db",
		"post": post,
	})
}

// update post func
func UpdatePost(c *gin.Context) {
	// get user id from token
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(400, gin.H{
			"msg": "unauthorised",
		})
		return
	}

	userId, ok := userIdInterface.(string)
	if !ok {
		c.JSON(400, gin.H{
			"msg": "invalid userId conversion",
		})
		return
	}

	// take post id from param nd check in db
	postId := c.Param("postId")

	var checkPost models.Post
	checkDb := utils.PostgresDB.Model(&models.Post{}).Where("id = ?", postId).First(&checkPost)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid post id",
		})
		return
	}

	// check ownership of post
	if checkPost.UserID != userId {
		c.JSON(400, gin.H{
			"msg": "couldn't update post",
		})
		return
	}

	// create struct for updating post
	type UpdatePost struct {
		Caption  *string `json:"caption"`
		Song     *string `json:"song"`
		ImageUrl *string `json:"image_url"`
		Location *string `json:"location"`
		IsPublic *bool   `json:"is_public"`
	}

	var updatePost UpdatePost
	if err := c.ShouldBindJSON(&updatePost); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// create updates
	updates := map[string]any{
		"updated_at": time.Now(),
	}

	if updatePost.Caption != nil {
		updates["caption"] = *updatePost.Caption
	}

	if updatePost.Song != nil {
		updates["song"] = *updatePost.Song
	}

	if updatePost.ImageUrl != nil {
		updates["image_url"] = *updatePost.ImageUrl
	}

	if updatePost.Location != nil {
		updates["location"] = *updatePost.Location
	}

	if updatePost.IsPublic != nil {
		updates["is_public"] = *updatePost.IsPublic
	}

	// check length if no updates passed in json form
	if len(updates) == 1 {
		c.JSON(400, gin.H{
			"msg": "ntohing to update",
		})
		return
	}

	// db Update
	dbUpdate := utils.PostgresDB.Model(&models.Post{}).Where("id = ?", checkPost.ID).Updates(updates)
	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{"msg": "invalid db err"})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{"msg": "invalid postId not found"})
		return
	}

	// del from redis get all nd get one
	_ = utils.RedisDelKey("post:" + postId)
	invalidateFeedCache()

	c.JSON(200, gin.H{
		"msg": "post updated✅",
	})

}

// delete post api
func DeletePost(c *gin.Context) {
	// get userId from token
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(400, gin.H{
			"msg": "unauthorised",
		})
		return
	}

	userId, ok := userIdInterface.(string)
	if !ok {
		c.JSON(400, gin.H{
			"msg": "invalid userId conversion",
		})
		return
	}

	// postid from token
	postId := c.Param("postId")

	// check post in db and ownership
	var checkPost models.Post
	checkDb := utils.PostgresDB.Model(&models.Post{}).Where("id = ?", postId).First(&checkPost)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err,post not found",
		})
		return
	}

	if checkPost.UserID != userId {
		c.JSON(400, gin.H{
			"msg": "unauthorised",
		})
		return
	}

	// delete from db
	deletePost := utils.PostgresDB.Model(&models.Post{}).Delete(&checkPost)
	if deletePost.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	// del from redis
	_ = utils.RedisDelKey("post:" + postId)
	invalidateFeedCache()

	c.JSON(200, gin.H{
		"msg": "post deleted✅",
	})

}
