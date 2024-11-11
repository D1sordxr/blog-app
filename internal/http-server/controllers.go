package http_server

import (
	"BlogWebApp/internal/storage/models"
	db "BlogWebApp/internal/storage/postgres"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := db.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
	return
}

func GetPostByID(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	if err := db.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func AddPost(c *gin.Context) {
	var json map[string]string
	if err := c.BindJSON(&json); err != nil {
		log.Printf("Error parsing json")
	}

	post := models.Post{
		Title:   json["title"],
		Content: json["content"],
	}

	db.DB.Create(&post)
	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created!"})
	return
}

func UpdatePost(c *gin.Context) {

}

func DeletePost(c *gin.Context) { // TODO: remake delete method
	var post models.Post

	id := c.Param("id")

	if err := db.DB.Delete(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
