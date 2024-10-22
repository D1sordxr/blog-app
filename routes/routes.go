package routes

import "github.com/gin-gonic/gin"

func Setup(app *gin.RouterGroup) {
	app.GET("/api/posts")  // read all posts
	app.POST("/api/posts") // add post
	// TODO: read one post
	// TODO: delete post
}
