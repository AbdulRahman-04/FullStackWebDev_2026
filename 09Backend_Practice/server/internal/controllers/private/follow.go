package private

import (
	"errors"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/models"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// send follow req
func SendFollowRequest(c *gin.Context) {
	// get fromuserid nd to userid
	fromUserId := c.GetString("userId")
	toUserId := c.Param("id")

	if fromUserId == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if fromUserId == toUserId {
		c.JSON(400, gin.H{"error": "cannot follow yourself"})
		return
	}

	// check already following ?
	var checkFollow models.Follow
	checkDb := utils.PostgresDB.Model(&models.Follow{}).Where("follower_id = ? AND following_id = ?", fromUserId, toUserId).First(&checkFollow)
	if checkDb.Error == nil {
		c.JSON(400, gin.H{"error": "already following"})
		return
	}

	if checkDb.Error != nil && !errors.Is(checkDb.Error, gorm.ErrRecordNotFound) {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	// already requested ?
	var checkRequest models.FollowRequest
	checkReq := utils.PostgresDB.Model(&models.FollowRequest{}).Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).First(&checkRequest)
	if checkReq.Error == nil {
		c.JSON(400, gin.H{"error": "request already sent"})
		return
	}

	if checkReq.Error != nil && !errors.Is(checkReq.Error, gorm.ErrRecordNotFound) {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	// create new flw request
	var newReq models.FollowRequest
	newReq.FromUserID = fromUserId
	newReq.ToUserID = toUserId
	newReq.Status = "pending"

	createReq := utils.PostgresDB.Model(&models.FollowRequest{}).Create(&newReq)
	if createReq.Error != nil {
		c.JSON(500, gin.H{"error": "could not send request"})
		return
	}

	// resp
	c.JSON(200, gin.H{
		"msg": "follow request sent✅",
	})
}

// accept request
func AcceptRequest(c *gin.Context) {
	// current user jo req accept krra uski id token se nikalo
	userId := c.GetString("userId")

	if userId == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// kis user ki req accpt krna h wo id param se lo
	requestID := c.Param("id")

	tx := utils.PostgresDB.Begin()
	if tx.Error != nil {
		c.JSON(400, gin.H{
			"msg": "transaction start failed",
		})
		return
	}

	defer tx.Rollback()

	// row lock and validate
	var req models.FollowRequest
	getReq := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ? AND to_user_id = ? AND status = ?", requestID, userId, "pending").First(&req)
	if getReq.Error != nil {
		c.JSON(404, gin.H{"error": "request not found"})
		return
	}

	// create follow
	var follow models.Follow
	follow.FollowerID = req.FromUserID
	follow.FollowingID = req.ToUserID

	createReq := tx.Create(&follow)
	if createReq.Error != nil && !errors.Is(createReq.Error, gorm.ErrDuplicatedKey) {
		c.JSON(500, gin.H{"error": "could not follow"})
		return
	}

	// delete row from flw request table
	deleteReq := tx.Delete(&req)
	if deleteReq.Error != nil {
		c.JSON(500, gin.H{"error": "could not delete request"})
		return
	}

	// comit tx
	commitDb := tx.Commit()
	if commitDb.Error != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "accepted",
	})
}

// reject req
func RejectRequest(c *gin.Context) {
	userId := c.GetString("userId")
	requestId := c.Param("id")

	if userId == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// st tx
	tx := utils.PostgresDB.Begin()
	if tx.Error != nil {
		c.JSON(401, gin.H{"error": "tx start failed"})
		return
	}

	defer tx.Rollback()

	// row lock and validate
	var req models.FollowRequest
	getReq := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ? AND to_user_id = ? AND status = ?", requestId, userId, "pending").First(&req)
	if getReq.Error != nil {
		c.JSON(404, gin.H{"error": "request not found"})
		return
	}

	// del req row
	delReq := tx.Delete(&req)
	if delReq.Error != nil {
		c.JSON(500, gin.H{"error": "could not reject"})
		return
	}

	// commit tx
	commitDb := tx.Commit()
	if commitDb.Error != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "req rejected✅",
	})
}

// get my follow request
func GetFollowRequests(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(500, gin.H{"error": "userid error"})
		return
	}

	var req []models.FollowRequest
	getReq := utils.PostgresDB.Model(&models.FollowRequest{}).Where("to_user_id = ? AND status = ?", userId, "pending").Find(&req)
	if getReq.Error != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	c.JSON(200, gin.H{
		"request": req,
	})
}

// unfollow api
func Unfollow(c *gin.Context) {
	userId := c.GetString("userId")
	targetId := c.Param("id")

	if userId == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if userId == targetId {
		c.JSON(400, gin.H{"error": "invalid user"})
		return
	}

	deleteFollow := utils.PostgresDB.Model(&models.Follow{}).Where("follower_id = ? AND following_id = ?", userId, targetId).Delete(&models.Follow{})
	if deleteFollow.Error != nil {
		c.JSON(500, gin.H{"error": "could not unfollow"})
		return
	}

	if deleteFollow.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "not following"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "unfollowed✅",
	})
}

// get my followers
func GetFollowers(c *gin.Context) {
	userId := c.GetString("userId")

	if userId == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var followers []models.Follow
	getFollowers := utils.PostgresDB.Model(&models.Follow{}).Where("following_id = ?", userId).Preload("Follower").Find(&followers)
	if getFollowers.Error != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	c.JSON(200, gin.H{
		"followers": followers,
	})
}

// get my following
func GetMyFollowing(c *gin.Context) {
	userId := c.GetString("userId")

	if userId == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var following []models.Follow
	getFollowing := utils.PostgresDB.Model(&models.Follow{}).Where("follower_id = ?", userId).Preload("Following").Find(&following)
	if getFollowing.Error != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	c.JSON(200, gin.H{
		"following": following,
	})
}
