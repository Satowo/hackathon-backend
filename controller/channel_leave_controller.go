package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type ChannelLeaveRequestForHTTPPost struct {
	UserId    string `json:"userId"`
	ChannelId string `json:"channelId"`
}

func ChannelLeaveController(w http.ResponseWriter, r *http.Request) {
	// リクエストのボディを読み込み
	var data ChannelJoinRequestForHTTPPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("fail: json.NewDecoder(r.Body).Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	channelId := data.ChannelId
	userId := data.UserId

	// データをデータベースに挿入しそれを含んだ全userのnameとageを返す
	userInfo, err := usecase.ChannelLeaveUseCase(channelId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のUserResForHTTPGetのリスト型に変換
	var userInfoRes UserInfoResForHTTPGet
	userInfoRes.UserId = userInfo.UserId
	userInfoRes.UserName = userInfo.UserName
	userInfoRes.Email = userInfo.Email

	userInfoRes.Channels = make([]Channel, 0)
	for _, u := range userInfo.Channels {
		userInfoRes.Channels = append(userInfoRes.Channels, Channel{
			ChannelId:   u.ChannelId,
			ChannelName: u.ChannelName,
		})
	}

	bytes, err := json.Marshal(userInfoRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
