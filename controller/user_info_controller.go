package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type UserInfoResForHTTPGet struct {
	UserId     string   `json:"userId"`
	UserName   string   `json:"userName"`
	Email      string   `json:"email"`
	InChannels []string `json:"inChannels"`
}

func UserInfoController(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータの文字列を取得、空文字の場合エラーコード400を返す
	email := r.URL.Query().Get("email")
	if email == "" {
		log.Println("fail: email is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo, err := usecase.UserInfoUseCase(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のUserResForHTTPGetのリスト型に変換
	var userInfoRes UserInfoResForHTTPGet
	userInfoRes.UserId = userInfo.UserId
	userInfoRes.UserName = userInfo.UserName
	userInfoRes.Email = userInfo.Email
	userInfoRes.InChannels = userInfo.InChannels

	bytes, err := json.Marshal(userInfoRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
