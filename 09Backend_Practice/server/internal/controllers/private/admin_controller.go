package private

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/models"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/utils"
	"github.com/gin-gonic/gin"
)

// get all users
func GetAllUsers(c*gin.Context) {
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

	offset := (page - 1) *limit

	cacheKey := fmt.Sprintf("admin:all:users:%d:%d",page,limit)

	// check redis first 
	cacheData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cacheData != "" {
		var users []models.User
		if json.Unmarshal([]byte(cacheData), &users) == nil {
			c.JSON(200, gin.H{
				"src":"redis",
				"users":users,
				"page":page,
				"limit":limit,
			})
			return
		}
	}

	// check db 
	var users []models.User
	checkDb := utils.PostgresDB.Model(&models.User{}).Order("created_at DESC").Offset(offset).Limit(limit).Find(&users)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg" : "invalid db err, no users found",
		})
		return
	}

	// set cache key and db value in redis 
	bytes, _ := json.Marshal(users)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)
	if err != nil {
		log.Printf("redis err: %v",err)
	}

	c.JSON(200, gin.H{
		"src": "db",
		"users": users,
		"page": page,
		"limit": limit,
	})
}

func GetOneUser(c*gin.Context) {
	userId := c.Param("userId")

	// cache key 
	cacheKey := "admin:user:"+userId

	// check redis first 
	cacheData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cacheData != "" {
		var user models.User
		if json.Unmarshal([]byte(cacheData), &user) == nil {
			c.JSON(200, gin.H{
				"src":"redis",
				"user":user,

			})
			return
		}
	}

	// check db 
	var user models.User
	checkDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?",userId).First(&user)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg" : "invalid db err, no users found",
		})
		return
	}

	// set cache key and db value in redis 
	bytes, _ := json.Marshal(user)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)
	if err != nil {
		log.Printf("redis err: %v",err)
	}

	c.JSON(200, gin.H{
		"src": "db",
		"user": user,
	})
}

// func delete one user
func DeleteOneUser(c*gin.Context) {
	// take user id from param
	userId := c.Param("userId")

	// check if user exists in db 
	var checkUser models.User
	checkDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?",userId).First(&checkUser)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "err, no user found",
		})
		return
	}

	// delete from db 
	deleteUser := utils.PostgresDB.Model(&models.User{}).Delete(&checkUser)
	if deleteUser.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err",
		})
		return
	}

	if deleteUser.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid user id not found",
		})
		return
	}

	// redis invalidate feed 
	_ = utils.RedisDelKey("user:me:"+userId)
	_ = utils.RedisDelKey("user:public:"+userId)

	c.JSON(200, gin.H{"msg": "user deleted✅"})
}


// get all users
func GetAllPostsAdmin(c*gin.Context) {
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

	offset := (page - 1) *limit

	cacheKey := fmt.Sprintf("admin:all:post:%d:%d",page,limit)

	// check redis first 
	cacheData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cacheData != "" {
		var posts []models.Post
		if json.Unmarshal([]byte(cacheData), &posts) == nil {
			c.JSON(200, gin.H{
				"src":"redis",
				"post":posts,
				"page":page,
				"limit":limit,
			})
			return
		}
	}

	// check db 
	var posts []models.Post
	checkDb := utils.PostgresDB.Model(&models.Post{}).Order("created_at DESC").Offset(offset).Limit(limit).Find(&posts)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg" : "invalid db err, no users found",
		})
		return
	}

	// set cache key and db value in redis 
	bytes, _ := json.Marshal(posts)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)
	if err != nil {
		log.Printf("redis err: %v",err)
	}

	c.JSON(200, gin.H{
		"src": "db",
		"posts": posts,
		"page": page,
		"limit": limit,
	})
}

func GetOnePostAdmin(c*gin.Context) {
	postId := c.Param("postId")

	// cache key 
	cacheKey := "admin:post:"+postId

	// check redis first 
	cacheData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cacheData != "" {
		var post models.Post
		if json.Unmarshal([]byte(cacheData), &post) == nil {
			c.JSON(200, gin.H{
				"src":"redis",
				"post":post,

			})
			return
		}
	}

	// check db 
	var post models.Post
	checkDb := utils.PostgresDB.Model(&models.Post{}).Where("id = ?",postId).First(&post)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg" : "invalid db err, no post found",
		})
		return
	}

	// set cache key and db value in redis 
	bytes, _ := json.Marshal(post)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)
	if err != nil {
		log.Printf("redis err: %v",err)
	}

	c.JSON(200, gin.H{
		"src": "db",
		"post": post,
	})
}

// func delete one user
func DeleteOnePost(c*gin.Context) {
	// take user id from param
	postId := c.Param("postId")

	// check if user exists in db 
	var checkPost models.Post
	checkDb := utils.PostgresDB.Model(&models.Post{}).Where("id = ?",postId).First(&checkPost)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "err, no post found",
		})
		return
	}

	// delete from db 
	deletePost := utils.PostgresDB.Model(&models.Post{}).Delete(&checkPost)
	if deletePost.Error != nil {
		c.JSON(400, gin.H{
			"msg": "db err",
		})
		return
	}

	if deletePost.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"msg": "invalid post id not found",
		})
		return
	}

	// redis invalidate feed 
	invalidateFeedCache()
	_ = utils.RedisDelKey("post:"+postId)

	c.JSON(200, gin.H{"msg": "post deleted✅"})
}








