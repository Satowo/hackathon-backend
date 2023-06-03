package model

type AppUser struct {
	UserId   string
	UserName string
	Email    string
	Password string
}

type Reply struct {
	ReplyId      string
	MessageId    string
	UserId       string
	ReplyContent string
}
