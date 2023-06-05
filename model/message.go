package model

type Message struct {
	MessageId      string
	UserId         string
	ChannelId      string
	MessageContent string
	Edited         bool
}
