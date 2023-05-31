package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UserSearchController(w http.ResponseWriter) {
	users, err := usecase.UserSearchUseCase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//レスポンス用のUserResForHTTPGetのリスト型に変換
	usersRes := make([]UserResForHTTPGet, 0)
	for _, u := range users {
		usersRes = append(usersRes, UserResForHTTPGet{
			Id:   u.Id,
			Name: u.Name,
			Age:  u.Age,
		})
	}

	bytes, err := json.Marshal(usersRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(bytes)
}
