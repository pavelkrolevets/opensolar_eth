package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/pavelkrolevets/opensolar_eth/db"
	"github.com/pavelkrolevets/opensolar_eth/models"
	"io/ioutil"
	"net/http"
)

func Auth (w http.ResponseWriter, r *http.Request){
	u := models.User{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &u)
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	fmt.Printf("User login, password: ", u)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	db.StoreUser(&u)
}

