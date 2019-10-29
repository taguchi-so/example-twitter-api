package usecase

import "github.com/taguchi-so/example-twitter-api/service"

// A A Usecase
type A interface {
	// FindImage
	FindImage()

	// SendRequest
	SendRequest()

	// ReceiveReplay
	ReceiveReplay()
}

type useCaseAImpl struct {
	twitterUserService service.TwitterUser
}

// NewA constructor
func NewA(twitterUserService service.TwitterUser) A {
	return &useCaseAImpl{twitterUserService}
}

func (u *useCaseAImpl) FindImage() {
	u.twitterUserService.FindImage()
}

func (u *useCaseAImpl) SendRequest() {
	u.twitterUserService.SendReplayMessageFromTweetID()
}

func (u *useCaseAImpl) ReceiveReplay() {
	u.twitterUserService.ReceiveReplay()
}
