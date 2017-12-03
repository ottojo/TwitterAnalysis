package main

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
	"log"
	"fmt"
	"time"
	"math"
)

var api *anaconda.TwitterApi

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify an action.")
		//TODO Display help
	}

	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)

	api = anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
	api.EnableThrottling(1*time.Second, math.MaxInt64)
	api.ReturnRateLimitError(true)

	switch os.Args[1] {
	case "toptweets":
		topTweets(os.Args[2:])
		break
	case "interaction":
		interaction(os.Args[2:])
		break
	default:
		fmt.Println("This is not a valid action.")
	}
}
