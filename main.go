package main

import (
	"gopkg.in/ini.v1"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Quiz struct {
	no       int
	quiz     string
	answer   string
	hashTags []string
}

func main() {

	http.HandleFunc("/tweet", Tweet)
	http.HandleFunc("/reply", Reply)
	http.ListenAndServe(":8080", nil)
}

func Tweet(w http.ResponseWriter, r *http.Request) {
	c, _ := ini.Load("config.conf")
	consumerKey := c.Section("twitterAPI").Key("consumerKey").String()
	consumerSecret := c.Section("twitterAPI").Key("consumerSecret").String()
	accessToken := c.Section("twitterAPI").Key("accessToken").String()
	accessTokenSecret := c.Section("twitterAPI").Key("accessTokenSecret").String()

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	//	txt := "顧客は、ある Web アプリケーションを使用して、Amazon S3 バケットに注文データをアッ プロードすることができます。すると、Amazon S3 イベントが発生し、Lambda 関数がトリ ガされ、メッセージが SQS キューに挿入されます。1 つの EC2 インスタンスによって、キュ ーからメッセージが読み取られて処理され、一意の注文番号で分割された DynamoDB テーブ ルに格納されます。来月のトラフィック量は 10 倍に増える見込みです。ソリューションアー キテクトは、スケーリングに関する問題がアーキテクチャに発生する可能性を調べています。 増加するトラフィックを処理するためにスケーリングできるようにする際、設計の見直しが最 も必要であると思われるコンポーネントはどれですか。"
	// txt := "Twitter Botを作ろう"
	quiz := getQuiz()

	//tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
	t, res, e := client.Statuses.Update(quiz.quiz, nil)
	if e != nil {
		log.Println("err", e)
	}
	// ツイート情報とhttpレスポンス
	log.Println("res", res)

	params := &twitter.StatusUpdateParams{
		InReplyToStatusID: t.ID,
	}

	_, res, e = client.Statuses.Update(quiz.answer, params)
	if e != nil {
		log.Println("err", e)
	}

	// ツイート情報とhttpレスポンス
	log.Println("tweet", res)
}

func Reply(w http.ResponseWriter, r *http.Request) {
	c, _ := ini.Load("config.conf")
	consumerKey := c.Section("twitterAPI").Key("consumerKey").String()
	consumerSecret := c.Section("twitterAPI").Key("consumerSecret").String()
	accessToken := c.Section("twitterAPI").Key("accessToken").String()
	accessTokenSecret := c.Section("twitterAPI").Key("accessTokenSecret").String()

	params := &twitter.StatusUpdateParams{
		InReplyToStatusID: 1249712564776230913,
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	txt := "test reply"
	_, res, e := client.Statuses.Update(txt, params)
	if e != nil {
		log.Println("err", e)
	}
	// ツイート情報とhttpレスポンス
	log.Println("tweet", res)
}

func getQuiz() *Quiz {

	// 何らかの方法でTweetを取得する
	quizStr :=
		`AWSQius 一日一問
	EC2インスタンスに付与できるタグの上限は次のうちどれですか？
	ア. 10
	イ. 20
	ウ. 40
	エ. 50
	`

	quiz := &Quiz{
		no:     1,
		quiz:   quizStr,
		answer: "[答え] エ",
	}

	return quiz
}
