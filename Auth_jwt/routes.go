package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

var client *redis.Client

//User is ...
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}
var (
	router = gin.Default()
)

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	w, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	print("Reply to PING From REDIS_DSN : localhost:6379", w)
}

//TokenDetails is ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

func main() {
	router.POST("/login", Login)
	router.POST("/todo", CreateTodo)
	router.POST("/logout", Logout)
	router.POST("/token/refresh", Refresh)
	log.Fatal(router.Run(":8080"))
}
