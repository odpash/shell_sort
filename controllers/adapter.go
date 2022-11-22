package controllers

import (
	"container/list"
	"fmt"
	"shell_sort/models"
	"strconv"
)

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func GetGraphicData(listSizeStr string, sortType string, percentStr string) {
	listSize, _ := strconv.Atoi(listSizeStr)
	disorderPercent, _ := strconv.Atoi(percentStr)
	var l *list.List
	var c int
	switch sortType {
	case models.UNSORTED_LIST:
		l = createUnsortedList(listSize)
	case models.SORTED_LIST:
		l = createSortedList(listSize)
	case models.REVERSE_SORTED_LIST:
		l = createReverseSortedList(listSize)
	case models.PARTLY_SORTED_LIST:
		l, c = createPartlySortedList(listSize, disorderPercent)
	}
	printList(l)
	fmt.Println(c)
}
