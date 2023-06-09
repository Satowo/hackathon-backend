package model

type AppUser struct {
	UserId   string
	UserName string
	Email    string
	Password string
}

type UserInfo struct {
	UserId   string
	UserName string
	Email    string
	Channels []Channel
}

type Reply struct {
	ReplyId      string
	MessageId    string
	UserId       string
	ReplyContent string
}
