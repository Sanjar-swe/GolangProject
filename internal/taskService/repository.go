package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Message) (Message, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Message, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task и ошибку
	UpdateTaskByID(id uint, task Message) (Message, error) 
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}


type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Message) (Message, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Message, error) {
	var tasks []Message
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Message) (Message, error) {
	result := r.db.Where("id = ?", id).Updates(&task)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Where("id = ?", id).Delete(&Message{})
	return result.Error
}
