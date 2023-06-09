package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type ChannelResForHTTPGet struct {
	ChannelId   string `json:"channelId"`
	ChannelName string `json:"channelName"`
}

func ChannelSearchController(w http.ResponseWriter) {
	channels, err := usecase.ChannelSearchUseCase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のUserResForHTTPGetのリスト型に変換
	channelsRes := make([]ChannelResForHTTPGet, 0)
	for _, u := range channels {
		channelsRes = append(channelsRes, ChannelResForHTTPGet{
			ChannelId:   u.ChannelId,
			ChannelName: u.ChannelName,
		})
	}

	bytes, err := json.Marshal(channelsRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
