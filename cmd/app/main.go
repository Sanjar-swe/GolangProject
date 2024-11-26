package main

import (
	"log"
	"net/http"

	"github.com/Sanjar-swe/GolangProject/internal/database"
	"github.com/Sanjar-swe/GolangProject/internal/handlers"
	"github.com/Sanjar-swe/GolangProject/internal/taskService"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


func main() {

	database.InitDB()
	database.DB.AutoMigrate(&taskService.Message{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewSerivce(repo)
	handler := handlers.NewHandler(service)
	router := mux.NewRouter()
	

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	// middleware.Logger() - логирование запросов
	e.Use(middleware.Logger())
	// middleware.Recover() - перехватывает панику
	e.Use(middleware.Recover())


	// Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler,nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}


	router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

