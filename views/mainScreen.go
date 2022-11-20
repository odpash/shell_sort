package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"shell_sort/controllers"
	"strconv"
)

var selectedSortType string
var a fyne.App
var w fyne.Window

func validation(listSize string, sortType string) bool {
	if listSize == "" && sortType == "" {
		dialog.ShowInformation("Ошибка валидации данных!", "Вы не указали размер двусвязанного списка или тип сортировки!", w)
		return false
	}
	if _, err := strconv.Atoi(listSize); err != nil {
		dialog.ShowInformation("Ошибка валидации данных!", "Вы указали не целое число в поле размера двусвязанного спискка!", w)
		return false
	}
	return true
}

func getSortTypeLblContainer() *fyne.Container {
	selectSortTypeLbl := canvas.NewText("Выберите тип сортировки:", color.White)
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), selectSortTypeLbl, layout.NewSpacer())
}

func getSortTypeSelectContainer() *fyne.Container {
	return container.NewVBox(widget.NewSelect([]string{"Неупорядоченный массив", "Упорядоченный массив", "Обратно упорядоченный массив", "Частично упорядоченный массив"}, func(value string) {
		log.Println("Select set to", value)
		selectedSortType = value
	}))
}

func getListSizeContainer() *fyne.Container {
	input := widget.NewEntry()
	input.SetPlaceHolder("Введите размер двусвязанного списка...")
	return container.NewVBox(input, widget.NewButton("Начать генерацию и сортировку", func() {
		if validation(input.Text, selectedSortType) {
			log.Println("All is OK!")
			controllers.GetGraphicData(input.Text, selectedSortType)
		}
	}))
}

func CreateApplication() {
	a = app.New()
	w = a.NewWindow("Shell Sort")
	w.SetContent(container.New(layout.NewVBoxLayout(), getSortTypeLblContainer(), getSortTypeSelectContainer(), getListSizeContainer()))
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(500, 150))
	w.ShowAndRun()
}
