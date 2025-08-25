package main

import (
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
	
	// Создаем содержимое для вкладок
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
	settingsContent := ui.NewVBox(
		ui.NewLabel("Настройки"),
		ui.NewCheckBox("Включить уведомления"),
		ui.NewCheckBox("Автосохранение"),
		ui.NewSelect([]string{"Светлая тема", "Темная тема"}),
		ui.NewButton("Применить настройки", func() {
			ui.ShowInfo("Настройки", "Настройки применены", window)
		}),
	)
	
	// Вкладка 3: Форма
	formContent := ui.NewForm(
		widget.NewFormItem("Имя:", ui.NewEntry()),
		widget.NewFormItem("Email:", ui.NewEntry()),
		widget.NewFormItem("Возраст:", ui.NewEntry()),
		widget.NewFormItem("Пол:", ui.NewSelect([]string{"Мужской", "Женский"})),
	)
	
	// Создаем вкладки
	tabs := ui.NewTabs(
		container.NewTabItem("Основная", mainContent),
		container.NewTabItem("Настройки", settingsContent),
		container.NewTabItem("Форма", formContent),
	)
	
	// Создаем строку состояния
	statusBar := ui.NewStatusBar("Готово")
	
	// Создаем полную компоновку
	content := ui.CreateFullLayout(toolbar, tabs, statusBar)
	
	// Устанавливаем содержимое окна
	window.SetContent(content)
	
	// Показываем окно и запускаем приложение
	window.ShowAndRun()
}
