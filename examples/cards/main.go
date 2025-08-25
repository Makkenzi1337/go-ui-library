package main

import (
	"fmt"
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
	window := app.NewWindow("Приложение с карточками")
	window.Resize(fyne.NewSize(900, 700))
	window.Center()
	
	status := ui.NewStatusBar("Готово")
	
	// Карточки с действиями
	card1 := ui.NewCard("Карточка 1", "Описание первой карточки",
		ui.NewVBox(
			ui.NewLabel("Это содержимое первой карточки"),
			ui.NewButton("Действие 1", func() {
				status.SetText("К1: действие 1 выполнено")
				ui.ShowInfo("Карточка 1", "Выполнено действие 1", window)
			}),
		),
	)

	card2Slider := ui.NewSlider(0, 100, 50)
	card2Value := ui.NewLabel("Значение: 50")
	card2Slider.OnChanged = func(v float64) { card2Value.SetText(fmt.Sprintf("Значение: %d", int(v))) }
	card2 := ui.NewCard("Карточка 2", "Описание второй карточки",
		ui.NewVBox(
			ui.NewLabel("Это содержимое второй карточки"),
			card2Slider,
			card2Value,
			ui.NewButton("Действие 2", func() {
				status.SetText("К2: действие 2 выполнено")
				ui.ShowInfo("Карточка 2", "Выполнено действие 2", window)
			}),
		),
	)

	opt1 := ui.NewCheckBox("Опция 1")
	opt2 := ui.NewCheckBox("Опция 2")
	card3 := ui.NewCard("Карточка 3", "Описание третьей карточки",
		ui.NewVBox(
			ui.NewLabel("Это содержимое третьей карточки"),
			opt1,
			opt2,
			ui.NewButton("Собрать", func() {
				status.SetText(fmt.Sprintf("К3: %v, %v", opt1.Check.Checked, opt2.Check.Checked))
			}),
		),
	)

	cardsGrid := ui.NewGrid(2, card1, card2, card3)
	
	// Список с обработчиками
	listItems := []string{"Элемент 1", "Элемент 2", "Элемент 3", "Элемент 4", "Элемент 5"}
	list := ui.NewList(
		func() int { return len(listItems) },
		func() fyne.CanvasObject { return ui.NewLabel("Template") },
		func(id widget.ListItemID, obj fyne.CanvasObject) { obj.(*ui.Label).SetText(listItems[id]) },
	)
	list.OnSelected = func(index int) {
		status.SetText(fmt.Sprintf("Список: выбран %s", listItems[index]))
	}

	// Таблица с обработчиками
	tableData := [][]string{
		{"Имя", "Возраст", "Город"},
		{"Иван", "25", "Москва"},
		{"Мария", "30", "Санкт-Петербург"},
		{"Петр", "35", "Новосибирск"},
		{"Анна", "28", "Екатеринбург"},
	}
	table := ui.NewTable(
		func() (int, int) { return len(tableData), len(tableData[0]) },
		func() fyne.CanvasObject { return ui.NewLabel("Template") },
		func(id widget.TableCellID, obj fyne.CanvasObject) { obj.(*ui.Label).SetText(tableData[id.Row][id.Col]) },
	)
	table.OnSelected = func(row, col int) {
		status.SetText(fmt.Sprintf("Таблица: [%d,%d] = %s", row, col, tableData[row][col]))
	}

	// Вкладки
	tabs := ui.NewTabs(
		container.NewTabItem("Карточки", cardsGrid),
		container.NewTabItem("Список", list),
		container.NewTabItem("Таблица", table),
	)

	// Тулбар
	toolbar := ui.NewToolbar(
		ui.ToolbarAction(theme.ContentAddIcon(), "Добавить", func() {
			status.SetText("Добавление элемента")
		}),
		ui.ToolbarAction(theme.DeleteIcon(), "Удалить", func() {
			status.SetText("Удаление элемента")
		}),
		ui.ToolbarSeparator(),
		ui.ToolbarAction(theme.SearchIcon(), "Поиск", func() {
			status.SetText("Поиск...")
		}),
	)

	content := ui.CreateFullLayout(toolbar, tabs, status)
	window.SetContent(content)
	window.ShowAndRun()
}
