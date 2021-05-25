package main

import (
	"log"
	"net/http"
	"os"
	"sync"
)

const dbFileName = "game.db.json"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, &sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	sMutext *sync.Mutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.sMutext.Lock()
	i.store[name]++
	i.sMutext.Unlock()
}

func (i *InMemoryPlayerStore) GetLeague() (league League) {
	for k, v := range i.store {
		league = append(league, Player{
			Name: k,
			Win: v,
		})
	}
	return
}

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}