package service

// TwitterUser twitter user context servcice interface
type TwitterUser interface {
	FindImage()
	SendReplayMessageFromTweetID()
	ReceiveReplay()
}
