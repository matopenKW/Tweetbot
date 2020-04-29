package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type AwsQuiz struct {
	SeqNo    string   `json: "SeqNo"`
	Quiz     string   `json: "Quiz"`
	Answer   string   `json: "Answer"`
	hashTags []string `json: "hashTags"`
}

func main() {
	http.HandleFunc("/tweet", Tweet)
	http.HandleFunc("/reply", Reply)
	http.HandleFunc("/getQuiz", func(w http.ResponseWriter, r *http.Request) {
		quiz, err := getQuiz()

		log.Println(quiz)

		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err)
		} else {
			fmt.Fprint(w, quiz)
		}
	})
	http.ListenAndServe(":80", nil)
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
	quiz, err := getQuiz()
	if err != nil {
		log.Println("err", err)
		return
	}

	//tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
	t, res, err := client.Statuses.Update(quiz.Quiz, nil)
	if err != nil {
		log.Println("err", err)
		return
	}
	// ツイート情報とhttpレスポンス
	log.Println("res", res)

	params := &twitter.StatusUpdateParams{
		InReplyToStatusID: t.ID,
	}

	_, res, err = client.Statuses.Update(quiz.Answer, params)
	if err != nil {
		log.Println("err", err)
		return
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
	_, res, err := client.Statuses.Update(txt, params)
	if err != nil {
		log.Println("err", err)
	}
	// ツイート情報とhttpレスポンス
	log.Println("tweet", res)
}

func getQuiz() (*AwsQuiz, error) {
	res, err := http.Get("http://localhost:8080/awsQuiz")
	defer res.Body.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(string(byteArray))

	var quiz AwsQuiz
	err = json.Unmarshal(byteArray, &quiz)
	log.Println(quiz)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &quiz, nil
}
