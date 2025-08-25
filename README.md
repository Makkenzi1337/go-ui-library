# Go UI Library

Современная библиотека для создания десктопных приложений на Go с простым и интуитивным API.

## Особенности

- 🚀 **Простота и скорость разработки** - интуитивный API для быстрого создания интерфейсов
- 🎨 **Современный стиль** - красивые компоненты из коробки
- 📦 **Готовые компоненты** - кнопки, слайдеры, выпадающие списки и многое другое
- 🖱️ **Drag & Drop** - визуальное размещение элементов
- 🔧 **Легкая настройка** - простые функции для обработки событий

## Установка

```bash
go get github.com/Makkenzi1337/go-ui-library
```

## Быстрый старт

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
    "github.com/Makkenzi1337/go-ui-library"
)

func main() {
    // Создаем приложение
    app := ui.NewApp()
    
    // Создаем главное окно
    window := app.NewWindow("Моё приложение")
    window.Resize(fyne.NewSize(400, 300))
    window.Center()
    
    // Создаем кнопку с обработчиком
    button := ui.NewButton("Нажми меня!", func() {
        ui.ShowInfo("Привет!", "Кнопка была нажата!", window)
    })
    
    // Создаем слайдер
    slider := ui.NewSlider(0, 100, 50)
    slider.OnChanged = func(value float64) {
        fmt.Printf("Значение: %.2f\n", value)
    }
    
    // Размещаем элементы
    content := ui.NewVBox(button, slider)
    window.SetContent(content)
    
    window.ShowAndRun()
}
```

## Компоненты

### Кнопки
```go
// Простая кнопка
button := ui.NewButton("Текст кнопки", func() {
    fmt.Println("Кнопка нажата!")
})

// Кнопка с иконкой
button := ui.NewButtonWithIcon("Текст", theme.InfoIcon(), func() {
    // обработчик
})
```

### Слайдеры
```go
// Слайдер с диапазоном
slider := ui.NewSlider(0, 100, 50)
slider.OnChanged = func(value float64) {
    fmt.Printf("Значение: %.2f\n", value)
}
```

### Выпадающие списки
```go
// Простой список
list := ui.NewSelect([]string{"Опция 1", "Опция 2", "Опция 3"})
list.OnChanged = func(selected string) {
    fmt.Printf("Выбрано: %s\n", selected)
}
```

### Поля ввода
```go
// Текстовое поле
input := ui.NewEntry()
input.SetPlaceHolder("Введите текст...")
input.OnChanged = func(text string) {
    fmt.Printf("Введено: %s\n", text)
}
```

## Примеры

Смотрите папку `examples/` для полных примеров приложений:

- **simple/** - Простой пример с базовыми компонентами
- **cards/** - Пример с карточками, списками и таблицами
- **advanced/** - Продвинутый пример с формами и вкладками

### Запуск примеров

```bash
# Простой пример
make simple

# Пример с карточками
make cards

# Продвинутый пример
make advanced

# Все примеры
make examples
```

## Требования

- Go 1.21+
- Fyne v2.4.1+

## Лицензия

MIT License
