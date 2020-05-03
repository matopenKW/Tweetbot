package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/tweet", tweet)
	router.GET("/tweetView", tweetView)
	router.Run()
}

func tweet(ctx *gin.Context) {
	c, _ := ini.Load("config.conf")
	consumerKey := c.Section("twitterAPI").Key("consumerKey").String()
	consumerSecret := c.Section("twitterAPI").Key("consumerSecret").String()
	accessToken := c.Section("twitterAPI").Key("accessToken").String()
	accessTokenSecret := c.Section("twitterAPI").Key("accessTokenSecret").String()

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	quiz, err := getQuiz()
	if err != nil {
		log.Println("err", err)
		return
	}

	//tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
	t, res, err := client.Statuses.Update(getHeader()+quiz.Quiz, nil)
	if err != nil {
		log.Println("err", err)
		return
	}
	// ツイート情報とhttpレスポンス
	log.Println("res", res)

	params := &twitter.StatusUpdateParams{
		InReplyToStatusID: t.ID,
	}

	_, res, err = client.Statuses.Update("答え: "+quiz.Answer, params)
	if err != nil {
		log.Println("err", err)
		return
	}

	// ツイート情報とhttpレスポンス
	log.Println("tweet", res)
}

func tweetView(ctx *gin.Context) {
	quiz, err := getQuiz()
	if err != nil {
		ctx.String(http.StatusHTTPVersionNotSupported, "err")
	}

	quiz.Quiz = getHeader() + quiz.Quiz

	log.Println(quiz)

	if err != nil {
		ctx.String(http.StatusHTTPVersionNotSupported, "err")
	} else {
		ctx.JSON(http.StatusOK, quiz)
	}
}

func getHeader() string {
	return "[AWS 認定ソリューションアーキテクト1日1問]\r\n"
}

func getQuiz() (*AwsQuiz, error) {
	c, _ := ini.Load("config.conf")
	url := c.Section("AwsQuiz").Key("url").String() + "/awsQuiz"
	res, err := http.Get(url)
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
