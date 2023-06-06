package controller

import (
	"encoding/json"
	"hackathon-backend/usecase"
	"log"
	"net/http"
)

type UserResForHTTPGet struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
			UserId:   u.UserId,
			UserName: u.UserName,
			Email:    u.Email,
			Password: u.Password,
		})
	}

	bytes, err := json.Marshal(usersRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
