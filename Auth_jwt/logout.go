package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//DeleteAuth is deleting the AccessUUID from Redis 
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

//Logout --> ExtractTokenMetadata -> VerifyToken -> ExtractToken -> DeleteAuth
func Logout(c *gin.Context) {
	au, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
