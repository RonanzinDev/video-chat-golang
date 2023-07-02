package server

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participant describy a single entity in the hashMap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// RoomMap  is the main HashMap [roomID string] -> [[]Participant]
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Init() => Initialises the RoomMap struct
func (rm *RoomMap) Init() {
	rm.Map = make(map[string][]Participant)
}

// Get() => Will get the array of Participants in the room
func (rm *RoomMap) Get(roomID string) []Participant {
	rm.Mutex.RLock()
	defer rm.Mutex.RUnlock()
	return rm.Map[roomID]
}

// CreateRoom() => Create a Unique ID and insert into the hashmap
func (rm *RoomMap) CreateRoom() string {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()
	rand.Seed(time.Now().UnixMilli())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
		fmt.Println(letters[rand.Intn(len(letters))])
	}
	roomID := string(b)
	rm.Map[roomID] = []Participant{}
	return roomID
}

// InsertIntoRoom() => will insert a participant into the room
func (rm *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	rm.Mutex.RLock()
	defer rm.Mutex.RUnlock()
	p := Participant{Host: host, Conn: conn}
	log.Printf("Insert into room with RoomID: %v\n", roomID)
	rm.Map[roomID] = append(rm.Map[roomID], p)
}

// DeleteRoom() => Deletes a room
func (rm *RoomMap) DeleteRoom(roomID string) {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()
	delete(rm.Map, roomID)
}
