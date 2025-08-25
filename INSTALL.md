# Инструкции по установке и использованию

## Требования

- Go 1.21 или выше
- Git
- Системные зависимости для Fyne (см. [документацию Fyne](https://fyne.io/docs/))

### Системные зависимости

#### Ubuntu/Debian:
```bash
sudo apt-get install gcc libc6-dev libx11-dev xorg-dev libxtst-dev libpng++-dev
```

#### Fedora:
```bash
sudo dnf install gcc libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libXxf86vm-devel
```

#### macOS:
```bash
# Установите Xcode Command Line Tools
xcode-select --install
```

#### Windows:
```bash
# Установите MinGW-w64 или используйте WSL
```

## Установка

### 1. Клонирование репозитория
```bash
git clone <repository-url>
cd go-ui-library
```

### 2. Установка зависимостей
```bash
go mod tidy
go mod download
```

### 3. Проверка установки
```bash
go test ./...
```

## Быстрый старт

### Простой пример
```bash
cd examples/simple
go run main.go
```

### Продвинутый пример
```bash
cd examples/advanced
go run main.go
```

### Пример с карточками
```bash
cd examples/cards
go run main.go
```

## Использование в вашем проекте

### 1. Создайте новый проект
```bash
mkdir my-app
cd my-app
go mod init my-app
```

### 2. Добавьте зависимость
```bash
go get <repository-url>
```

### 3. Создайте main.go
```go
package main

import (
    "go-ui-library/ui"
)

func main() {
    app := ui.NewApp()
    window := app.NewWindow("Моё приложение")
    
    button := ui.NewButton("Привет!", func() {
        ui.ShowInfo("Привет", "Кнопка нажата!", window)
    })
    
    window.SetContent(button)
    window.ShowAndRun()
}
```

### 4. Запустите приложение
```bash
go run main.go
```

## Использование Makefile

Библиотека включает Makefile для удобной работы:

```bash
# Показать справку
make help

# Установить зависимости
make install

# Запустить тесты
make test

# Собрать все примеры
make examples

# Запустить простой пример
make simple

# Запустить продвинутый пример
make advanced

# Запустить пример с карточками
make cards

# Проверить код
make check

# Очистить сборки
make clean
```

## Структура проекта

```
go-ui-library/
├── ui.go                 # Основной файл библиотеки
├── go.mod               # Зависимости Go
├── README.md            # Документация
├── INSTALL.md           # Инструкции по установке
├── Makefile             # Команды для сборки
├── internal/            # Внутренние пакеты
│   └── ui/             # Основная логика UI
│       ├── ui.go       # Базовые компоненты
│       ├── components.go # Дополнительные компоненты
│       ├── utils.go    # Утилиты и хелперы
│       └── ui_test.go  # Тесты
├── examples/            # Примеры использования
│   ├── simple/         # Простой пример
│   ├── advanced/       # Продвинутый пример
│   └── cards/          # Пример с карточками
└── docs/               # Документация
    └── API.md          # API документация
```

## Компоненты библиотеки

### Базовые компоненты
- **Button** - Кнопки с обработчиками событий
- **Label** - Текстовые метки
- **Entry** - Поля ввода
- **Slider** - Слайдеры
- **Select** - Выпадающие списки
- **CheckBox** - Чекбоксы
- **ProgressBar** - Прогресс-бары

### Контейнеры
- **VBox** - Вертикальные контейнеры
- **HBox** - Горизонтальные контейнеры
- **Grid** - Сеточные контейнеры
- **Border** - Контейнеры с границами
- **Card** - Карточки

### Сложные компоненты
- **Tabs** - Вкладки
- **Form** - Формы
- **List** - Списки
- **Table** - Таблицы
- **Toolbar** - Панели инструментов
- **MenuBar** - Меню

### Диалоговые окна
- **ShowInfo** - Информационные окна
- **ShowError** - Окна с ошибками
- **ShowConfirm** - Окна подтверждения
- **ShowInput** - Окна ввода
- **ShowFileOpen** - Окна выбора файлов
- **ShowFileSave** - Окна сохранения файлов
- **ShowProgress** - Окна с прогрессом

## Темы и стили

Библиотека поддерживает светлую и темную темы:

```go
// Установка светлой темы
ui.SetLightTheme()

// Установка темной темы
ui.SetDarkTheme()
```

## Лучшие практики

1. **Обработка событий**: Всегда используйте обработчики событий
2. **Компоновка**: Используйте подходящие контейнеры
3. **Темы**: Применяйте темы для консистентности
4. **Обработка ошибок**: Используйте диалоговые окна
5. **Прогресс**: Показывайте прогресс для длительных операций

## Устранение неполадок

### Проблемы с зависимостями
```bash
go clean -modcache
go mod tidy
```

### Проблемы с системными зависимостями
Убедитесь, что установлены все необходимые системные зависимости для вашей ОС.

### Проблемы с компиляцией
```bash
go vet ./...
go fmt ./...
```

## Поддержка

- Документация: `docs/API.md`
- Примеры: `examples/`
- Тесты: `internal/ui/ui_test.go`

## Лицензия

MIT License
