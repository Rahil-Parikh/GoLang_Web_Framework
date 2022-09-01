package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//User is ...
	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Creates a gin router with default middleware
	var router *gin.Engine = gin.Default()
	// router := gin.Default()
	// A handler for GET request on /example
	router.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		}) // gin.H is a shortcut for map[string]interface{}

	})

	router.GET("/get", func(c *gin.Context) {
		message, _ := c.GetQuery("m")
		c.String(http.StatusOK, "Get works! you sent: "+message)
	})

	router.GET("/getPara/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := "Message -> " + name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/post", func(c *gin.Context) {
		var u User
		c.BindJSON(&u)
		c.JSON(http.StatusOK, gin.H{
			"user": u.Username,
			"pass": u.Password,
		})
	})
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.OPTIONS("/options", func(c *gin.Context) {
		c.Header("Allow", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"/ping":       "GET",
			"/get?m=....": "GET",
			"/post":       "POST",
		})
		c.Status(http.StatusOK)
	})

	router.Run() // listen and serve on port 8080
}

// router.Run(":3000")// router.Run(":3000") for a hard coded port
