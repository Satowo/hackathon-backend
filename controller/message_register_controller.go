package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
	"unicode/utf8"
)

type MessageResForHTTPPost struct {
	UserId         string `json:"user_id"`
	ChannelId      string `json:"channel_id"`
	MessageContent string `json:"message_content"`
}

type Message struct {
	MessageId      string
	UserId         string
	ChannelId      string
	MessageContent string
}

func MessageRegisterController(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータの文字列を取得、空文字の場合エラーコード400を返す
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		log.Println("fail: userId is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// リクエストのボディを読み込み
	var data MessageResForHTTPPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("fail: json.NewDecoder(r.Body).Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Nameが空文字か50字以上、Ageが81歳以上19歳以下の時エラーコード400を返す
	if data.MessageContent == "" || utf8.RuneCountInString(data.MessageContent) > 500 {
		log.Printf("fail: invalidinput")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	messageContent := data.MessageContent

	// データをデータベースに挿入しそれを含んだ全userのnameとageを返す
	err = usecase.UserRegisterUseCase(userId, channelId, messageContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
