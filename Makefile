# Makefile для Go UI Library

.PHONY: help build test clean examples simple advanced cards

# Переменные
BINARY_NAME=go-ui-library
BUILD_DIR=build
EXAMPLES_DIR=examples

# Цвета для вывода
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

help: ## Показать справку
	@echo "$(GREEN)Доступные команды:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

build: ## Собрать библиотеку
	@echo "$(GREEN)Сборка библиотеки...$(NC)"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

test: ## Запустить тесты
	@echo "$(GREEN)Запуск тестов...$(NC)"
	go test ./...

clean: ## Очистить сборки
	@echo "$(GREEN)Очистка...$(NC)"
	rm -rf $(BUILD_DIR)
	go clean

examples: ## Собрать все примеры
	@echo "$(GREEN)Сборка примеров...$(NC)"
	@mkdir -p $(BUILD_DIR)/examples
	@for dir in $(EXAMPLES_DIR)/*; do \
		if [ -d "$$dir" ]; then \
			echo "Сборка $$(basename $$dir)..."; \
			go build -o $(BUILD_DIR)/examples/$$(basename $$dir) $$dir; \
		fi \
	done

simple: ## Запустить простой пример
	@echo "$(GREEN)Запуск простого примера...$(NC)"
	cd $(EXAMPLES_DIR)/simple && go run main.go

advanced: ## Запустить продвинутый пример
	@echo "$(GREEN)Запуск продвинутого примера...$(NC)"
	cd $(EXAMPLES_DIR)/advanced && go run main.go

cards: ## Запустить пример с карточками
	@echo "$(GREEN)Запуск примера с карточками...$(NC)"
	cd $(EXAMPLES_DIR)/cards && go run main.go

install: ## Установить зависимости
	@echo "$(GREEN)Установка зависимостей...$(NC)"
	go mod tidy
	go mod download

deps: install ## Алиас для install

run-all: ## Запустить все примеры по очереди
	@echo "$(GREEN)Запуск всех примеров...$(NC)"
	@$(MAKE) simple
	@echo "$(YELLOW)Нажмите Enter для запуска следующего примера...$(NC)"
	@read
	@$(MAKE) advanced
	@echo "$(YELLOW)Нажмите Enter для запуска следующего примера...$(NC)"
	@read
	@$(MAKE) cards

dev: ## Режим разработки (запуск с горячей перезагрузкой)
	@echo "$(GREEN)Режим разработки...$(NC)"
	@echo "$(YELLOW)Установите air для горячей перезагрузки: go install github.com/cosmtrek/air@latest$(NC)"
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "$(YELLOW)Air не установлен. Запускаем обычный режим...$(NC)"; \
		go run examples/simple/main.go; \
	fi

lint: ## Проверка кода
	@echo "$(GREEN)Проверка кода...$(NC)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "$(YELLOW)golangci-lint не установлен. Установите: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(NC)"; \
	fi

format: ## Форматирование кода
	@echo "$(GREEN)Форматирование кода...$(NC)"
	go fmt ./...

vet: ## Проверка кода с помощью go vet
	@echo "$(GREEN)Проверка кода...$(NC)"
	go vet ./...

check: format vet lint ## Проверить код (форматирование + vet + lint)

release: ## Создать релиз
	@echo "$(GREEN)Создание релиза...$(NC)"
	@echo "$(YELLOW)Не забудьте обновить версию в go.mod$(NC)"
	git tag -a v1.0.0 -m "Release v1.0.0"
	git push origin v1.0.0

.DEFAULT_GOAL := help
