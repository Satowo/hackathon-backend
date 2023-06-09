package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type MessageDeleteRequestForHTTPPost struct {
	MessageId string `json:"messageId"`
	ChannelId string `json:"channelId"`
}

func MessageDeleteController(w http.ResponseWriter, r *http.Request) {
	// リクエストのボディを読み込み
	var data MessageDeleteRequestForHTTPPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("fail: json.NewDecoder(r.Body).Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	messageId := data.MessageId
	channelId := data.ChannelId

	// データベースのmessageContentとeditedを更新しそれを含んだチャンネル内の全メッセージの情報を返す。
	messages, err := usecase.MessageDeleteUseCase(messageId, channelId)
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
	w.Write(bytes)
}
