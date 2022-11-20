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
	"shell_sort/models"
	"strconv"
)

var selectedSortType string
var a fyne.App
var w fyne.Window

func validation(listSize string, sortType string, percent string) bool {
	if listSize == "" && sortType == "" {
		dialog.ShowInformation("Ошибка валидации данных!", "Вы не указали размер двусвязанного списка или тип сортировки!", w)
		return false
	}
	if _, err := strconv.Atoi(listSize); err != nil {
		dialog.ShowInformation("Ошибка валидации данных!", "Вы указали не целое число в поле размера двусвязанного списка!", w)
		return false
	}
	if _, err := strconv.Atoi(percent); err != nil {
		dialog.ShowInformation("Ошибка валидации данных!", "Вы указали не целое число в поле размера процента неупорядоченности", w)
		return false
	}
	return true
}

func getSortTypeLblContainer() *fyne.Container {
	selectSortTypeLbl := canvas.NewText("Выберите тип двусвязанного списка:", color.White)
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), selectSortTypeLbl, layout.NewSpacer())
}

func getSortTypeSelectContainer() *fyne.Container {
	return container.NewVBox(widget.NewSelect([]string{models.UNSORTED_LIST, models.SORTED_LIST, models.REVERSE_SORTED_LIST, models.PARTLY_SORTED_LIST}, func(value string) {
		log.Println("Выбран -> ", value)
		selectedSortType = value
	}))
}

func getListSizeContainer() *fyne.Container {
	inputSize := widget.NewEntry()
	inputSize.SetPlaceHolder("Введите размер двусвязанного списка...")
	inputPercent := widget.NewEntry()
	inputPercent.SetPlaceHolder("Введите процент упорядоченности двусвязанного списка...")

	return container.NewVBox(inputSize, inputPercent, widget.NewButton("Начать генерацию и сортировку", func() {
		if validation(inputSize.Text, selectedSortType, inputPercent.Text) {
			log.Println("All is OK!")
			controllers.GetGraphicData(inputSize.Text, selectedSortType)
		}
	}))
}

func CreateApplication() {
	a = app.New()
	w = a.NewWindow("Сортировка Шелла")
	w.SetFixedSize(true)
	w.SetContent(container.New(layout.NewVBoxLayout(), getSortTypeLblContainer(), getSortTypeSelectContainer(), getListSizeContainer()))
	w.Resize(fyne.NewSize(550, 200))
	w.ShowAndRun()
}

//
