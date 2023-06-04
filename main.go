package main

import (
	_ "github.com/go-sql-driver/mysql"
	"hackathon-backend/controller"
	"hackathon-backend/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func headerSet(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://hackathon-frontend-d7e49jyei-satowo.vercel.app/register")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Content-Type", "application/json")
	return
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	headerSet(w)
	switch r.Method {
	case http.MethodOptions:
		w.Header()
	case http.MethodGet:
		controller.UserSearchController(w)

	case http.MethodPost:
		controller.UserRegisterController(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func channelHandler(w http.ResponseWriter, r *http.Request) {
	headerSet(w)
	switch r.Method {
	case http.MethodOptions:
		w.Header()
	case http.MethodGet:
		controller.ChannelSearchController(w, r)

	//case http.MethodPost:

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func channelMemberHandler(w http.ResponseWriter, r *http.Request) {
	headerSet(w)
	switch r.Method {
	case http.MethodOptions:
		w.Header()
	case http.MethodGet:
		controller.ChannelMemberSearchController(w, r)

	//case http.MethodPost:

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	headerSet(w)
	switch r.Method {
	case http.MethodOptions:
		w.Header()
	case http.MethodGet:
		controller.MessageSearchController(w, r)

	//case http.MethodPost:

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// /userでsign upまたはsigh inのどちらかのリクエストを受け取る
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/channel", channelHandler)
	http.HandleFunc("/message", MessageHandler)
	http.HandleFunc("/channel_member", channelMemberHandler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		err := dao.DataBaseClose()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
