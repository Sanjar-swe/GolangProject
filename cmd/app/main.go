package main

import (
	"log"
	"net/http"

	"github.com/Sanjar-swe/GolangProject/internal/database"
	"github.com/Sanjar-swe/GolangProject/internal/handlers"
	"github.com/Sanjar-swe/GolangProject/internal/taskService"
	"github.com/Sanjar-swe/GolangProject/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {

	database.InitDB()
	database.DB.AutoMigrate(&taskService.Message{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewSerivce(repo)
	handler := handlers.NewHandler(service)
	// router := mux.NewRouter()
	

	// Инициализируем echo
	e := echo.New()

	
	// middleware.Logger() - логирование запросов
	e.Use(middleware.Logger())
	// middleware.Recover() - перехватывает панику
	e.Use(middleware.Recover())



	// Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler,nil)
	tasks.RegisterHandlers(e, strictHandler)

	// Регистрация маршрутов с адаптерами
	e.GET("/api/tasks", func(c echo.Context) error {
		response, err := handler.GetTasks(c.Request().Context(), tasks.GetTasksRequestObject{})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/api/tasks", func(c echo.Context) error {
		var request tasks.PostTasksRequestObject
		if err := c.Bind(&request); err != nil {
			return err
		}
		response, err := handler.PostTasks(c.Request().Context(), request)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
		})

	// 	e.PATCH("/api/tasks/:id", func(c echo.Context) error {
	// 		var request tasks.PatchTasksIdRequestObject
	// 		id := c.Param("id") // Получаем id из параметров
	// 		if err := c.Bind(&request); err != nil {
	// 			return err
	// 		}
	// 		response, err := handler.PatchTasksId(c.Request().Context(), request) // Используем request
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return c.JSON(http.StatusOK, response)
	// 	})
		
	// 	e.DELETE("/api/tasks/:id", func(c echo.Context) error {
	// 		id := c.Param("id") // Получаем id из параметров
	// 		// Предполагается, что DeleteTasksId принимает id как строку
	// 		err := handler.DeleteTasksId(c.Request().Context(), id) // Убедитесь, что DeleteTasksId принимает id
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return c.NoContent(http.StatusNoContent)
	// 	})

	

	// Регистрация маршрутов
	// e.GET("/api/get", handler.GetTaskHandler)
	// e.POST("/api/post", handler.PostTaskHandler)
	// e.PATCH("/api/tasks/:id", handler.PatchTaskHandler)
	// e.DELETE("/api/tasks/:id", handler.DeleteTaskHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}


	

	// router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	// router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	// router.HandleFunc("/api/tasks/{id}", handler.PatchTaskHandler).Methods("PATCH")
	// router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	// http.ListenAndServe(":8080", router)
}

