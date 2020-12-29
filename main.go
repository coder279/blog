package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"message",
		})
	})
	go http.ListenAndServe("127.0.0.1:8081",http.DefaultServeMux)
	r.Run()
}
