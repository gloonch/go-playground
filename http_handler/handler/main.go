package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user", getUser)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	user := User{
		Name:    "John Doe",
		Age:     42,
		Address: "Highway 53, Number 2",
	}

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
