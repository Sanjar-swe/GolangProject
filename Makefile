
# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:postgres@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)


# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер

commit:
	@read -p "Enter commit message: " msg; \
	git add .; \
	git commit -m "$$msg"; \
	git push -u origin main
