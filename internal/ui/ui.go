package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Button представляет собой улучшенную кнопку с простым API
type Button struct {
	*widget.Button
	onClick func()
}

// NewButton создает новую кнопку с обработчиком клика
func NewButton(text string, onClick func()) *Button {
	btn := &Button{
		Button:  widget.NewButton(text, onClick),
		onClick: onClick,
	}
	
	// Применяем современные стили
	btn.Importance = widget.HighImportance
	
	return btn
}

// NewButtonWithIcon создает кнопку с иконкой
func NewButtonWithIcon(text string, icon fyne.Resource, onClick func()) *Button {
	btn := &Button{
		Button:  widget.NewButtonWithIcon(text, icon, onClick),
		onClick: onClick,
	}
	
	btn.Importance = widget.HighImportance
	
	return btn
}

// Slider представляет собой улучшенный слайдер
type Slider struct {
	*widget.Slider
	OnChanged func(float64)
}

// NewSlider создает новый слайдер с диапазоном значений
func NewSlider(min, max, value float64) *Slider {
	slider := &Slider{
		Slider: widget.NewSlider(min, max),
	}
	
	slider.SetValue(value)
	slider.OnChanged = func(v float64) {
		if slider.OnChanged != nil {
			slider.OnChanged(v)
		}
	}
	
	return slider
}

// Select представляет собой улучшенный выпадающий список
type Select struct {
	*widget.Select
	OnChanged func(string)
}

// NewSelect создает новый выпадающий список
func NewSelect(options []string) *Select {
	selectWidget := &Select{
		Select: widget.NewSelect(options, nil),
	}
	
	selectWidget.OnChanged = func(selected string) {
		if selectWidget.OnChanged != nil {
			selectWidget.OnChanged(selected)
		}
	}
	
	return selectWidget
}

// Entry представляет собой улучшенное поле ввода
type Entry struct {
	*widget.Entry
	OnChanged func(string)
}

// NewEntry создает новое поле ввода
func NewEntry() *Entry {
	entry := &Entry{
		Entry: widget.NewEntry(),
	}
	
	entry.OnChanged = func(text string) {
		if entry.OnChanged != nil {
			entry.OnChanged(text)
		}
	}
	
	return entry
}

// SetPlaceHolder устанавливает placeholder для поля ввода
func (e *Entry) SetPlaceHolder(placeholder string) {
	e.Entry.PlaceHolder = placeholder
}

// CheckBox представляет собой улучшенный чекбокс
type CheckBox struct {
	*widget.Check
	OnChanged func(bool)
}

// NewCheckBox создает новый чекбокс
func NewCheckBox(text string) *CheckBox {
	checkbox := &CheckBox{
		Check: widget.NewCheck(text, nil),
	}
	
	checkbox.OnChanged = func(checked bool) {
		if checkbox.OnChanged != nil {
			checkbox.OnChanged(checked)
		}
	}
	
	return checkbox
}

// ProgressBar представляет собой улучшенный прогресс-бар
type ProgressBar struct {
	*widget.ProgressBar
}

// NewProgressBar создает новый прогресс-бар
func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		ProgressBar: widget.NewProgressBar(),
	}
}

// SetValue устанавливает значение прогресс-бара
func (p *ProgressBar) SetValue(value float64) {
	p.ProgressBar.SetValue(value)
}

// Label представляет собой улучшенную метку
type Label struct {
	*widget.Label
}

// NewLabel создает новую метку
func NewLabel(text string) *Label {
	return &Label{
		Label: widget.NewLabel(text),
	}
}

// SetText устанавливает текст метки
func (l *Label) SetText(text string) {
	l.Label.SetText(text)
}

// Container представляет собой улучшенный контейнер
type Container struct {
	*fyne.Container
}

// NewVBox создает вертикальный контейнер
func NewVBox(objects ...fyne.CanvasObject) *Container {
	return &Container{
		Container: container.NewVBox(objects...),
	}
}

// NewHBox создает горизонтальный контейнер
func NewHBox(objects ...fyne.CanvasObject) *Container {
	return &Container{
		Container: container.NewHBox(objects...),
	}
}

// NewGrid создает сеточный контейнер
func NewGrid(cols int, objects ...fyne.CanvasObject) *Container {
	return &Container{
		Container: container.NewGridWithColumns(cols, objects...),
	}
}

// NewBorder создает контейнер с границами
func NewBorder(top, bottom, left, right fyne.CanvasObject, center fyne.CanvasObject) *Container {
	return &Container{
		Container: container.NewBorder(top, bottom, left, right, center),
	}
}
