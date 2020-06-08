package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	db "github.com/pavelkrolevets/opensolar_eth/db"
	"github.com/pavelkrolevets/opensolar_eth/models"
	"io/ioutil"
	"log"
	"net/http"
)
type Auth struct {
 	auth bool
 }
func UserAuth (w http.ResponseWriter, r *http.Request) {
	var u *models.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &u)
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	if u.Login == nil {
		log.Fatal("Login is missing or null!")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Provide username"))
	}
	if u.Password == nil {
		log.Fatal("Password is missing or null!")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Provide password"))
	}
	var store db.Store
	store.Path = "bolt_main"
	_, err = store.GetUser(u)
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - User doesnt exist"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Auth{auth: true})
}

 func NewUser (w http.ResponseWriter, r *http.Request) {
	 var u *models.User
	 body, _ := ioutil.ReadAll(r.Body)
	 err := json.Unmarshal(body, &u)
	 if err != nil {
		 fmt.Println("Err", err)
		 return
	 }
	 if u.Login == nil {
		 log.Fatal("Login is missing or null!")
		 w.WriteHeader(http.StatusInternalServerError)
		 w.Write([]byte("500 - Provide username"))
	 }
	 if u.Password == nil {
		 log.Fatal("Password is missing or null!")
		 w.WriteHeader(http.StatusInternalServerError)
		 w.Write([]byte("500 - Provide password"))
	 }
	 // Save hash of user password to db
	 h := sha256.New()
	 h.Write([]byte(*u.Password))
	 hex := base64.URLEncoding.EncodeToString(h.Sum(nil))
	 u.Password = &hex
	 var store db.Store
	 store.Path = "bolt_main"
	 err = store.StoreUser(u)
	 if err!= nil {
		 w.WriteHeader(http.StatusInternalServerError)
		 w.Write([]byte("500 - User exist, please login"))
		 return
	 }
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusCreated)
	 json.NewEncoder(w).Encode(u)
 }

 func GetUser (w http.ResponseWriter, r *http.Request){
	 var u *models.User
	 body, _ := ioutil.ReadAll(r.Body)
	 err := json.Unmarshal(body, &u)
	 if err != nil {
		 fmt.Println("Err", err)
		 return
	 }
	 if u.Login == nil {
		 log.Fatal("Login is missing or null!")
	 }
	 if u.Login == nil {
		 log.Fatal("Password is missing or null!")
	 }
	 var store db.Store
	 store.Path = "bolt_main"
	 usr, err := store.GetUser(u)
	 if err!= nil {
		 fmt.Println("cant get user")
	 }
	 fmt.Println("Successfully got user", *usr.Password)
 }
