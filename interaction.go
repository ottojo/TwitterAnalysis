package main

import (
	"github.com/ChimeraCoder/anaconda"
	"fmt"
	"sort"
	"net/url"
	"strconv"
	"log"
)

func interaction(args []string) {
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

	tweetChan := make(chan anaconda.Tweet)
	go getTweets(tweetChan, args[0])
	scores := interactionScore(tweetChan, args[0])
	if n == 0 {
		n = len(scores)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d: [%d] %s\n", i+1, scores[i].Score, scores[i].User)
	}
}

func interactionScore(tweetChannel chan anaconda.Tweet, username string) InteractionScoreList {
	scoreMap := make(map[string]int)
	for tweet := range tweetChannel {
		if tweet.QuotedStatus != nil {
			scoreMap[tweet.QuotedStatus.User.ScreenName]++
		} else if tweet.RetweetedStatus != nil && tweet.RetweetedStatus.User.ScreenName != username {
			scoreMap[tweet.RetweetedStatus.User.ScreenName]++
		}
		if tweet.InReplyToScreenName != "" && tweet.InReplyToScreenName != username {
			scoreMap[tweet.InReplyToScreenName]++
		}
	}

	var likedTweets []anaconda.Tweet

	v := url.Values{}
	v.Set("screen_name", username)
	v.Set("count", "200")
	v.Set("include_entities", "false")

	for {
		newLikes, _ := api.GetFavorites(v)
		log.Printf("Got %d tweets liked by %s", len(newLikes), username)
		if len(newLikes) == 0 {
			break
		}
		likedTweets = append(likedTweets, newLikes...)
		v.Set("max_id", strconv.FormatInt(newLikes[len(newLikes)-1].Id-1, 10))
	}

	for _, t := range likedTweets {
		if t.User.ScreenName != username {
			scoreMap[t.User.ScreenName]++
		}
	}

	scoreList := make(InteractionScoreList, len(scoreMap))
	i := 0
	for user, score := range scoreMap {
		scoreList[i] = InteractionScore{User: user, Score: score}
		i++
	}

	sort.Sort(sort.Reverse(scoreList))

	return scoreList
}

type InteractionScore struct {
	User  string
	Score int
}

type InteractionScoreList []InteractionScore

func (l InteractionScoreList) Len() int           { return len(l) }
func (l InteractionScoreList) Less(i, j int) bool { return l[i].Score < l[j].Score }
func (l InteractionScoreList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
