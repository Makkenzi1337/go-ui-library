// Package ui предоставляет простой и современный API для создания десктопных приложений на Go
package ui

// Экспортируем все компоненты из внутреннего пакета
import "go-ui-library/internal/ui"

// Re-export всех типов и функций для удобства использования

// App types
type App = ui.App
type Window = ui.Window

// Basic components
type Button = ui.Button
type Label = ui.Label
type Entry = ui.Entry
type CheckBox = ui.CheckBox
type Slider = ui.Slider
type Select = ui.Select
type ProgressBar = ui.ProgressBar

// Container types
type Container = ui.Container
type Card = ui.Card
type List = ui.List
type Table = ui.Table
type Form = ui.Form
type FormItem = ui.FormItem
type Tabs = ui.Tabs
type TabItem = ui.TabItem

// Layout types
type SplitContainer = ui.SplitContainer
type Layout = ui.Layout

// Toolbar and menu types
type Toolbar = ui.Toolbar
type MenuBar = ui.MenuBar

// Dialog types
type Dialog = ui.Dialog
type ProgressDialog = ui.ProgressDialog

// Utility types
type Color = ui.Color
type Size = ui.Size
type Position = ui.Position

// Re-export constructors
var (
	// App constructors
	NewApp = ui.NewApp

	// Basic component constructors
	NewButton        = ui.NewButton
	NewButtonWithIcon = ui.NewButtonWithIcon
	NewLabel         = ui.NewLabel
	NewEntry         = ui.NewEntry
	NewCheckBox      = ui.NewCheckBox
	NewSlider        = ui.NewSlider
	NewSelect        = ui.NewSelect
	NewProgressBar   = ui.NewProgressBar

	// Container constructors
	NewVBox    = ui.NewVBox
	NewHBox    = ui.NewHBox
	NewGrid    = ui.NewGrid
	NewBorder  = ui.NewBorder
	NewCard    = ui.NewCard
	NewList    = ui.NewList
	NewTable   = ui.NewTable
	NewForm    = ui.NewForm
	NewFormItem = ui.NewFormItem
	NewTabs    = ui.NewTabs
	NewTabItem = ui.NewTabItem
	NewTabItemWithIcon = ui.NewTabItemWithIcon

	// Split container constructors
	NewHSplit = ui.NewHSplit
	NewVSplit = ui.NewVSplit

	// Toolbar constructors
	NewToolbar        = ui.NewToolbar
	ToolbarAction     = ui.ToolbarAction
	ToolbarSeparator  = ui.ToolbarSeparator

	// Menu constructors
	NewMenuBar        = ui.NewMenuBar
	Menu              = ui.Menu
	MenuItem          = ui.MenuItem
	MenuItemSeparator = ui.MenuItemSeparator
	MenuItemCheck     = ui.MenuItemCheck
	MenuItemRadio     = ui.MenuItemRadio
	MenuItemSubMenu   = ui.MenuItemSubMenu

	// Status bar constructors
	NewStatusBar = ui.NewStatusBar

	// Layout constructors
	NewPaddedLayout = ui.NewPaddedLayout
	NewMaxLayout    = ui.NewMaxLayout
	NewCenterLayout = ui.NewCenterLayout

	// Scroll constructors
	NewHScroll   = ui.NewHScroll
	NewVScroll   = ui.NewVScroll
	NewBothScroll = ui.NewBothScroll

	// Color constructors
	NewColor = ui.NewColor
	NewRGBA  = ui.NewRGBA

	// Size and position constructors
	NewSize     = ui.NewSize
	NewPosition = ui.NewPosition

	// Layout helper constructors
	CreateFormLayout      = ui.CreateFormLayout
	CreateToolbarLayout   = ui.CreateToolbarLayout
	CreateStatusBarLayout = ui.CreateStatusBarLayout
	CreateFullLayout      = ui.CreateFullLayout
	CreateSidebarLayout   = ui.CreateSidebarLayout
	CreateSplitLayout     = ui.CreateSplitLayout

	// Dialog functions
	ShowInfo      = ui.ShowInfo
	ShowError     = ui.ShowError
	ShowConfirm   = ui.ShowConfirm
	ShowInput     = ui.ShowInput
	ShowFileOpen  = ui.ShowFileOpen
	ShowFileSave  = ui.ShowFileSave
	ShowFolderOpen = ui.ShowFolderOpen
	ShowProgress  = ui.ShowProgress

	// Theme functions
	SetTheme      = ui.SetTheme
	SetLightTheme = ui.SetLightTheme
	SetDarkTheme  = ui.SetDarkTheme

	// Utility functions
	FormatSize     = ui.FormatSize
	FormatDuration = ui.FormatDuration
)

// Common colors
var (
	ColorRed   = ui.ColorRed
	ColorGreen = ui.ColorGreen
	ColorBlue  = ui.ColorBlue
	ColorBlack = ui.ColorBlack
	ColorWhite = ui.ColorWhite
	ColorGray  = ui.ColorGray
)
