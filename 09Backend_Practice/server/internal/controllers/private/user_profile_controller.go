package private

import (
	"encoding/json"
	"log"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/models"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/utils"
	"github.com/gin-gonic/gin"
)

// get my profile
func GetMyProfile(c *gin.Context) {
	// get userId from token
	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(400, gin.H{
			"msg": "unauthorised",
		})
		return
	}

	userId, ok := userIDInterface.(string)
	if !ok {
		c.JSON(400, gin.H{
			"msg": "conversion err",
		})
		return
	}

	// create cache key nd check redis
	cacheKey := "user:me:" + userId

	cachedData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cachedData != "" {
		var profile models.User
		if json.Unmarshal([]byte(cachedData), &profile) == nil {
			c.JSON(200, gin.H{
				"src":     "redis",
				"profile": profile,
			})
			return
		}
	}

	// check db
	var profile models.User
	checkDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?", userId).First(&profile)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid userid",
		})
		return
	}

	// set cache key nd db data in redis
	bytes, _ := json.Marshal(profile)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)
	if err != nil {
		log.Printf("redis err %v:", err)
	}

	c.JSON(200, gin.H{
		"src":     "db",
		"profile": profile,
	})
}

// get other user profile
func GetOtherUser(c *gin.Context) {
	// take userid from param
	userId := c.Param("userId")

	// create redis key and check in redis
	cacheKey := "user:public:" + userId

	cacheData, err := utils.RedisGetKey(cacheKey)
	if err == nil && cacheData != "" {
		var userPfp models.User
		if json.Unmarshal([]byte(cacheData), &userPfp) == nil {
			c.JSON(200, gin.H{
				"src":     "redis",
				"profile": userPfp,
			})
			return
		}
	}

	// check db
	var userPfp models.User
	checkDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?", userId).First(&userPfp)
	if checkDb.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid id no user found",
		})
		return
	}

	// set cache key and db data into redis
	bytes, _ := json.Marshal(userPfp)
	err = utils.RedisSetKey(cacheKey, string(bytes), 5*time.Hour)

	if err != nil {
		log.Printf("redis err %v:", err)
	}

	c.JSON(200, gin.H{
		"src":     "db",
		"profile": userPfp,
	})
}

func UpdateProfile(c *gin.Context) {
	// get userId from token
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(400, gin.H{"msg": "unauthorised"})
		return
	}

	userId, ok := userIdInterface.(string)
	if !ok {
		c.JSON(400, gin.H{"msg": "invalid userId conversion"})
		return
	}

	// struct with pointers - form + json dono
	type UpdateProfile struct {
		FullName *string `json:"full_name" form:"full_name"`
		Email    *string `json:"email" form:"email"`
		Phone    *string `json:"phone" form:"phone"`
	}

	var updateProfile UpdateProfile
	_ = c.ShouldBind(&updateProfile)

	// updates map
	updates := map[string]any{
		"updated_at": time.Now(),
	}

	if updateProfile.FullName != nil {
		updates["full_name"] = *updateProfile.FullName
	}

	// email duplicate check
	if updateProfile.Email != nil {
		var count int64
		utils.PostgresDB.Model(&models.User{}).
			Where("email = ? AND id <> ?", *updateProfile.Email, userId).
			Count(&count)
		if count > 0 {
			c.JSON(400, gin.H{"msg": "email already exists"})
			return
		}
		updates["email"] = *updateProfile.Email
	}

	// phone duplicate check
	if updateProfile.Phone != nil {
		var count int64
		utils.PostgresDB.Model(&models.User{}).
			Where("phone = ? AND id <> ?", *updateProfile.Phone, userId).
			Count(&count)
		if count > 0 {
			c.JSON(400, gin.H{"msg": "phone already exists"})
			return
		}
		updates["phone"] = *updateProfile.Phone
	}

	// profile picture upload
	file, _ := c.FormFile("file")
	if file != nil {
		url, err := utils.UploadFile(c)
		if err != nil {
			c.JSON(400, gin.H{"msg": "file upload failed"})
			return
		}
		updates["profile_picture"] = url
	}

	// nothing to update
	if len(updates) == 1 {
		c.JSON(400, gin.H{"msg": "nothing to update"})
		return
	}

	// db update
	dbUpdate := utils.PostgresDB.Model(&models.User{}).Where("id = ?", userId).Updates(updates)
	if dbUpdate.Error != nil {
		c.JSON(400, gin.H{"msg": "db err"})
		return
	}

	if dbUpdate.RowsAffected == 0 {
		c.JSON(400, gin.H{"msg": "user not found"})
		return
	}

	// redis cleanup
	_ = utils.RedisDelKey("user:me:" + userId)
	_ = utils.RedisDelKey("user:public:" + userId)

	c.JSON(200, gin.H{"msg": "profile updated✅"})
}

// delete profile
func DeleteMyProfile(c *gin.Context) {
	// get userid
	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"msg": "unauthorised"})
		return
	}

	userId, ok := userIDInterface.(string)
	if !ok {
		c.JSON(401, gin.H{"msg": "invalid user id format"})
		return
	}

	// delete user row from db
	deleteUser := utils.PostgresDB.Model(&models.User{}).Where("id = ?", userId).Delete(&models.User{})

	if deleteUser.Error != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db err",
		})
		return
	}

	if deleteUser.RowsAffected == 0 {
		c.JSON(401, gin.H{"msg": "invalid user id"})
		return
	}

	// invalidate from redis db
	_ = utils.RedisDelKey("user:me:" + userId)
	_ = utils.RedisDelKey("user:public:" + userId)

	c.JSON(200, gin.H{
		"msg": "user account deleted💔",
	})

}

// logout user
func LogOutUser(c *gin.Context) {
	// get token from id
	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"msg": "unauthorised"})
		return
	}

	userId, ok := userIDInterface.(string)
	if !ok {
		c.JSON(401, gin.H{"msg": "invalid user"})
		return
	}

	// update db
	updateDb := utils.PostgresDB.Model(&models.User{}).Where("id = ?", userId).Updates(map[string]any{
		"refresh_token":  "",
		"refresh_expiry": time.Time{},
	})

	if updateDb.Error != nil {
		c.JSON(401, gin.H{"msg": "db err"})
		return
	}

	if updateDb.RowsAffected == 0 {
		c.JSON(401, gin.H{"msg": "invalid user id"})
		return
	}

	// redis feed invalidate
	_ = utils.RedisDelKey("user:me:" + userId)

	c.JSON(200, gin.H{
		"msg": "logged out✅",
	})
}
