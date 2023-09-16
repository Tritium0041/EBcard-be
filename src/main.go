package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//ginserver
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.Run() // listen and serve on
	r.POST("/upload")
}

func UploadBusinessCard(c *gin.Context) {
	db := getdb()
	defer db.Close()
	
}
