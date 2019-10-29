package infra

import (
	"errors"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const twiterTokenEndPoint = "https://api.twitter.com/oauth2/token"

// NewTwitterConf return Client Credentials Config
func NewTwitterConf(consumerKey, consumerSecret string) (*clientcredentials.Config, error) {
	if consumerKey == "" || consumerSecret == "" {
		return nil, errors.New("Invalid ConsumerKey or ConsumerSecret")
	}
	return &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     twiterTokenEndPoint,
	}, nil
}

// NewTwitterClientWithNoContext return twitter client
func NewTwitterClientWithNoContext(config *clientcredentials.Config) *twitter.Client {
	httpClient := config.Client(oauth2.NoContext)
	return twitter.NewClient(httpClient)
}

// NewTwitterClientWithUserContext return twitter client
func NewTwitterClientWithUserContext(config *clientcredentials.Config) *twitter.Client {
	httpClient := config.Client(oauth2.NoContext)
	return twitter.NewClient(httpClient)
}
