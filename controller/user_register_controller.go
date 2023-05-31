package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
	"unicode/utf8"
)

type UserResForHTTPPost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
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
	if data.Name == "" || utf8.RuneCountInString(data.Name) > 50 || data.Age < 20 || data.Age > 80 {
		log.Printf("fail: invalidinput, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//受け取ったデータをキーごとに取り出す
	name := data.Name
	age := data.Age

	// データをデータベースに挿入しそれを含んだ全userのnameとageを返す
	users, err := usecase.UserRegisterUseCase(name, age)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	////レスポンス用のUserResForHTTPPostのリスト型に変換
	usersRes := make([]UserResForHTTPPost, 0)
	for _, u := range users {
		usersRes = append(usersRes, UserResForHTTPPost{
			Name: u.Name,
			Age:  u.Age,
		})
	}

	// 変換したユーザーデータリストをjson形式に変換
	jsonResp, err := json.Marshal(usersRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
