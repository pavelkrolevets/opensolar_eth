package models

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
)

type User struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func (u *User) StoreUser() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("opensolar_eth.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	b, err := tx.CreateBucketIfNotExists(usersBucket)
	if err != nil {
		log.Fatal(err)
	}

	encoded, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	err = b.Put([]byte(u.Login), encoded)
	if err != nil {
		log.Fatal(err)
	}
}
