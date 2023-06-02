package model

//動作確認用

/*type User struct {
	Id   string
	Name string
	Age  int
}*/

//以下ハッカソン用

type AppUser struct {
	UserId   string
	UserName string
	Email    string
	Password string
}

type Channel struct {
	ChannelId   string
	ChannelName string
}

type ChannelMember struct {
	ChannelMemberId string
	ChannelId       string
	UserId          string
}

type Message struct {
	MessageId      string
	UserId         string
	MessageContent string
}

type Reply struct {
	ReplyId      string
	MessageId    string
	UserId       string
	ReplyContent string
}
