package ui

import (
	"testing"
	"fyne.io/fyne/v2/test"
)

func TestNewButton(t *testing.T) {
	clicked := false
	button := NewButton("Test Button", func() {
		clicked = true
	})
	
	if button.Text != "Test Button" {
		t.Errorf("Expected button text to be 'Test Button', got '%s'", button.Text)
	}
	
	// Симулируем клик
	test.Tap(button)
	
	if !clicked {
		t.Error("Button click handler was not called")
	}
}

func TestNewLabel(t *testing.T) {
	label := NewLabel("Test Label")
	
	if label.Text != "Test Label" {
		t.Errorf("Expected label text to be 'Test Label', got '%s'", label.Text)
	}
}

func TestNewEntry(t *testing.T) {
	entry := NewEntry()
	entry.SetPlaceHolder("Test Placeholder")
	
	if entry.PlaceHolder != "Test Placeholder" {
		t.Errorf("Expected placeholder to be 'Test Placeholder', got '%s'", entry.PlaceHolder)
	}
	
	// Тестируем обработчик изменений
	changed := false
	entry.OnChanged = func(text string) {
		changed = true
	}
	
	test.Type(entry, "test")
	
	if !changed {
		t.Error("Entry OnChanged handler was not called")
	}
}

func TestNewSlider(t *testing.T) {
	slider := NewSlider(0, 100, 50)
	
	if slider.Value != 50 {
		t.Errorf("Expected slider value to be 50, got %f", slider.Value)
	}
	
	// Тестируем обработчик изменений
	changed := false
	slider.OnChanged = func(value float64) {
		changed = true
	}
	
	slider.SetValue(75)
	
	if !changed {
		t.Error("Slider OnChanged handler was not called")
	}
}

func TestNewSelect(t *testing.T) {
	options := []string{"Option 1", "Option 2", "Option 3"}
	selectWidget := NewSelect(options)
	
	if len(selectWidget.Options) != 3 {
		t.Errorf("Expected 3 options, got %d", len(selectWidget.Options))
	}
	
	// Тестируем обработчик изменений
	changed := false
	selectWidget.OnChanged = func(selected string) {
		changed = true
	}
	
	selectWidget.SetSelected("Option 2")
	
	if !changed {
		t.Error("Select OnChanged handler was not called")
	}
}

func TestNewCheckBox(t *testing.T) {
	checkbox := NewCheckBox("Test Checkbox")
	
	if checkbox.Text != "Test Checkbox" {
		t.Errorf("Expected checkbox text to be 'Test Checkbox', got '%s'", checkbox.Text)
	}
	
	// Тестируем обработчик изменений
	changed := false
	checkbox.OnChanged = func(checked bool) {
		changed = true
	}
	
	checkbox.SetChecked(true)
	
	if !changed {
		t.Error("CheckBox OnChanged handler was not called")
	}
}

func TestNewProgressBar(t *testing.T) {
	progress := NewProgressBar()
	
	progress.SetValue(0.5)
	
	if progress.Value != 0.5 {
		t.Errorf("Expected progress value to be 0.5, got %f", progress.Value)
	}
}

func TestNewVBox(t *testing.T) {
	label1 := NewLabel("Label 1")
	label2 := NewLabel("Label 2")
	
	container := NewVBox(label1, label2)
	
	if len(container.Objects) != 2 {
		t.Errorf("Expected 2 objects in container, got %d", len(container.Objects))
	}
}

func TestNewHBox(t *testing.T) {
	label1 := NewLabel("Label 1")
	label2 := NewLabel("Label 2")
	
	container := NewHBox(label1, label2)
	
	if len(container.Objects) != 2 {
		t.Errorf("Expected 2 objects in container, got %d", len(container.Objects))
	}
}

func TestNewGrid(t *testing.T) {
	label1 := NewLabel("Label 1")
	label2 := NewLabel("Label 2")
	label3 := NewLabel("Label 3")
	label4 := NewLabel("Label 4")
	
	container := NewGrid(2, label1, label2, label3, label4)
	
	if len(container.Objects) != 4 {
		t.Errorf("Expected 4 objects in container, got %d", len(container.Objects))
	}
}

func TestNewCard(t *testing.T) {
	content := NewLabel("Card Content")
	card := NewCard("Test Title", "Test Subtitle", content)
	
	if card.Title != "Test Title" {
		t.Errorf("Expected card title to be 'Test Title', got '%s'", card.Title)
	}
	
	if card.Subtitle != "Test Subtitle" {
		t.Errorf("Expected card subtitle to be 'Test Subtitle', got '%s'", card.Subtitle)
	}
}

func TestColorConstructors(t *testing.T) {
	color1 := NewColor(1.0, 0.0, 0.0, 1.0)
	color2 := NewRGBA(255, 0, 0, 255)
	
	if color1.R != 1.0 || color1.G != 0.0 || color1.B != 0.0 || color1.A != 1.0 {
		t.Error("NewColor did not create correct color values")
	}
	
	if color2.R != 1.0 || color2.G != 0.0 || color2.B != 0.0 || color2.A != 1.0 {
		t.Error("NewRGBA did not create correct color values")
	}
}

func TestSizeAndPosition(t *testing.T) {
	size := NewSize(100, 200)
	pos := NewPosition(50, 75)
	
	if size.Width != 100 || size.Height != 200 {
		t.Error("NewSize did not create correct size values")
	}
	
	if pos.X != 50 || pos.Y != 75 {
		t.Error("NewPosition did not create correct position values")
	}
}

func TestUtilityFunctions(t *testing.T) {
	// Тест форматирования размера
	sizeStr := FormatSize(1024)
	if sizeStr != "1.0 KB" {
		t.Errorf("Expected '1.0 KB', got '%s'", sizeStr)
	}
	
	// Тест форматирования длительности
	durationStr := FormatDuration(65)
	if durationStr != "1m 5s" {
		t.Errorf("Expected '1m 5s', got '%s'", durationStr)
	}
}
