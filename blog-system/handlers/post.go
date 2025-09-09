package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"study/blog-system/models"
)

type PostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreatePost 创建新文章
func CreatePost(c *gin.Context) {
	// 从上下文对象中获取用户 ID
	userID, _ := c.Get("user_id")
	// 检查用户 ID 是否有效
	if userID.(int) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// 绑定请求体参数
	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从上下文对象中获取数据库连接
	db := c.MustGet("db").(*gorm.DB)
	// 创建文章对象
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID.(uint),
	}
	// 保存文章到数据库
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	// 返回创建的文章
	c.JSON(http.StatusOK, post)
}

// GetPosts 获取所有文章
func GetPosts(c *gin.Context) {
	// 从上下文对象中获取数据库连接
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post
	// 查询所有文章
	if err := db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost 获取单篇文章
func GetPost(c *gin.Context) {
	// 从 URL 参数中获取文章 ID
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var post models.Post
	// 查询文章
	if err := db.First(&post, id).Error; err != nil {
		// 处理查询错误
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	// 从 URL 参数中获取文章 ID
	id := c.Param("id")
	// 从上下文对象中获取用户 ID
	userID, _ := c.Get("user_id")
	if userID.(int) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// 绑定请求体参数
	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 从上下文对象中获取数据库连接
	db := c.MustGet("db").(*gorm.DB)
	// 查询文章
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	// 检查用户是否有权限更新文章
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this post"})
		return
	}
	// 更新文章字段
	post.Title = input.Title
	post.Content = input.Content
	// 保存更新后的文章
	db.Save(&post)
	c.JSON(http.StatusOK, post)
}

// DeletePost 删除文章
func DeletePost(c *gin.Context) {
	// 从 URL 参数中获取文章 ID
	id := c.Param("id")
	// 从上下文对象中获取用户 ID
	userID, _ := c.Get("user_id")
	if userID.(int) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// 从上下文对象中获取数据库连接
	db := c.MustGet("db").(*gorm.DB)
	var post models.Post
	// 查询文章
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	// 检查用户是否有权限删除文章
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this post"})
		return
	}
	// 删除文章
	db.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
