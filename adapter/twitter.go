package adapter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/taguchi-so/example-twitter-api/service"
)

type twitterUserImpl struct {
	client *twitter.Client
}

// NewTwitterUser new twitter user conetext service
func NewTwitterUser(client *twitter.Client) service.TwitterUser {
	return &twitterUserImpl{client}
}

func (s *twitterUserImpl) FindImage() {
	print("TwitterUserServie FindImage")
}

func (s *twitterUserImpl) SendReplayMessageFromTweetID() {
	print("TwitterUserServie SendReplayMessageFromTweetID")
}

func (s *twitterUserImpl) ReceiveReplay() {
	print("TwitterUserServie ReceiveReplay")
}
