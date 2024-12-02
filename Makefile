
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


# Генерация кода на основе openapi
# gen:
# 	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml &gt; ./internal/web/tasks/api.gen.go
gen:
	mkdir -p ./internal/web/tasks && oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go



# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # Сервер Запущен
#commit main
commit:
	@read -p "Enter commit message: " msg; \
	git add .; \
	git commit -m "$$msg"; \
	git push -u origin main

commit-test:
	@read -p "Enter commit message: " msg; \
	git add .; \
	git commit -m "$$msg"; \
	git push -u origin test

# 
# Имя выходного документа
DOC_FILE := generated_doc.md

# Команда для генерации документации
MKDOC_CMD := mkdoc

# Цель по умолчанию
.PHONY: doc
doc:
	@for file in *.go; do \
		echo "Generating documentation for $$file..."; \
		$(MKDOC_CMD) -file $$file -output $(DOC_FILE); \
		echo "Generated documentation saved in $(DOC_FILE)"; \
		less $(DOC_FILE); \
		rm -f $(DOC_FILE); \
		echo "Documentation for $$file has been removed."; \
	done
#