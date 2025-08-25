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
go get github.com/yourusername/go-ui-library
```

## Быстрый старт

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/yourusername/go-ui-library/ui"
)

func main() {
    myApp := app.New()
    window := myApp.NewWindow("Моё приложение")
    
    // Создаем кнопку с обработчиком
    button := ui.NewButton("Нажми меня!", func() {
        // Ваш код здесь
    })
    
    // Создаем слайдер
    slider := ui.NewSlider(0, 100, 50)
    
    // Размещаем элементы
    content := container.NewVBox(button, slider)
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
button := ui.NewButtonWithIcon("Текст", icon, func() {
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

Смотрите папку `examples/` для полных примеров приложений.

## Лицензия

MIT License
