package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
	"unicode/utf8"
)

type UserResForHTTPPost struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegisterController(w http.ResponseWriter, r *http.Request) {
	// リクエストのボディを読み込み
	var data UserResForHTTPPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("fail: json.NewDecoder(r.Body).Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Nameが空文字か50字以上、Ageが81歳以上19歳以下の時エラーコード400を返す
	if data.UserName == "" || utf8.RuneCountInString(data.UserName) > 10 ||
		data.Email == "" || utf8.RuneCountInString(data.Email) > 30 ||
		data.Password == "" || utf8.RuneCountInString(data.Password) > 20 {
		log.Printf("fail: invalidinput")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	name := data.UserName
	email := data.Email
	pwd := data.Password

	// データをデータベースに挿入しそれを含んだ全userのnameとageを返す
	err = usecase.UserRegisterUseCase(name, email, pwd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
