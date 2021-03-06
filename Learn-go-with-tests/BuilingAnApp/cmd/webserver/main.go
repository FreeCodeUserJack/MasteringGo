package main

import (
	"log"
	"net/http"

	poker "github.com/FreeCodeUserJack/MasteringGo/Learn-go-with-tests/BuilingAnApp"
)

const dbFileName = "game.db.json"

// func NewInMemoryPlayerStore() *InMemoryPlayerStore {
// 	return &InMemoryPlayerStore{map[string]int{}, &sync.Mutex{}}
// }

// type InMemoryPlayerStore struct {
// 	store map[string]int
// 	sMutext *sync.Mutex
// }

// func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
// 	return i.store[name]
// }

// func (i *InMemoryPlayerStore) RecordWin(name string) {
// 	i.sMutext.Lock()
// 	i.store[name]++
// 	i.sMutext.Unlock()
// }

// func (i *InMemoryPlayerStore) GetLeague() (league League) {
// 	for k, v := range i.store {
// 		league = append(league, Player{
// 			Name: k,
// 			Win: v,
// 		})
// 	}
// 	return
// }

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}