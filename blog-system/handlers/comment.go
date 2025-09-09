package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"study/blog-system/models"
)

type CommentInput struct {
	Content string `json:"content" binding:"required"`
}

// CreateComment 创建新评论
func CreateComment(c *gin.Context) {
	// 从上下文对象中获取用户 ID
	userID, _ := c.Get("user_id")
	// 检查用户 ID 是否有效
	if userID.(int) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// 从 URL 参数中获取文章 ID
	postID := c.Param("id")
	// 绑定请求体参数
	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 从上下文对象中获取数据库连接
	db := c.MustGet("db").(*gorm.DB)
	// 创建评论对象
	comment := models.Comment{
		Content: input.Content,
		UserID:  userID.(uint),
		PostID:  uintFromString(postID),
	}
	// 保存评论到数据库
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetComments 获取文章的所有评论
func GetComments(c *gin.Context) {
	// 从 URL 参数中获取文章 ID
	postID := c.Param("id")
	// 从上下文对象中获取数据库连接
	db := c.MustGet("db").(*gorm.DB)
	var comments []models.Comment
	// 查询文章的所有评论
	if err := db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// 辅助函数：将字符串转为 uint
func uintFromString(s string) uint {
	var i uint
	fmt.Sscanf(s, "%d", &i)
	return i
}
