package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"net/url"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

type Quiz struct {
	no       int
	quiz     string
	ansewr   string
	hashTags []string
}

func main() {

	http.HandleFunc("/tweet", Tweet)
	http.HandleFunc("/tweet2", Tweet2)
	http.ListenAndServe(":8080", nil)
}

func Tweet(w http.ResponseWriter, r *http.Request) {

	c, _ := ini.Load("config.conf")
	consumerKey := c.Section("twitterAPI").Key("consumerKey").String()
	consumerSecret := c.Section("twitterAPI").Key("consumerSecret").String()
	accessToken := c.Section("twitterAPI").Key("accessToken").String()
	accessTokenSecret := c.Section("twitterAPI").Key("accessTokenSecret").String()

	// ConsumerKeyのセット
	anaconda.SetConsumerKey(consumerKey)

	// ConsumerSecretのセット
	anaconda.SetConsumerSecret(consumerSecret)

	// AccessTokenとAccessTokenSecretのセット
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	ctx := appengine.NewContext(r)

	// apiのHttpClient.TransportにurlfetchのTransportを利用する
	api.HttpClient.Transport = &urlfetch.Transport{Context: ctx}
	v := url.Values{}

	// ツイートする

	_, err := api.PostTweet(getTweet(), v)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, "Tweet!")
	}
}

func Tweet2(w http.ResponseWriter, r *http.Request) {
	c, _ := ini.Load("config.conf")
	consumerKey := c.Section("twitterAPI").Key("consumerKey").String()
	consumerSecret := c.Section("twitterAPI").Key("consumerSecret").String()
	accessToken := c.Section("twitterAPI").Key("accessToken").String()
	accessTokenSecret := c.Section("twitterAPI").Key("accessTokenSecret").String()

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	txt := "test tweet"

	//tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
	_, res, e := client.Statuses.Update(txt, nil)
	if e != nil {
		log.Println("err", e)
	}
	// ツイート情報とhttpレスポンス
	log.Println("tweet", res)

}

func getTweet() string {

	// 何らかの方法でTweetを取得する

	return "test anaconda"
}
