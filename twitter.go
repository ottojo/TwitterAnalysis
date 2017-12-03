package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"strconv"
	"log"
)

func getTweets(tweetChannel chan anaconda.Tweet, user string) {
	v := url.Values{}
	v.Set("screen_name", user)
	//v.Set("trim_user", "true")
	v.Set("count", "200")
	for {
		newTweets, _ := api.GetUserTimeline(v)
		log.Printf("Got %d tweets from %s", len(newTweets), user)
		if len(newTweets) == 0 {
			break
		}
		for _, tweet := range newTweets {
			tweetChannel <- tweet
		}
		v.Set("max_id", strconv.FormatInt(newTweets[len(newTweets)-1].Id-1, 10))
	}
	close(tweetChannel)
}
