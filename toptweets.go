package main

import (
	"github.com/ChimeraCoder/anaconda"
	"strconv"
	"fmt"
	"log"
)

func topTweets(args []string) {
	if len(args) < 1 {
		log.Fatal("Please provide a username")
	}
	n := 5
	if len(args) > 1 {
		i, err := strconv.Atoi(args[1])
		if err == nil {
			n = i
		}
	}
	printBestTweets(n, args[0])
}

func printBestTweets(n int, user string) {
	tweetChannel := make(chan anaconda.Tweet, 200)
	go getTweets(tweetChannel, user)
	sortedTweets := sortTweets(tweetChannel)
	if n == 0 {
		n = len(sortedTweets)
	}
	for i := len(sortedTweets) - 1; i >= len(sortedTweets)-n; i-- {
		fmt.Printf("%d: [%d] %s https://twitter.com/%s/status/%d\r\n", len(sortedTweets)-i, sortedTweets[i].FavoriteCount, sortedTweets[i].Text, user, sortedTweets[i].Id)
	}
}

func sortTweets(tweetChannel chan anaconda.Tweet) []anaconda.Tweet {
	sorted := []anaconda.Tweet{<-tweetChannel}
	for input := range tweetChannel {
		sorted = append(sorted, anaconda.Tweet{})
		for i := range sorted {
			if sorted[i].FavoriteCount > input.FavoriteCount {
				copy(sorted[i+1:], sorted[i:])
				sorted[i] = input
				break
			}
			if i == len(sorted)-1 {
				sorted[i] = input
			}
		}
	}
	return sorted
}
