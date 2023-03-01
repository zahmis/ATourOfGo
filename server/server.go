package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ProteinusMember struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	member := ProteinusMember{Name: "sizma", Age: 33}
	json.NewEncoder(w).Encode(member)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	now := time.Now()

	json.NewEncoder(w).Encode(now)
}

func main() {
	http.HandleFunc("/person", personHandler)
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}