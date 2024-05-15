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

	// Api ini berguna untuk kebutuhan utama bisnis.
	r.GET("/posts", middlewares.AuthMiddleWare, handlers.GetPostsHandler)
	r.GET("/posts/:id", middlewares.AuthMiddleWare, handlers.GetPostHandler)
	r.POST("/posts", middlewares.AuthMiddleWare, handlers.CreatePostHandler)
	r.PUT("/posts/:id", middlewares.AuthMiddleWare, handlers.UpdatePostHandler)
	r.DELETE("/posts/:id", middlewares.AuthMiddleWare, handlers.DeletePostHandler)

	// Api ini berguna untuk melakukan authentikasi dan deauthentikasi.
	r.POST("/signup", handlers.SignupHandler)
	r.POST("/login", handlers.LoginHandler)
	r.POST("/logout", middlewares.AuthMiddleWare, handlers.LogOutHandler)

	// Api ini berguna untuk CMS.
	r.GET("/users/:id", middlewares.AuthMiddleWare, handlers.GetUserHandler)
	r.PUT("/users/:id", middlewares.AuthMiddleWare, handlers.UpdateUserHandler)
	r.DELETE("/users/:id", middlewares.AuthMiddleWare, handlers.DeleteUserHandler)

	r.Run(":8080")
}
