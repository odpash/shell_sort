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

func GetGraphicData(listSizeStr string, sortType string) {
	listSize, _ := strconv.Atoi(listSizeStr)
	var l *list.List
	switch sortType {
	case models.UNSORTED_LIST:
		l = createUnsortedList(listSize)
	case models.SORTED_LIST:
		l = createSortedList(listSize)
	case models.REVERSE_SORTED_LIST:
		l = createReverseSortedList(listSize)
	case models.PARTLY_SORTED_LIST:
		l = createPartlySortedList(listSize, 20)
	}
	printList(l)
}
