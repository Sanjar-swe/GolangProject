package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
    
var DB *gorm.DB

func InitDB() {
    dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
    var err error

    // Подключение к базе данных
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    // Проверка на ошибки подключения
    if err != nil {
        log.Fatal("Не удалось подключиться к базе данных", err)
    }
}
