package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error) 

	DeleteTaskByID(id uint) error
}


type TaskRepository struct {
	db *gorm.DB
}

func newTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *TaskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	result := r.db.Where("id = ?", id).Updates(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *TaskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Where("id = ?", id).Delete(&Task{})
	return result.Error
}
