package model

type Message struct {
	MessageId      string
	UserId         string
	UserName       string
	ChannelId      string
	MessageContent string
	Edited         bool
}
