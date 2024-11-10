package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts = make([]Post, 0)

func GetPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, posts)
}

func AddPost(ctx *gin.Context) {
	var newPost Post
	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newPost.ID = func() string {
		return fmt.Sprintf("%d", len(posts)+1)
	}()

	posts = append(posts, newPost)
	ctx.JSON(http.StatusCreated, newPost)

}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedPost Post

	if err := ctx.ShouldBindJSON(&updatedPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts[i].Title = updatedPost.Title
			posts[i].Content = updatedPost.Content
			ctx.JSON(http.StatusOK, posts[i])
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var result []Post // Создаем новый срез для хранения оставшихся постов

	for _, post := range posts {
		if post.ID != id {
			result = append(result, post) // Добавляем пост в результат, если его ID не совпадает
		}
	}

	posts = result // Обновляем глобальный срез posts
	ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
