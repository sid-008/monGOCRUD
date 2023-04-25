package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sid-008/monGoCRUD/routes"
)

func main() {
	router := gin.Default()

	router.POST("/", routes.CreatePost)

	router.GET("/getOne/:postId", routes.ReadOnePost)

	router.PUT("/update/:postId", routes.UpdatePost)

	router.DELETE("/delete/:postId", routes.DeletePost)

	err := router.Run("localhost: 3000")
	if err != nil {
		log.Fatal(err)
	}
}
