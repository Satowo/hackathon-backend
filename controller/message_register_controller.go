package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
	"unicode/utf8"
)

type MessageRequestForHTTPPost struct {
	UserId         string `json:"userId"`
	ChannelId      string `json:"channelId"`
	MessageContent string `json:"messageContent"`
}

type MessageResForHTTPPost struct {
	MessageId      string `json:"messageId"`
	UserId         string `json:"userId"`
	UserName       string `json:"userName"`
	ChannelId      string `json:"channelId"`
	MessageContent string `json:"messageContent"`
	Edited         bool   `json:"edited"`
}

func MessageRegisterController(w http.ResponseWriter, r *http.Request) {
	// リクエストのボディを読み込み
	var data MessageRequestForHTTPPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("fail: json.NewDecoder(r.Body).Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//MessageContentが空文字か500字以上の時エラーコード400を返す
	if data.MessageContent == "" || utf8.RuneCountInString(data.MessageContent) > 500 {
		log.Printf("fail: invalidinput")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	userId := data.UserId
	channelId := data.ChannelId
	messageContent := data.MessageContent

	// データをデータベースに挿入しそれを含んだチャンネル内のmessageContent,userName,editedを返す。
	messages, err := usecase.MessageRegisterUseCase(userId, channelId, messageContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のMessageResForHTTPPostのリスト型に変換
	messagesRes := make([]MessageResForHTTPPost, 0)
	for _, u := range messages {
		messagesRes = append(messagesRes, MessageResForHTTPPost{
			MessageId:      u.UserId,
			UserId:         u.UserId,
			UserName:       u.UserName,
			ChannelId:      u.ChannelId,
			MessageContent: u.MessageContent,
			Edited:         u.Edited,
		})
	}

	bytes, err := json.Marshal(messagesRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
