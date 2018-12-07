package service

import (
	"fmt"

	"github.com/rdip/twitter/src/domain"
)

type TweetManager struct {
	tweetsMap map[string][]*domain.TextTweet
	tweets    []*domain.TextTweet
}

func NewTweetManager() *TweetManager {
	var tweetsMap = map[string][]*domain.TextTweet{}
	var tweets = []*domain.TextTweet{}

	var manager = TweetManager{tweetsMap, tweets}

	return &manager
}

func (manager *TweetManager) PublishTweet(tweet1 *domain.TextTweet) (int, error) {
	if tweet1.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	}
	if tweet1.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(tweet1.GetText()) > 140 {
		return 0, fmt.Errorf("tweet must have no more than 140 characters")
	}

	max_id := 0

	for _, tweet := range manager.tweets {
		if tweet.GetId() > max_id {
			max_id = tweet.GetId()
		}
	}

	max_id = max_id + 1

	tweet1.Id = max_id

	manager.tweets = append(manager.tweets, tweet1)

	manager.tweetsMap[tweet1.User] = append(manager.tweetsMap[tweet1.User], tweet1)

	return max_id, nil
}

func (manager *TweetManager) GetTweets() []*domain.TextTweet {
	return manager.tweets
}

func (manager *TweetManager) GetTweetById(id int) *domain.TextTweet {
	for _, tweet := range manager.tweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func (manager *TweetManager) CountTweetsByUser(user string) int {
	count := 0
	for _, tweet := range manager.tweets {
		if tweet.User == user {
			count = count + 1
		}
	}
	return count
}

func (manager *TweetManager) GetTweetsByUser(user string) []*domain.TextTweet {
	return manager.tweetsMap[user]
}

func 