package main

import (
	"log"
	"net/http"

	"github.com/ronanzindev/video-chat-golang/server"
)

func main() {
	server.AllRooms.Init()
	http.HandleFunc("/create", server.CreateRoomRequestHadler)
	http.HandleFunc("/join", server.JoinRoomRequestHadler)

	log.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
