package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"log"
	"net/http"
	"os"
)

type Quiz struct {
	No     string `csv:"No"`
	Quiz   string `csv:"Quiz"`
	Ansewr string `csv:"Ansewr"`
}

func main() {
	router := gin.Default()
	router.GET("/awsQuiz", awsQuiz)

	router.Run()
}

func awsQuiz(ctx *gin.Context) {

	quizList, err := getCsvList()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusHTTPVersionNotSupported, "error")
		return
	}

	num := 0

	var quiz *Quiz
	for i, v := range quizList {
		if i == num {
			quiz = v
		}
	}
	log.Println(quiz)

	ctx.JSON(http.StatusOK, &quiz)
}

func getCsvList() ([]*Quiz, error) {
	file, err := os.Open("csv/aws.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	quizList := []*Quiz{}
	err = gocsv.UnmarshalFile(file, &quizList)
	if err != nil {
		return nil, err
	}

	log.Println(quizList)

	return quizList, nil
}
