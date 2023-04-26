package main

import (
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
