package main

import (
	"github.com/gin-gonic/gin"

	"gin-blog-app/database"
	"gin-blog-app/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := database.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/posts", handlers.GetPostsHandler)
	r.GET("/posts/:id", handlers.GetPostHandler)
	r.POST("/posts", handlers.CreatePostHandler)
	r.PUT("/posts/:id", handlers.UpdatePostHandler)
	r.DELETE("/posts/:id", handlers.DeletePostHandler)

	r.Run(":8080")
}
