package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
	"unicode/utf8"
)

type UserRequestForHTTPPost struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegisterController(w http.ResponseWriter, r *http.Request) {
	// リクエストのボディを読み込み
	var data UserRequestForHTTPPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("fail: json.NewDecoder(r.Body).Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Nameが空文字か50字以上、Ageが81歳以上19歳以下の時エラーコード400を返す
	if data.UserName == "" || utf8.RuneCountInString(data.UserName) > 10 ||
		data.Email == "" || utf8.RuneCountInString(data.Email) > 30 ||
		data.Password == "" || utf8.RuneCountInString(data.Password) > 30 {
		log.Printf("fail: invalidinput")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	userName := data.UserName
	email := data.Email
	password := data.Password

	// データをデータベースに挿入しそれを含んだ全userのnameとageを返す
	userInfo, err := usecase.UserRegisterUseCase(userName, email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のUserResForHTTPGetのリスト型に変換
	var userInfoRes UserInfoResForHTTPGet
	userInfoRes.UserId = userInfo.UserId
	userInfoRes.UserName = userInfo.UserName
	userInfoRes.Email = userInfo.Email
	userInfoRes.Channels = []Channel{}

	bytes, err := json.Marshal(userInfoRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
