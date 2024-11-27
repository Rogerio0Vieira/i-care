package main

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Birth  `json:"birth"`
}

func main(){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
