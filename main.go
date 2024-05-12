package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{})

	r := gin.Default()

	r.GET("/posts", GetPostsHandler(db))
	r.GET("/posts/:id", GetPostHandler(db))
	r.POST("/posts", CreatePostHandler(db))
	r.PUT("/posts/:id", UpdatePostHandler(db))
	r.DELETE("/posts/:id", DeletePostHandler(db))

	r.Run(":8080")
}
