package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type ChannelMemberResForHTTPGet struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ChannelMemberSearchController(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータの文字列を取得、空文字の場合エラーコード400を返す
	channelId := r.URL.Query().Get("channel_id")
	if channelId == "" {
		log.Println("fail: channel_id is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//User構造体のリスト型のデータを受け取り、
	members, err := usecase.ChannelMemberSearchUseCase(channelId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のUserResForHTTPGetのリスト型に変換
	membersRes := make([]ChannelMemberResForHTTPGet, 0)
	for _, u := range members {
		membersRes = append(membersRes, ChannelMemberResForHTTPGet{
			UserId:   u.UserId,
			UserName: u.UserName,
			Email:    u.Email,
			Password: u.Password,
		})
	}

	// 変換したユーザーデータリストをjson形式に変換
	bytes, err := json.Marshal(membersRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
