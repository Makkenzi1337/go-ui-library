package ui

import (
	"fmt"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// App представляет собой улучшенное приложение
type App struct {
	fyne.App
	windows []fyne.Window
}

// NewApp создает новое приложение
func NewApp() *App {
	return &App{
		App:     app.New(),
		windows: make([]fyne.Window, 0),
	}
}

// NewWindow создает новое окно
func (a *App) NewWindow(title string) *Window {
	window := &Window{
		Window: a.App.NewWindow(title),
	}
	a.windows = append(a.windows, window.Window)
	return window
}

// Run запускает приложение
func (a *App) Run() {
	a.App.Run()
}

// Window представляет собой улучшенное окно
type Window struct {
	fyne.Window
}

// SetContent устанавливает содержимое окна
func (w *Window) SetContent(content fyne.CanvasObject) {
	w.Window.SetContent(content)
}

// Show показывает окно
func (w *Window) Show() {
	w.Window.Show()
}

// ShowAndRun показывает окно и запускает приложение
func (w *Window) ShowAndRun() {
	w.Window.ShowAndRun()
}

// Resize изменяет размер окна
func (w *Window) Resize(size fyne.Size) {
	w.Window.Resize(size)
}

// Center центрирует окно на экране
func (w *Window) Center() {
	w.Window.CenterOnScreen()
}

// SetFullScreen устанавливает полноэкранный режим
func (w *Window) SetFullScreen(fullscreen bool) {
	w.Window.SetFullScreen(fullscreen)
}

// SetIcon устанавливает иконку окна
func (w *Window) SetIcon(icon fyne.Resource) {
	w.Window.SetIcon(icon)
}

// Dialog представляет собой диалоговое окно
type Dialog struct {
	dialog.Dialog
}

// ShowInfo показывает информационное диалоговое окно
func ShowInfo(title, message string, window fyne.Window) {
	dialog.ShowInformation(title, message, window)
}

// ShowError показывает диалоговое окно с ошибкой
func ShowError(err error, window fyne.Window) {
	dialog.ShowError(err, window)
}

// ShowConfirm показывает диалоговое окно подтверждения
func ShowConfirm(title, message string, callback func(bool), window fyne.Window) {
	dialog.ShowConfirm(title, message, callback, window)
}

// ShowInput показывает диалоговое окно ввода
func ShowInput(title, message string, callback func(string), window fyne.Window) {
	entry := widget.NewEntry()
	formItem := &widget.FormItem{
		Text:   "Введите текст:",
		Widget: entry,
	}
	dialog.ShowForm(title, "Подтвердить", "Отмена", []*widget.FormItem{formItem}, func(confirmed bool) {
		if confirmed {
			callback(entry.Text)
		}
	}, window)
}

// ShowFileOpen показывает диалоговое окно открытия файла
func ShowFileOpen(callback func(fyne.URI), window fyne.Window) {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			ShowError(err, window)
			return
		}
		if reader != nil {
			callback(reader.URI())
		}
	}, window)
}

// ShowFileSave показывает диалоговое окно сохранения файла
func ShowFileSave(callback func(fyne.URIWriteCloser), window fyne.Window) {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			ShowError(err, window)
			return
		}
		if writer != nil {
			callback(writer)
		}
	}, window)
}

// ShowFolderOpen показывает диалоговое окно открытия папки
func ShowFolderOpen(callback func(fyne.ListableURI), window fyne.Window) {
	dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
		if err != nil {
			ShowError(err, window)
			return
		}
		if list != nil {
			callback(list)
		}
	}, window)
}

// ProgressDialog представляет собой диалоговое окно с прогрессом
type ProgressDialog struct {
	*dialog.ProgressDialog
}

// ShowProgress показывает диалоговое окно с прогрессом
func ShowProgress(title, message string, window fyne.Window) *ProgressDialog {
	progress := dialog.NewProgress(title, message, window)
	progress.Show()
	return &ProgressDialog{
		ProgressDialog: progress,
	}
}

// SetProgress устанавливает значение прогресса
func (p *ProgressDialog) SetProgress(value float64) {
	p.ProgressDialog.SetValue(value)
}

// Hide скрывает диалоговое окно
func (p *ProgressDialog) Hide() {
	p.ProgressDialog.Hide()
}

// Layout представляет собой улучшенный макет
type Layout struct {
	fyne.Layout
}

// NewPaddedLayout создает контейнер с отступами
func NewPaddedLayout(objects ...fyne.CanvasObject) *Container {
	return &Container{
		Container: container.NewPadded(objects...),
	}
}

// NewMaxLayout создает контейнер с максимальным размером
func NewMaxLayout(objects ...fyne.CanvasObject) *Container {
	return &Container{
		Container: container.NewMax(objects...),
	}
}

// NewCenterLayout создает центрированный контейнер
func NewCenterLayout(objects ...fyne.CanvasObject) *Container {
	return &Container{
		Container: container.NewCenter(objects...),
	}
}

// Scroll представляет собой скроллируемый контейнер
type Scroll struct {
	*container.Scroll
}

// NewHScroll создает горизонтальный скролл
func NewHScroll(content fyne.CanvasObject) *Scroll {
	return &Scroll{
		Scroll: container.NewHScroll(content),
	}
}

// NewVScroll создает вертикальный скролл
func NewVScroll(content fyne.CanvasObject) *Scroll {
	return &Scroll{
		Scroll: container.NewVScroll(content),
	}
}

// NewBothScroll создает двунаправленный скролл
func NewBothScroll(content fyne.CanvasObject) *Scroll {
	return &Scroll{
		Scroll: container.NewScroll(content),
	}
}

// Theme представляет собой тему приложения
type Theme struct {
	fyne.Theme
}

// SetTheme устанавливает тему приложения
func SetTheme(theme fyne.Theme) {
	fyne.CurrentApp().Settings().SetTheme(theme)
}

// SetLightTheme устанавливает светлую тему
func SetLightTheme() {
	SetTheme(theme.LightTheme())
}

// SetDarkTheme устанавливает темную тему
func SetDarkTheme() {
	SetTheme(theme.DarkTheme())
}

// Color представляет собой цвет
type Color = color.Color

// NewColor создает новый цвет
func NewColor(r, g, b, a uint8) Color {
	return color.RGBA{r, g, b, a}
}

// NewRGBA создает новый цвет из RGBA значений
func NewRGBA(r, g, b, a uint8) Color {
	return color.RGBA{r, g, b, a}
}

// Common colors
var (
	ColorRed   = NewRGBA(255, 0, 0, 255)
	ColorGreen = NewRGBA(0, 255, 0, 255)
	ColorBlue  = NewRGBA(0, 0, 255, 255)
	ColorBlack = NewRGBA(0, 0, 0, 255)
	ColorWhite = NewRGBA(255, 255, 255, 255)
	ColorGray  = NewRGBA(128, 128, 128, 255)
)

// Size представляет собой размер
type Size struct {
	fyne.Size
}

// NewSize создает новый размер
func NewSize(width, height float32) Size {
	return Size{
		Size: fyne.NewSize(width, height),
	}
}

// Position представляет собой позицию
type Position struct {
	fyne.Position
}

// NewPosition создает новую позицию
func NewPosition(x, y float32) Position {
	return Position{
		Position: fyne.NewPos(x, y),
	}
}

// Helper функции для создания часто используемых компоновок

// CreateFormLayout создает компоновку формы
func CreateFormLayout(items ...fyne.CanvasObject) *fyne.Container {
	return container.NewVBox(items...)
}

// CreateToolbarLayout создает компоновку с панелью инструментов
func CreateToolbarLayout(toolbar fyne.CanvasObject, content fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(toolbar, nil, nil, nil, content)
}

// CreateStatusBarLayout создает компоновку со строкой состояния
func CreateStatusBarLayout(content fyne.CanvasObject, statusBar fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, statusBar, nil, nil, content)
}

// CreateFullLayout создает полную компоновку с панелью инструментов и строкой состояния
func CreateFullLayout(toolbar, content, statusBar fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(toolbar, statusBar, nil, nil, content)
}

// CreateSidebarLayout создает компоновку с боковой панелью
func CreateSidebarLayout(sidebar, content fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, sidebar, nil, content)
}

// Split представляет собой разделитель
type Split struct {
	*container.Split
}

// CreateSplitLayout создает компоновку с разделителем
func CreateSplitLayout(left, right fyne.CanvasObject) *Split {
	return &Split{
		Split: container.NewHSplit(left, right),
	}
}

// Utility функции

// FormatSize форматирует размер в читаемый вид
func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// FormatDuration форматирует длительность в читаемый вид
func FormatDuration(seconds int64) string {
	if seconds < 60 {
		return fmt.Sprintf("%ds", seconds)
	}
	if seconds < 3600 {
		return fmt.Sprintf("%dm %ds", seconds/60, seconds%60)
	}
	return fmt.Sprintf("%dh %dm", seconds/3600, (seconds%3600)/60)
}
