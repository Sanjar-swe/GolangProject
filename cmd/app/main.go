package main

import (
	"net/http"

	"github.com/Sanjar-swe/GolangProject/internal/database"
	"github.com/Sanjar-swe/GolangProject/internal/handlers"
	"github.com/Sanjar-swe/GolangProject/internal/taskService"
	"github.com/gorilla/mux"
)


func main() {

	database.InitDB()
	database.DB.AutoMigrate(&taskService.Message{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewSerivce(repo)
	handler := handlers.NewHandler(service)
	router := mux.NewRouter()
	// router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	// router.HandleFunc("/api/task", PostHandler).Methods("POST")
	// router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	// router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	// router.HandleFunc("/api/messages/{id}", UpdateTaskHandler).Methods("PATCH")
	// router.HandleFunc("/api/messages/{id}", DeleteTaskHandler).Methods("DELETE")

	router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}


// var task string

// type Task struct {
// 	Task string `json:"task"`
// }

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	if task == "" {
// 		fmt.Fprintln(w, "Hello, World!")
// 	} else {
// 		fmt.Fprintf(w, "Hello, %s!", task)
// 	}
// }

// func GetMessages (w http.ResponseWriter, r *http.Request) {
// 	// Обновить GET ручку, чтобы она выводила слайс task (все message, которые лежат в БД)
// 	var messages []Message

//     // Получаем все сообщения из базы данных
// 	if err := database.DB.Find(&messages).Error; err != nil {
// 		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
// 		return
// 	}
//     // Отправляем все сообщения в ответе
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(messages)
// }

// func CreateMessage (w http.ResponseWriter, r *http.Request) {
// 	// Обновить POST ручку, чтобы она записывала содержимое task в БД (Передаем джейсон с полями task и is_done)
// 	var data Message
//     // Декодирование JSON из запроса в структуру
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}

//     // Сохранение в базе данных
// 	if err := database.DB.Create(&data).Error; err != nil{
// 		http.Error(w, "Failed to create message", http.StatusInternalServerError)
// 		return
// 	}

//     // Ответ с подтверждением
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Message created successfully"})
// }

// // POST
// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	var data struct {
// 		Task string `json:"task"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}
// 	task = data.Task
// 	fmt.Println(w, "Task updates succesfully")

// }
// // PATCH handler для обновления задачи по ID
// func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	// Получаем ID из URL
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid task ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Проверяем, существует ли задача с таким ID
// 	var task Message
// 	if err := database.DB.First(&task, id).Error; err != nil {
// 		http.Error(w, "Task not found", http.StatusNotFound)
// 		return
// 	}
// 	// Декодируем тело запроса
// 	var updates struct {
// 		Task *string `json:"task"`
// 		IsDone *bool `json:"is_done"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}
// 	// Обновляем только те поля, которые переданы
// 	if updates.Task != nil {
// 		task.Task = *updates.Task
// 	}
// 	if updates.IsDone != nil {
// 		task.IsDone = *updates.IsDone
// 	}
// 	// Сохраняем изменения в базе данных
// 	if err := database.DB.Save(&task).Error; err != nil{
// 		http.Error(w, "Failed to update task", http.StatusInternalServerError)
// 		return
// 	}
// 	// Возвращаем обновлённую задачу
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(task)
// }

// // DeleteTaskHandler удаляет задачу из базы данных по её ID
// func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	// Получаем ID задачи из параметров URL и преобразуем его в число
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid task ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Проверяем, существует ли задача с указанным ID в базе данных
// 	var task taskService.Message
// 	if err := database.DB.First(&task, id).Error; err != nil {
// 		http.Error(w, "Task not found", http.StatusNotFound)
// 		return
// 	}

// 	// Удаляем найденную задачу из базы данных
// 	if err := database.DB.Delete(&task).Error; err !UpdateTaskHandler= nil {
// 		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
// 		return
// 	}

// 	// Возвращаем подтверждение успешного удаления
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})

// }