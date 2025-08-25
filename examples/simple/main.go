package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"go-ui-library"
)

func main() {
	// Создаем приложение
	app := ui.NewApp()
	
	// Создаем главное окно
	window := app.NewWindow("Простое приложение")
	window.Resize(fyne.NewSize(400, 300))
	window.Center()
	
	// Создаем компоненты
	label := ui.NewLabel("Добро пожаловать в наше приложение!")
	
	// Поле ввода с валидацией
	entry := ui.NewEntry()
	entry.SetPlaceHolder("Введите имя (мин 3 символа)...")
	
	// Кнопка сохранения (по умолчанию неактивна)
	saveBtn := ui.NewButtonWithIcon("Сохранить", theme.ConfirmIcon(), func() {
		ui.ShowInfo("Успех", fmt.Sprintf("Привет, %s!", entry.Text), window)
	})
	saveBtn.Disable()
	
	// Текст статуса внизу
	status := ui.NewLabel("")
	
	// Слайдер громкости с отображением значения
	volumeLabel := ui.NewLabel("Громкость: 50%")
	slider := ui.NewSlider(0, 100, 50)
	slider.OnChanged = func(value float64) {
		volumeLabel.SetText(fmt.Sprintf("Громкость: %d%%", int(value)))
	}
	
	// Чекбокс согласия
	agree := ui.NewCheckBox("Согласен с условиями")
	agree.OnChanged = func(checked bool) {
		// Кнопка доступна только при валидном вводе и согласии
		if checked && len(entry.Text) >= 3 {
			saveBtn.Enable()
			status.SetText("")
		} else {
			saveBtn.Disable()
			if !checked {
				status.SetText("Нужно принять условия")
			}
		}
	}
	
	// Обработка ввода: валидация длины
	entry.OnChanged = func(text string) {
		if len(text) < 3 {
			status.SetText("Имя слишком короткое (>=3)")
			saveBtn.Disable()
			return
		}
		if agree.Check.Checked { // если согласие уже проставлено
			saveBtn.Enable()
			status.SetText("")
		}
	}
	
	// Выпадающий список с реакцией
	selectWidget := ui.NewSelect([]string{"Опция 1", "Опция 2", "Опция 3"})
	selectWidget.OnChanged = func(selected string) {
		status.SetText(fmt.Sprintf("Выбрано: %s", selected))
	}
	
	// Дополнительная кнопка показа информации
	infoBtn := ui.NewButton("Показать диалог", func() {
		ui.ShowInfo("Информация", "Это демонстрация обработчиков событий", window)
	})
	
	// Компоновка
	content := ui.NewVBox(
		label,
		ui.NewLabel("Имя:"),
		entry,
		agree,
		saveBtn,
		infoBtn,
		ui.NewLabel("Настройка громкости:"),
		slider,
		volumeLabel,
		ui.NewLabel("Выберите опцию:"),
		selectWidget,
		status,
	)
	
	// Устанавливаем содержимое окна
	window.SetContent(content)
	
	// Показываем окно и запускаем приложение
	window.ShowAndRun()
}
