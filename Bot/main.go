package main

import (
	"appengine/urlfetch"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/http"
	"net/url"

	"appengine"
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

	http.HandleFunc("/tweet", Tweet)

}

func Tweet(w http.ResponseWriter, r *http.Request) {

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

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

func getTweet() string {

	// 何らかの方法でTweetを取得する

	return "hello anaconda"
}
