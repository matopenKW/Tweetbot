package main

import (
	"GetQuiz/domain/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/awsQuiz", handler.AwsQuiz)

	router.Run()
}
