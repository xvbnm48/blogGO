package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/blogGO/handler"
	"github.com/xvbnm48/blogGO/post"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/blogGo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection database error, error: ", err)
	}

	fmt.Println("connection database success")
	postRepository := post.NewRepository(db)
	postService := post.NewService(postRepository)
	postHandler := handler.NewPostHandler(postService)
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	v1.POST("/post", postHandler.CreatePost)
	v1.GET("/post", postHandler.FindAllPost)
	v1.PUT("/post/:id", postHandler.UpdatePost)
	v1.DELETE("/post/:id", postHandler.DeletePost)

	router.Run(":8080")
}
