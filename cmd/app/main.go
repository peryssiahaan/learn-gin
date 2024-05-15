package main

import (
	"github.com/gin-gonic/gin"

	"gin-blog-app/database"
	"gin-blog-app/handlers"
	"gin-blog-app/middlewares"
)

func main() {

	if err := database.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/posts", middlewares.AuthMiddleWare, handlers.GetPostsHandler)
	r.GET("/posts/:id", middlewares.AuthMiddleWare, handlers.GetPostHandler)
	r.POST("/posts", middlewares.AuthMiddleWare, handlers.CreatePostHandler)
	r.PUT("/posts/:id", middlewares.AuthMiddleWare, handlers.UpdatePostHandler)
	r.DELETE("/posts/:id", middlewares.AuthMiddleWare, handlers.DeletePostHandler)

	r.POST("/signup", handlers.SignupHandler)
	r.POST("/login", handlers.LoginHandler)
	r.POST("/logout", middlewares.AuthMiddleWare, handlers.LogOutHandler)
	r.GET("/users/:id", middlewares.AuthMiddleWare, handlers.GetUserHandler)
	r.PUT("/users/:id", middlewares.AuthMiddleWare, handlers.UpdateUserHandler)
	r.DELETE("/users/:id", middlewares.AuthMiddleWare, handlers.DeleteUserHandler)

	r.Run(":8080")
}
