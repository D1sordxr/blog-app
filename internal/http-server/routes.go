package http_server

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/posts", GetPosts)
		api.POST("/posts", AddPost)
		api.GET("/posts/:id", GetPostByID)
		api.PATCH("/posts/:id", UpdatePost)
		api.DELETE("/posts/:id", DeletePost)
	}
}
