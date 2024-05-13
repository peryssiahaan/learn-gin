package main

import (
	"github.com/gin-gonic/gin"

	"gin-blog-app/database"
	"gin-blog-app/handlers"
)

func main() {

	if err := database.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/posts", handlers.GetPostsHandler)
	r.GET("/posts/:id", handlers.GetPostHandler)
	r.POST("/posts", handlers.CreatePostHandler)
	r.PUT("/posts/:id", handlers.UpdatePostHandler)
	r.DELETE("/posts/:id", handlers.DeletePostHandler)

	r.POST("/signup", handlers.SignupHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/users/:id", handlers.GetUserHandler)
	r.PUT("/users/:id", handlers.UpdateUserHandler)
	r.DELETE("/users/:id", handlers.DeleteUserHandler)

	r.Run(":8080")
}
