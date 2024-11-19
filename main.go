package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// POST
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
// PATCH handler для обновления задачи по ID
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL	
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Проверяем, существует ли задача с таким ID
	var task Message
	if err := DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	// Декодируем тело запроса
	var updates struct {
		Task *string `json:"task"`
		IsDone *bool `json:"is_done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	// Обновляем только те поля, которые переданы
	if updates.Task != nil {
		task.Task = *updates.Task
	}	
	if updates.IsDone != nil {
		task.IsDone = *updates.IsDone
	}
	// Сохраняем изменения в базе данных
	if err := DB.Save(&task).Error; err != nil{
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
	// Возвращаем обновлённую задачу
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)	
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
	router.HandleFunc("/api/messages/{id}", UpdateTaskHandler).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}