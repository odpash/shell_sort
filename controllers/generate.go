package controllers

import (
	"container/list"
	"math/rand"
	"shell_sort/models"
)

func contains(s []int, element int) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}

	return false
}

func createUnsortedList(listSize int) *list.List {
	linkedList := list.New()
	for i := 0; i < listSize; i++ {
		linkedList.PushFront(rand.Intn(listSize * 7))
	}
	return linkedList
}

func createSortedList(listSize int) *list.List {
	linkedList := list.New()
	step := 10
	lastElement := 0
	for i := 0; i < listSize; i++ {
		currentElement := rand.Intn(lastElement + step)
		for currentElement <= lastElement {
			currentElement = rand.Intn(lastElement + step)
		}
		lastElement = currentElement
		linkedList.PushBack(lastElement)
	}
	return linkedList
}

func createReverseSortedList(listSize int) *list.List {
	linkedList := list.New()
	step := 10
	lastElement := 0
	for i := 0; i < listSize; i++ {
		currentElement := rand.Intn(lastElement + step)
		for currentElement <= lastElement {
			currentElement = rand.Intn(lastElement + step)
		}
		lastElement = currentElement
		linkedList.PushFront(lastElement)
	}
	return linkedList
}

func createPartlySortedList(listSize int, disordersPercent int) *list.List {
	ordersPercent := float32(100 - disordersPercent)                                  // процент упорядоченных чисел
	countOrders := int(float32(listSize) * (ordersPercent / 100))                     // количество элементов которые должны быть упорядочены
	divisionsCount := rand.Intn(countOrders)                                          // количество делений
	maxOrderedElementCount := countOrders/divisionsCount + countOrders%divisionsCount // максимальное количество элементов в одном делении
	defaultOrderedElementCount := countOrders / divisionsCount
	var doubleDimArray []models.DimArray
	var randomIndexes []int

	digitToAdd := 0
	println("YES")

	for i := 0; i < divisionsCount; i++ {
		place := rand.Intn(listSize - maxOrderedElementCount)
		isInside := false
		for j := place; place < place+maxOrderedElementCount; j++ {
			if contains(randomIndexes, j) {
				isInside = true
			}
		}
		if !isInside {
			randomIndexes = append(randomIndexes, place)
			var arrToAdd []int
			for z := 0; z < defaultOrderedElementCount; z++ {
				digitToAdd++
				arrToAdd = append(arrToAdd, digitToAdd)
			}
			doubleDimArray = append(doubleDimArray, models.DimArray{Index: place, Arr: arrToAdd})
		}
	}

	println("YES")
	linkedList := list.New()
	for i := 0; i < listSize; i++ {
		isInside := false
		for j := 0; j < len(doubleDimArray); j++ {
			if i == doubleDimArray[j].Index {
				isInside = true
				for z := 0; z < len(doubleDimArray[j].Arr); z++ {
					linkedList.PushBack(doubleDimArray[j].Arr[z])
					i++
				}
			}
		}
		if !isInside {
			linkedList.PushBack(rand.Intn(listSize-digitToAdd+1) + digitToAdd + 1)
		}

	}
	return linkedList
}
