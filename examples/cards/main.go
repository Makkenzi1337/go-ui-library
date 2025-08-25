package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"go-ui-library"
)

func main() {
	// Создаем приложение
	app := ui.NewApp()
	
	// Создаем главное окно
	window := app.NewWindow("Приложение с карточками")
	window.Resize(fyne.NewSize(900, 700))
	window.Center()
	
	// Создаем карточки
	card1 := ui.NewCard("Карточка 1", "Описание первой карточки", 
		ui.NewVBox(
			ui.NewLabel("Это содержимое первой карточки"),
			ui.NewButton("Действие 1", func() {
				ui.ShowInfo("Карточка 1", "Выполнено действие 1", window)
			}),
		),
	)
	
	card2 := ui.NewCard("Карточка 2", "Описание второй карточки",
		ui.NewVBox(
			ui.NewLabel("Это содержимое второй карточки"),
			ui.NewSlider(0, 100, 50),
			ui.NewButton("Действие 2", func() {
				ui.ShowInfo("Карточка 2", "Выполнено действие 2", window)
			}),
		),
	)
	
	card3 := ui.NewCard("Карточка 3", "Описание третьей карточки",
		ui.NewVBox(
			ui.NewLabel("Это содержимое третьей карточки"),
			ui.NewCheckBox("Опция 1"),
			ui.NewCheckBox("Опция 2"),
			ui.NewButton("Действие 3", func() {
				ui.ShowInfo("Карточка 3", "Выполнено действие 3", window)
			}),
		),
	)
	
	// Создаем сетку карточек
	cardsGrid := ui.NewGrid(2, card1, card2, card3)
	
	// Создаем список элементов
	listItems := []string{"Элемент 1", "Элемент 2", "Элемент 3", "Элемент 4", "Элемент 5"}
	
	list := ui.NewList(
		func() int { return len(listItems) },
		func() fyne.CanvasObject {
			return ui.NewLabel("Template")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*ui.Label).SetText(listItems[id])
		},
	)
	
	list.OnSelected = func(index int) {
		ui.ShowInfo("Список", fmt.Sprintf("Выбран элемент: %s", listItems[index]), window)
	}
	
	// Создаем таблицу
	tableData := [][]string{
		{"Имя", "Возраст", "Город"},
		{"Иван", "25", "Москва"},
		{"Мария", "30", "Санкт-Петербург"},
		{"Петр", "35", "Новосибирск"},
		{"Анна", "28", "Екатеринбург"},
	}
	
	table := ui.NewTable(
		func() (int, int) { return len(tableData), len(tableData[0]) },
		func() fyne.CanvasObject {
			return ui.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*ui.Label).SetText(tableData[id.Row][id.Col])
		},
	)
	
	table.OnSelected = func(row, col int) {
		ui.ShowInfo("Таблица", fmt.Sprintf("Выбрана ячейка [%d,%d]: %s", row, col, tableData[row][col]), window)
	}
	
	// Создаем вкладки
	tabs := ui.NewTabs(
		container.NewTabItem("Карточки", cardsGrid),
		container.NewTabItem("Список", list),
		container.NewTabItem("Таблица", table),
	)
	
	// Создаем панель инструментов
	toolbar := ui.NewToolbar(
		ui.ToolbarAction(theme.ContentAddIcon(), "Добавить", func() {
			ui.ShowInfo("Панель инструментов", "Добавить новый элемент", window)
		}),
		ui.ToolbarAction(theme.DeleteIcon(), "Удалить", func() {
			ui.ShowInfo("Панель инструментов", "Удалить элемент", window)
		}),
		ui.ToolbarSeparator(),
		ui.ToolbarAction(theme.SearchIcon(), "Поиск", func() {
			ui.ShowInfo("Панель инструментов", "Поиск", window)
		}),
	)
	
	// Создаем строку состояния
	statusBar := ui.NewStatusBar("Готово к работе")
	
	// Создаем полную компоновку
	content := ui.CreateFullLayout(toolbar, tabs, statusBar)
	
	// Устанавливаем содержимое окна
	window.SetContent(content)
	
	// Показываем окно и запускаем приложение
	window.ShowAndRun()
}
