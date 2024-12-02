package handlers

import (
	"context"

	"github.com/Sanjar-swe/GolangProject/internal/taskService"
	"github.com/Sanjar-swe/GolangProject/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	panic("unimplemented")
}

// DeleteTasksId implements tasks.StrictServerInterface.
// func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
// 	// Обращаемся к сервису и удаляем задачу по ID
// 	err := h.Service.DeleteTaskByID(uint(request.Id))
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Возвращаем сообщение об успешном удалении
// 	return tasks.DeleteTasksId200JSONResponse{
// 		Message: "Task deleted successfully",
// 	}, nil
// }

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Распаковываем тело запроса
	taskRequest := request.Body

	// Создаем объект для обновления задачи
	taskToUpdate := taskService.Message{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	// Обновляем задачу в сервисе
	updatedTask, err := h.Service.UpdateTaskByID(uint(request.Id), taskToUpdate)
	if err != nil {
		return nil, err
	}

	// Создаем структуру ответа
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	// Возвращаем респонс
	return response, nil
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Message{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	reponse := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Возвращаем респонс
	return reponse, nil

}

// Нужна для создания структуры Handler на этапе инициализации приложения
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

// func (h *Handler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	tasks, error := h.Service.GetAllTasks()
// 	if error != nil {
// 		http.Error(w, error.Error(), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// }

// func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	var task taskService.Message
// 	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}
// 	createdTask, error := h.Service.CreateTask(task)
// 	if error != nil {
// 		http.Error(w, error.Error(), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(createdTask)
// }

// func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid task ID", http.StatusBadRequest)
// 		return
// 	}
// 	var task taskService.Message
// 	if err := database.DB.First(&task, id).Error; err != nil {
// 		http.Error(w, "Task not found", http.StatusNotFound)
// 		return
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}
// 	updatedTask, error := h.Service.UpdateTaskByID(uint(id), task)
// 	if error != nil {
// 		http.Error(w, error.Error(), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(updatedTask)
// }

// func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid task ID", http.StatusBadRequest)
// 		return
// 	}
// 	error := h.Service.DeleteTaskByID(uint(id))
// 	if error != nil {
// 		http.Error(w, error.Error(), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
// }
