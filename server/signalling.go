package server

import (
	"fmt"
	"net/http"
)

// AllRooms => is the all global hashMap for the server
var AllRooms RoomMap

// CreateRoomRequestHadler(w,r) => Create a room and return RoomID
func CreateRoomRequestHadler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// JoinRoomRequestHadler(w,r) => Will join the client in a particular room
func JoinRoomRequestHadler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
