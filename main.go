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
	Choice   string   `json: "Choice"`
	Answer   string   `json: "Answer"`
	hashTags []string `json: "hashTags"`
}

func (quiz *AwsQuiz) cnvQuizList() []string {
	limitCount := 150
	ret := make([]string, 0, 0)

	runes := []rune(quiz.Quiz)
	for i := 0; i < len(runes); i += limitCount {
		if i+limitCount < len(runes) {
			ret = append(ret, string(runes[i:(i+limitCount)]))
		} else {
			ret = append(ret, string(runes[i:]))
		}
	}
	return ret
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
	params := &twitter.StatusUpdateParams{
		InReplyToStatusID: t.ID,
	}

	t, res, err = client.Statuses.Update(quiz.Choice, params)
	if err != nil {
		log.Println("err", err)
		return
	}

	params = &twitter.StatusUpdateParams{
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
		log.Println(err)
		ctx.String(http.StatusHTTPVersionNotSupported, "error")
		return
	}

	str := getHeader() + quiz.Quiz

	str = str + "\r\n" + quiz.Choice

	if err != nil {
		log.Println(err)
		ctx.String(http.StatusHTTPVersionNotSupported, "err")
	} else {
		ctx.String(http.StatusOK, str)
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
		return nil, err
	}

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(byteArray))

	var quiz AwsQuiz
	err = json.Unmarshal(byteArray, &quiz)

	if err != nil {
		return nil, err
	}

	return &quiz, nil
}
