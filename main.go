package main

import (
	"encoding/json"
	"time"

	// "io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "time"
)

type RPSLog struct {
	// ID     string `json:"id"`
	Time   time.Time `json:"timestamp"`
	Winner string    `json:"winner"`
	Looser string    `json:"looser"`
	Combo  string    `json:"combo"`
}

var logs []RPSLog

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/index.html")
}

func apiGETHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func apiPOSTHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newLog RPSLog

	err := json.NewDecoder(r.Body).Decode(&newLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newLog.Time = time.Now()

	logs = append(logs, newLog)
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/get", apiGETHandler).Methods("GET")
	router.HandleFunc("/api/post", apiPOSTHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
