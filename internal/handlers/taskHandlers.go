package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *taskService.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, error := h.Service.GetAllTasks()
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	createdTask, error := h.Service.CreateTask(task)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	updatedTask, error := h.Service.UpdateTaskByID(task.ID, task)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	error := h.Service.DeleteTaskByID(id)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}