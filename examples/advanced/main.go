package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"go-ui-library/ui"
)

func main() {
	// Создаем приложение
	app := ui.NewApp()
	
	// Создаем главное окно
	window := app.NewWindow("Продвинутое приложение")
	window.Resize(ui.NewSize(800, 600))
	window.Center()
	
	// Создаем меню
	fileMenu := ui.Menu("Файл",
		ui.MenuItem("Новый", theme.DocumentCreateIcon(), func() {
			ui.ShowInfo("Файл", "Создать новый файл", window)
		}),
		ui.MenuItem("Открыть", theme.FolderOpenIcon(), func() {
			ui.ShowFileOpen(func(uri fyne.URI) {
				ui.ShowInfo("Файл", "Открыт файл: "+uri.Name(), window)
			}, window)
		}),
		ui.MenuItem("Сохранить", theme.DocumentSaveIcon(), func() {
			ui.ShowFileSave(func(writer fyne.URIWriteCloser) {
				ui.ShowInfo("Файл", "Файл сохранен", window)
			}, window)
		}),
		ui.MenuItemSeparator(),
		ui.MenuItem("Выход", theme.CancelIcon(), func() {
			window.Close()
		}),
	)
	
	editMenu := ui.Menu("Правка",
		ui.MenuItem("Вырезать", theme.ContentCutIcon(), func() {
			ui.ShowInfo("Правка", "Вырезать", window)
		}),
		ui.MenuItem("Копировать", theme.ContentCopyIcon(), func() {
			ui.ShowInfo("Правка", "Копировать", window)
		}),
		ui.MenuItem("Вставить", theme.ContentPasteIcon(), func() {
			ui.ShowInfo("Правка", "Вставить", window)
		}),
	)
	
	helpMenu := ui.Menu("Помощь",
		ui.MenuItem("О программе", theme.HelpIcon(), func() {
			ui.ShowInfo("О программе", "Продвинутое приложение v1.0\nСоздано с помощью Go UI Library", window)
		}),
	)
	
	menuBar := ui.NewMenuBar(fileMenu, editMenu, helpMenu)
	
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
		ui.NewFormItem("Имя:", ui.NewEntry()),
		ui.NewFormItem("Email:", ui.NewEntry()),
		ui.NewFormItem("Возраст:", ui.NewEntry()),
		ui.NewFormItem("Пол:", ui.NewSelect([]string{"Мужской", "Женский"})),
	)
	
	// Создаем вкладки
	tabs := ui.NewTabs(
		ui.NewTabItem("Основная", mainContent),
		ui.NewTabItem("Настройки", settingsContent),
		ui.NewTabItem("Форма", formContent),
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
