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

	router.Run() // listen and serve on port 8080
}
