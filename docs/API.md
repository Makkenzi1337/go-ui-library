# API Documentation

## Обзор

Go UI Library предоставляет простой и интуитивный API для создания современных десктопных приложений на Go. Библиотека построена на основе Fyne и предоставляет улучшенный интерфейс для работы с UI компонентами.

## Основные концепции

### Приложение (App)
Приложение является корневым объектом, который управляет жизненным циклом приложения.

```go
app := ui.NewApp()
window := app.NewWindow("Моё приложение")
app.Run()
```

### Окно (Window)
Окно представляет собой основную область приложения.

```go
window := app.NewWindow("Заголовок")
window.Resize(ui.NewSize(800, 600))
window.Center()
window.ShowAndRun()
```

## Базовые компоненты

### Кнопка (Button)
```go
// Простая кнопка
button := ui.NewButton("Текст кнопки", func() {
    // Обработчик клика
})

// Кнопка с иконкой
button := ui.NewButtonWithIcon("Текст", icon, func() {
    // Обработчик клика
})
```

### Метка (Label)
```go
label := ui.NewLabel("Текст метки")
label.SetText("Новый текст")
```

### Поле ввода (Entry)
```go
entry := ui.NewEntry()
entry.SetPlaceHolder("Подсказка")
entry.OnChanged = func(text string) {
    // Обработчик изменений
}
```

### Слайдер (Slider)
```go
slider := ui.NewSlider(0, 100, 50)
slider.OnChanged = func(value float64) {
    // Обработчик изменений
}
```

### Выпадающий список (Select)
```go
selectWidget := ui.NewSelect([]string{"Опция 1", "Опция 2"})
selectWidget.OnChanged = func(selected string) {
    // Обработчик изменений
}
```

### Чекбокс (CheckBox)
```go
checkbox := ui.NewCheckBox("Текст чекбокса")
checkbox.OnChanged = func(checked bool) {
    // Обработчик изменений
}
```

### Прогресс-бар (ProgressBar)
```go
progress := ui.NewProgressBar()
progress.SetValue(0.5) // 50%
```

## Контейнеры

### Вертикальный контейнер (VBox)
```go
container := ui.NewVBox(
    ui.NewLabel("Элемент 1"),
    ui.NewLabel("Элемент 2"),
)
```

### Горизонтальный контейнер (HBox)
```go
container := ui.NewHBox(
    ui.NewLabel("Элемент 1"),
    ui.NewLabel("Элемент 2"),
)
```

### Сетка (Grid)
```go
container := ui.NewGrid(2, // 2 колонки
    ui.NewLabel("Ячейка 1"),
    ui.NewLabel("Ячейка 2"),
    ui.NewLabel("Ячейка 3"),
    ui.NewLabel("Ячейка 4"),
)
```

### Контейнер с границами (Border)
```go
container := ui.NewBorder(
    top,    // верхняя граница
    bottom, // нижняя граница
    left,   // левая граница
    right,  // правая граница
    center, // центральная область
)
```

### Карточка (Card)
```go
card := ui.NewCard(
    "Заголовок",
    "Подзаголовок",
    content, // содержимое карточки
)
```

## Вкладки (Tabs)

```go
tabs := ui.NewTabs(
    ui.NewTabItem("Вкладка 1", content1),
    ui.NewTabItem("Вкладка 2", content2),
    ui.NewTabItemWithIcon("Вкладка 3", icon, content3),
)
```

## Формы (Form)

```go
form := ui.NewForm(
    ui.NewFormItem("Имя:", ui.NewEntry()),
    ui.NewFormItem("Email:", ui.NewEntry()),
    ui.NewFormItem("Возраст:", ui.NewEntry()),
)
```

## Панель инструментов (Toolbar)

```go
toolbar := ui.NewToolbar(
    ui.ToolbarAction(icon, "Подсказка", func() {
        // Действие
    }),
    ui.ToolbarSeparator(),
    ui.ToolbarAction(icon2, "Подсказка 2", func() {
        // Действие 2
    }),
)
```

## Меню (Menu)

```go
fileMenu := ui.Menu("Файл",
    ui.MenuItem("Новый", icon, func() {
        // Действие
    }),
    ui.MenuItem("Открыть", icon, func() {
        // Действие
    }),
    ui.MenuItemSeparator(),
    ui.MenuItem("Выход", icon, func() {
        // Действие
    }),
)

menuBar := ui.NewMenuBar(fileMenu)
```

## Диалоговые окна

### Информационное окно
```go
ui.ShowInfo("Заголовок", "Сообщение", window)
```

### Окно с ошибкой
```go
ui.ShowError(err, window)
```

### Окно подтверждения
```go
ui.ShowConfirm("Заголовок", "Сообщение", func(confirmed bool) {
    if confirmed {
        // Пользователь подтвердил
    }
}, window)
```

### Окно ввода
```go
ui.ShowInput("Заголовок", "Сообщение", func(text string) {
    // Обработка введенного текста
}, window)
```

### Окно выбора файла
```go
ui.ShowFileOpen(func(uri fyne.URI) {
    // Обработка выбранного файла
}, window)
```

### Окно сохранения файла
```go
ui.ShowFileSave(func(writer fyne.URIWriteCloser) {
    // Обработка сохранения файла
}, window)
```

### Окно с прогрессом
```go
progress := ui.ShowProgress("Заголовок", "Сообщение", window)
progress.SetProgress(0.5) // 50%
progress.Hide()
```

## Утилиты

### Цвета
```go
// Создание цвета из RGBA значений
color := ui.NewRGBA(255, 0, 0, 255) // Красный

// Создание цвета из float значений
color := ui.NewColor(1.0, 0.0, 0.0, 1.0) // Красный

// Предопределенные цвета
ui.ColorRed
ui.ColorGreen
ui.ColorBlue
ui.ColorBlack
ui.ColorWhite
ui.ColorGray
```

### Размеры и позиции
```go
size := ui.NewSize(800, 600)
pos := ui.NewPosition(100, 100)
```

### Форматирование
```go
// Форматирование размера файла
sizeStr := ui.FormatSize(1024) // "1.0 KB"

// Форматирование длительности
durationStr := ui.FormatDuration(65) // "1m 5s"
```

## Компоновка

### Готовые компоновки
```go
// Компоновка с панелью инструментов
content := ui.CreateToolbarLayout(toolbar, mainContent)

// Компоновка со строкой состояния
content := ui.CreateStatusBarLayout(mainContent, statusBar)

// Полная компоновка
content := ui.CreateFullLayout(toolbar, mainContent, statusBar)

// Компоновка с боковой панелью
content := ui.CreateSidebarLayout(sidebar, mainContent)

// Разделенная компоновка
content := ui.CreateSplitLayout(leftPanel, rightPanel)
```

## Темы

```go
// Установка светлой темы
ui.SetLightTheme()

// Установка темной темы
ui.SetDarkTheme()

// Установка пользовательской темы
ui.SetTheme(customTheme)
```

## Лучшие практики

1. **Обработка событий**: Всегда используйте обработчики событий для реагирования на действия пользователя
2. **Компоновка**: Используйте подходящие контейнеры для организации элементов
3. **Темы**: Применяйте темы для обеспечения консистентного внешнего вида
4. **Обработка ошибок**: Используйте диалоговые окна для отображения ошибок
5. **Прогресс**: Показывайте прогресс для длительных операций

## Примеры

Смотрите папку `examples/` для полных примеров использования библиотеки:

- `examples/simple/` - Простой пример с базовыми компонентами
- `examples/advanced/` - Продвинутый пример с вкладками и меню
- `examples/cards/` - Пример с карточками и списками
