package model

type Channel struct {
	ChannelId   string
	ChannelName string
}

type ChannelMember struct {
	ChannelMemberId string
	ChannelId       string
	UserId          string
}
