package main

import (
	"encoding/json"
	"net/http"
	"fmt"
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

func main() {
	http.HandleFunc("/person", personHandler)
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}