package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type Task struct {
	Task string `json:"task"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if task == "" {
		fmt.Fprintln(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", task)
	}
}


func PostHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Task string `json:"task"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	task = data.Task
	fmt.Println(w, "Task updates succesfully")

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}