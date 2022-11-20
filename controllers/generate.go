package controllers

import (
	"container/list"
	"log"
	"math/rand"
)

func createUnsortedList(listSize int) *list.List {
	linkedList := list.New()
	for i := 0; i < listSize; i++ {
		linkedList.PushFront(rand.Intn(listSize))
	}
	return linkedList
}

func createSortedList(listSize int) *list.List {
	linkedList := list.New()
	lastElement := 0
	for i := 0; i < listSize; i++ {
		step := rand.Intn(10) + 1
		linkedList.PushBack(lastElement + step)
		lastElement = lastElement + step
		log.Println(lastElement, " ", i)
	}
	return linkedList
}

func createReverseSortedList(listSize int) *list.List {
	linkedList := list.New()
	lastElement := 0
	for i := 0; i < listSize; i++ {
		step := rand.Intn(10) + 1
		linkedList.PushFront(lastElement + step)
		lastElement = lastElement + step
		log.Println(lastElement, " ", i)
	}
	return linkedList
}

func createPartlySortedList(listSize int, disordersPercent int) *list.List {
	linkedList := list.New()
	lastElement := 0
	counter := 0
	for i := 0; i < listSize; i++ {
		step := rand.Intn(10) + 1
		r := rand.Intn(101)
		var result int
		if r < disordersPercent {
			result = rand.Intn(listSize)
			counter++
		} else {
			result = lastElement + step
			lastElement = result
		}
		linkedList.PushBack(result)
	}
	log.Println(counter)
	return linkedList
}
