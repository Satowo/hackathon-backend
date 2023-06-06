package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
	"unicode/utf8"
)

type MessageEditRequestForHTTPPost struct {
	MessageId      string `json:"messageId"`
	ChannelId      string `json:"channelId"`
	MessageContent string `json:"messageContent"`
}

func MessageEditController(w http.ResponseWriter, r *http.Request) {
	// リクエストのボディを読み込み
	var data MessageEditRequestForHTTPPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("fail: json.NewDecoder(r.Body).Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//MessageContentが空文字か500字以上の時エラーコード400を返す
	if data.MessageContent == "" || utf8.RuneCountInString(data.MessageContent) > 500 {
		log.Printf("fail: invalIdInput")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	messageId := data.MessageId
	messageContent := data.MessageContent
	channelId := data.ChannelId

	// データベースのmessageContentとeditedを更新しそれを含んだチャンネル内の全メッセージの情報を返す。
	messages, err := usecase.MessageEditUseCase(messageId, channelId, messageContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のMessageResForHTTPPostのリスト型に変換
	messagesRes := make([]MessageResForHTTPPost, 0)
	for _, u := range messages {
		messagesRes = append(messagesRes, MessageResForHTTPPost{
			MessageId:      u.MessageId,
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
