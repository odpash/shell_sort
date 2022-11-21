package sorts

import (
	"container/list"
	"time"
)

func speedTest(l *list.List, f func(*list.List) *list.List) []int64 {
	var times []int64

	for i := 1; i < 21; i++ { // по 5 %
		testListSize := l.Len() / 100 * i * 5
		testList := list.New()
		count := 0
		for temp := l.Front(); temp != nil; temp = temp.Next() {
			count++
			testList.PushBack(temp.Value)
			if testListSize == count {
				break
			}
		}
		start := time.Now()
		f(testList)
		duration := time.Since(start)
		times = append(times, duration.Nanoseconds())
	}
	return times
}

func RunTests(l *list.List) ([]int64, []int64, []int64) {
	knutResults := speedTest(l, knutSort)
	divByTwoResults := speedTest(l, divByTwoSort)
	virtResults := speedTest(l, virtSort)
	return knutResults, divByTwoResults, virtResults
}
