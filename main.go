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

func CreateMessage (w http.ResponseWriter, r *http.Request) {
	// Обновить POST ручку, чтобы она записывала содержимое task в БД (Передаем джейсон с полями task и is_done)
	var data Message
    // Декодирование JSON из запроса в структуру
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

    // Сохранение в базе данных
	if err := DB.Create(&data).Error; err != nil{
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

    // Ответ с подтверждением
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Message created successfully"})
}

func GetMessages (w http.ResponseWriter, r *http.Request) {
	// Обновить GET ручку, чтобы она выводила слайс task (все message, которые лежат в БД)
	var messages []Message

    // Получаем все сообщения из базы данных
	if err := DB.Find(&messages).Error; err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}
    // Отправляем все сообщения в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages) 
}

func main() {

	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}