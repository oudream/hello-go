package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	//router.LoadHTMLGlob("templates/*.tmpl.html")

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=UTF-8")
		c.File("2.html")
	})

	router.POST("/markdown", func(c *gin.Context) {
		body := c.PostForm("body")
		log.Println(body)
	})

	router.Run(":5000")
}