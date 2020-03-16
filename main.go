package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Include the project ID here
var projectID string = "go-rpc-270604"

type RPSLog struct {
	// ID     string `json:"id"`
	Time   time.Time `json:"timestamp"`
	Winner string    `json:"winner"`
	Looser string    `json:"looser"`
	Combo  string    `json:"combo"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/index.html")
}

func apiGETHandler(w http.ResponseWriter, r *http.Request) {

	// Set up the client to access the database
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Client creation failed: %v", err)
	}
	defer client.Close()

	// For each of the saved games, add it to an empty logs
	var logs []RPSLog
	iter := client.Collection("games").Documents(ctx)
	for {
		// Get the doc
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		// Load the doc into a struct
		var tempLog *RPSLog
		doc.DataTo(&tempLog)
		logs = append(logs, *tempLog)
	}

	// Send the logs formated in JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func apiPOSTHandler(w http.ResponseWriter, r *http.Request) {

	// Read the new log from the incoming request
	defer r.Body.Close()
	var newLog RPSLog
	err := json.NewDecoder(r.Body).Decode(&newLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	newLog.Time = time.Now() // Timestamp that sucker

	// Open up a connectiont to the database
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		log.Fatalf("Failed to create firebase client: %v", err)
	}
	defer client.Close()

	// Add the entry to the database
	_, _, err = client.Collection("games").Add(ctx, newLog)
	if err != nil {
		log.Fatalf("Failed to add entry: %v", err)
	}
}

func handleRequests() {

	// Good ol' mux router
	router := mux.NewRouter().StrictSlash(true)

	// Handle the various requests
	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/get", apiGETHandler).Methods("GET")
	router.HandleFunc("/api/post", apiPOSTHandler).Methods("POST")

	// Start that sucker up
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
