package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	consumerKey = "******************************"

	consumerSecret = "******************************"

	accessToken = "******************************"

	accessTokenSecret = "******************************"
)

type Quiz struct {
	No       int
	Quiz     string
	Ansewr   string
	HashTags []string
}

func main() {
	router := gin.Default()
	router.GET("/awsQuiz", awsQuiz)

	router.Run()
}

func awsQuiz(ctx *gin.Context) {

	quiz := &Quiz{
		1,
		`顧客関係管理 (CRM) アプリケーションが、Application Load Balancer の背後にある複数のア
		ベイラビリティーゾーン内の Amazon EC2 インスタンスで実行されています。
		これらのインスタンスのいずれかに障害が発生した場合、どうなりますか。
		A. ロードバランサーが、障害が発生したインスタンスへのリクエスト送信を停止する。
		B. ロードバランサーが、障害が発生したインスタンスを終了する。
		C. ロードバランサーが、障害が発生したインスタンスを自動的に置換する。
		D. ロードバランサーが、インスタンスが置換されるまで 504 ゲートウェイタイムアウトエ
		ラーを返す。`,
		"",
		[]string{"#aws, "},
	}
	log.Println(quiz)

	ctx.JSON(http.StatusOK, quiz)
}
