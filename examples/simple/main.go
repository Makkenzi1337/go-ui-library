package main

import (
	"fmt"
	"fyne.io/fyne/v2/theme"
	"go-ui-library"
)

func main() {
	// Создаем приложение
	app := ui.NewApp()
	
	// Создаем главное окно
	window := app.NewWindow("Простое приложение")
	window.Resize(ui.NewSize(400, 300))
	window.Center()
	
	// Создаем компоненты
	label := ui.NewLabel("Добро пожаловать в наше приложение!")
	
	// Кнопка с обработчиком
	button := ui.NewButton("Нажми меня!", func() {
		ui.ShowInfo("Привет!", "Кнопка была нажата!", window)
	})
	
	// Кнопка с иконкой
	iconButton := ui.NewButtonWithIcon("Кнопка с иконкой", theme.InfoIcon(), func() {
		ui.ShowInfo("Иконка", "Кнопка с иконкой была нажата!", window)
	})
	
	// Слайдер
	slider := ui.NewSlider(0, 100, 50)
	slider.OnChanged = func(value float64) {
		fmt.Printf("Значение слайдера: %.2f\n", value)
	}
	
	// Чекбокс
	checkbox := ui.NewCheckBox("Согласен с условиями")
	checkbox.OnChanged = func(checked bool) {
		fmt.Printf("Чекбокс: %t\n", checked)
	}
	
	// Выпадающий список
	selectWidget := ui.NewSelect([]string{"Опция 1", "Опция 2", "Опция 3"})
	selectWidget.OnChanged = func(selected string) {
		fmt.Printf("Выбрано: %s\n", selected)
	}
	
	// Поле ввода
	entry := ui.NewEntry()
	entry.SetPlaceHolder("Введите текст...")
	entry.OnChanged = func(text string) {
		fmt.Printf("Введено: %s\n", text)
	}
	
	// Создаем компоновку
	content := ui.NewVBox(
		label,
		button,
		iconButton,
		ui.NewLabel("Слайдер:"),
		slider,
		checkbox,
		ui.NewLabel("Выберите опцию:"),
		selectWidget,
		ui.NewLabel("Введите текст:"),
		entry,
	)
	
	// Устанавливаем содержимое окна
	window.SetContent(content)
	
	// Показываем окно и запускаем приложение
	window.ShowAndRun()
}
