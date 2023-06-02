package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type ChannelResForHTTPGet struct {
	ChannelId   string `json:"id"`
	ChannelName string `json:"name"`
}

func ChannelSearchController(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(bytes)
}
