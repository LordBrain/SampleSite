package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//Basic HTML page
	r.Static("/images", "./html/images")
	r.Static("/css", "./html/css")
	r.LoadHTMLGlob("html/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Main website", //IGNORE THIS
		})
	})

	r.GET("/big", func(c *gin.Context) {
		c.HTML(200, "big.html", gin.H{
			"title": "Main website", //IGNORE THIS
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
