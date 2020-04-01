package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

const (
	consumerKey = "******************************"

	consumerSecret = "******************************"

	accessToken = "******************************"

	accessTokenSecret = "******************************"
)

type Quiz struct {
	no       int
	quiz     string
	ansewr   string
	hashTags []string
}

func main() {
	router := gin.Default()
	router.POST("/awsQuiz", awsQuiz)

	router.Run()
}

func awsQuiz(ctx *gin.Context) {

	quiz := &Quiz{
		1,
		"",
		"",
		[]string{""},
	}

	bytes := 

	return ctx.JSON(http.StatusOK, quiz)
}
