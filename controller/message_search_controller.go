package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type MessageResForHTTPGet struct {
	MessageId      string `json:"message_id"`
	UserName       string `json:"user_name"`
	ChannelId      string `json:"channel_id"`
	MessageContent string `json:"message_content"`
	Edited         bool   `json:"edited"`
}

func MessageSearchController(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータの文字列を取得、空文字の場合エラーコード400を返す
	channelId := r.URL.Query().Get("channelId")
	if channelId == "" {
		log.Println("fail: channel_id is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, err := usecase.MessageSearchUseCase(channelId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のUserResForHTTPGetのリスト型に変換
	messagesRes := make([]MessageResForHTTPGet, 0)
	for _, u := range messages {
		messagesRes = append(messagesRes, MessageResForHTTPGet{
			MessageId:      u.MessageId,
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
