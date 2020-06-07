package db

import (
	"encoding/json"
	"fmt"
	"github.com/pavelkrolevets/opensolar_eth/models"
	"log"
	"github.com/boltdb/bolt"
)

type Store struct {
	Path string
}

func (s *Store) StoreUser(u *models.User) error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Users"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(u)
		if err != nil {
			return err
		}
		return b.Put([]byte(u.Login), encoded)
	})
	return err
}

func (s *Store) GetUser (u *models.User) {
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		v := b.Get([]byte(u.Login))
		err := json.Unmarshal(v, u)
		if err!= nil{
			log.Fatal("cant unmarshal json", err)
			return err
		}
		fmt.Printf("User is : %s\n", )
		return nil
	})

}
