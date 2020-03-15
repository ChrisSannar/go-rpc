package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "time"
)

type RPSLog struct {
	ID string `json:"id"`
	// Time string `json:"timestamp"`
	Winner string `json:"winner"`
	Looser string `json:"looser"`
	Combo  string `json:"combo"`
}

var logs []RPSLog

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/index.html")
}

func getGamePage(w http.ResponseWriter, r *http.Request) {
	logs = append(logs, RPSLog{ID: "1", Winner: "Dude", Looser: "Guy", Combo: "PR"})
	fmt.Fprintf(w, "Game page")
}

func apiGETHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func apiPOSTHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
	fmt.Println("POST")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/game", getGamePage).Methods("GET")
	router.HandleFunc("/api/get", apiGETHandler).Methods("GET")
	router.HandleFunc("/api/post", apiGETHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
