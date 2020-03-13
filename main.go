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
	Combo  string `json:"combo"`
}

type RPSLogs []RPSLog

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to RPS!")
	// Input their name, and check a hashtable if someone else has that name already
}

func getGamePage(w http.ResponseWriter, r *http.Request) {
	logs := RPSLogs{
		RPSLog{ID: "1", Winner: "Duders", Combo: "R P"},
	}

	// fmt.Fprintf(w, "Game page")
	json.NewEncoder(w).Encode(logs)
}

func postGamePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST working")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/game", getGamePage).Methods("GET")
	router.HandleFunc("/game", postGamePage).Methods("POST")
	router.HandleFunc("/api", apiHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
