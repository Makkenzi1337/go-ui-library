package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Card представляет собой карточку с заголовком и содержимым
type Card struct {
	*widget.Card
}

// NewCard создает новую карточку
func NewCard(title, subtitle string, content fyne.CanvasObject) *Card {
	return &Card{
		Card: widget.NewCard(title, subtitle, content),
	}
}

// List представляет собой улучшенный список
type List struct {
	*widget.List
	OnSelected func(int)
}

// NewList создает новый список
func NewList(length func() int, createItem func() fyne.CanvasObject, updateItem func(widget.ListItemID, fyne.CanvasObject)) *List {
	list := &List{
		List: widget.NewList(length, createItem, updateItem),
	}
	
	list.OnSelected = func(id widget.ListItemID) {
		if list.OnSelected != nil {
			list.OnSelected(int(id))
		}
	}
	
	return list
}

// Table представляет собой улучшенную таблицу
type Table struct {
	*widget.Table
	OnSelected func(int, int)
}

// NewTable создает новую таблицу
func NewTable(length func() (int, int), createCell func() fyne.CanvasObject, updateCell func(widget.TableCellID, fyne.CanvasObject)) *Table {
	table := &Table{
		Table: widget.NewTable(length, createCell, updateCell),
	}
	
	return table
}

// Form представляет собой улучшенную форму
type Form struct {
	*widget.Form
}

// NewForm создает новую форму
func NewForm(items ...*widget.FormItem) *Form {
	return &Form{
		Form: widget.NewForm(items...),
	}
}

// FormItem представляет собой элемент формы
type FormItem struct {
	Text   string
	Widget fyne.CanvasObject
}

// NewFormItem создает новый элемент формы
func NewFormItem(text string, widget fyne.CanvasObject) *FormItem {
	return &FormItem{
		Text:   text,
		Widget: widget,
	}
}

// Tabs представляет собой улучшенные вкладки
type Tabs struct {
	*container.AppTabs
}

// NewTabs создает новые вкладки
func NewTabs(items ...*container.TabItem) *Tabs {
	return &Tabs{
		AppTabs: container.NewAppTabs(items...),
	}
}

// TabItem представляет собой вкладку
type TabItem struct {
	*container.TabItem
}

// NewTabItem создает новую вкладку
func NewTabItem(text string, content fyne.CanvasObject) *TabItem {
	return &TabItem{
		TabItem: container.NewTabItem(text, content),
	}
}

// NewTabItemWithIcon создает новую вкладку с иконкой
func NewTabItemWithIcon(text string, icon fyne.Resource, content fyne.CanvasObject) *TabItem {
	return &TabItem{
		TabItem: container.NewTabItemWithIcon(text, icon, content),
	}
}

// Toolbar представляет собой улучшенную панель инструментов
type Toolbar struct {
	*widget.Toolbar
}

// NewToolbar создает новую панель инструментов
func NewToolbar(items ...widget.ToolbarItem) *Toolbar {
	return &Toolbar{
		Toolbar: widget.NewToolbar(items...),
	}
}

// ToolbarAction создает действие для панели инструментов
func ToolbarAction(icon fyne.Resource, tooltip string, onTapped func()) widget.ToolbarItem {
	return widget.NewToolbarAction(icon, onTapped)
}

// ToolbarSeparator создает разделитель для панели инструментов
func ToolbarSeparator() widget.ToolbarItem {
	return widget.NewToolbarSeparator()
}

// MenuBar представляет собой улучшенное меню
type MenuBar struct {
	*fyne.MainMenu
}

// NewMenuBar создает новое меню
func NewMenuBar(menus ...*fyne.Menu) *MenuBar {
	return &MenuBar{
		MainMenu: fyne.NewMainMenu(menus...),
	}
}

// Menu создает меню
func Menu(label string, items ...*fyne.MenuItem) *fyne.Menu {
	return fyne.NewMenu(label, items...)
}

// MenuItem создает элемент меню
func MenuItem(label string, icon fyne.Resource, action func()) *fyne.MenuItem {
	return fyne.NewMenuItem(label, action)
}

// MenuItemSeparator создает разделитель в меню
func MenuItemSeparator() *fyne.MenuItem {
	return fyne.NewMenuItemSeparator()
}

// MenuItemCheck создает элемент меню с чекбоксом
func MenuItemCheck(label string, icon fyne.Resource, checked bool, action func(bool)) *fyne.MenuItem {
	return fyne.NewMenuItem(label, func() {
		action(!checked)
	})
}

// MenuItemRadio создает элемент меню с радиокнопкой
func MenuItemRadio(label string, icon fyne.Resource, checked bool, action func(bool)) *fyne.MenuItem {
	return fyne.NewMenuItem(label, func() {
		action(!checked)
	})
}

// MenuItemSubMenu создает подменю
func MenuItemSubMenu(label string, icon fyne.Resource, subMenu *fyne.Menu) *fyne.MenuItem {
	return fyne.NewMenuItem(label, nil)
}

// StatusBar представляет собой строку состояния
type StatusBar struct {
	*widget.Label
}

// NewStatusBar создает новую строку состояния
func NewStatusBar(text string) *StatusBar {
	return &StatusBar{
		Label: widget.NewLabel(text),
	}
}

// SetStatus устанавливает текст строки состояния
func (s *StatusBar) SetStatus(text string) {
	s.Label.SetText(text)
}

// SplitContainer представляет собой контейнер с разделителем
type SplitContainer struct {
	*container.Split
}

// NewHSplit создает горизонтальный разделенный контейнер
func NewHSplit(leading, trailing fyne.CanvasObject) *SplitContainer {
	return &SplitContainer{
		Split: container.NewHSplit(leading, trailing),
	}
}

// NewVSplit создает вертикальный разделенный контейнер
func NewVSplit(top, bottom fyne.CanvasObject) *SplitContainer {
	return &SplitContainer{
		Split: container.NewVSplit(top, bottom),
	}
}

// SetOffset устанавливает позицию разделителя
func (s *SplitContainer) SetOffset(offset float64) {
	s.Split.SetOffset(offset)
}

// GetOffset возвращает позицию разделителя
func (s *SplitContainer) GetOffset() float64 {
	return s.Split.Offset
}
