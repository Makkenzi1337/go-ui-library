package main

import (
	"fmt"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go-ui-library"
)

func main() {
	// Создаем приложение
	app := ui.NewApp()
	
	// Создаем главное окно
	window := app.NewWindow("Продвинутое приложение")
	window.Resize(fyne.NewSize(800, 600))
	window.Center()
	
	// Создаем панель инструментов
	toolbar := ui.NewToolbar(
		ui.ToolbarAction(theme.DocumentCreateIcon(), "Новый", func() {
			ui.ShowInfo("Панель инструментов", "Создать новый", window)
		}),
		ui.ToolbarAction(theme.FolderOpenIcon(), "Открыть", func() {
			ui.ShowInfo("Панель инструментов", "Открыть файл", window)
		}),
		ui.ToolbarAction(theme.DocumentSaveIcon(), "Сохранить", func() {
			ui.ShowInfo("Панель инструментов", "Сохранить файл", window)
		}),
		ui.ToolbarSeparator(),
		ui.ToolbarAction(theme.ContentCutIcon(), "Вырезать", func() {
			ui.ShowInfo("Панель инструментов", "Вырезать", window)
		}),
		ui.ToolbarAction(theme.ContentCopyIcon(), "Копировать", func() {
			ui.ShowInfo("Панель инструментов", "Копировать", window)
		}),
		ui.ToolbarAction(theme.ContentPasteIcon(), "Вставить", func() {
			ui.ShowInfo("Панель инструментов", "Вставить", window)
		}),
	)
	
	// Вкладка 1: Основная информация
	mainContent := ui.NewVBox(
		ui.NewLabel("Основная информация"),
		ui.NewButton("Показать информацию", func() {
			ui.ShowInfo("Информация", "Это основная вкладка приложения", window)
		}),
		ui.NewProgressBar(),
		ui.NewSlider(0, 100, 25),
	)
	
	// Вкладка 2: Настройки
	notify := ui.NewCheckBox("Включить уведомления")
	autosave := ui.NewCheckBox("Автосохранение")
	themeSelect := ui.NewSelect([]string{"Светлая тема", "Темная тема"})
	applyBtn := ui.NewButton("Применить настройки", func() {
		mode := themeSelect.Selected
		if mode == "Темная тема" {
			ui.SetDarkTheme()
		} else {
			ui.SetLightTheme()
		}
		ui.ShowInfo("Настройки", "Настройки применены", window)
	})
	settingsContent := ui.NewVBox(
		ui.NewLabel("Настройки"),
		notify,
		autosave,
		themeSelect,
		applyBtn,
	)
	
	// Вкладка 3: Форма с валидацией
	nameEntry := ui.NewEntry()
	emailEntry := ui.NewEntry()
	ageEntry := ui.NewEntry()
	genderSelect := ui.NewSelect([]string{"Мужской", "Женский"})
	form := ui.NewForm(
		widget.NewFormItem("Имя:", nameEntry),
		widget.NewFormItem("Email:", emailEntry),
		widget.NewFormItem("Возраст:", ageEntry),
		widget.NewFormItem("Пол:", genderSelect),
	)
	submitBtn := ui.NewButton("Отправить", func() {
		if len(nameEntry.Text) < 3 {
			ui.ShowError(fmt.Errorf("Имя слишком короткое"), window)
			return
		}
		if !strings.Contains(emailEntry.Text, "@") {
			ui.ShowError(fmt.Errorf("Некорректный email"), window)
			return
		}
		ui.ShowInfo("Форма", "Форма успешно отправлена", window)
	})
	formContent := ui.NewVBox(form, submitBtn)
	
	// Создаем вкладки
	tabs := ui.NewTabs(
		container.NewTabItem("Основная", mainContent),
		container.NewTabItem("Настройки", settingsContent),
		container.NewTabItem("Форма", formContent),
	)
	
	// Полная компоновка с тулбаром и статусом
	status := ui.NewStatusBar("Готово к работе")
	content := ui.CreateFullLayout(toolbar, tabs, status)
	
	// Устанавливаем содержимое окна
	window.SetContent(content)
	
	// Показываем окно и запускаем приложение
	window.ShowAndRun()
}
