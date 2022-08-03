package main

import (
	"fmt"
	"container/heap"
)

type MedianOfAges struct {
	maxHeap *MaxHeap // containing first half of numbers
	minHeap *minHeap // containing second half of numbers
}

// initialize the new heaps
func new() *MedianOfAges {
	min := &MinHeap{}
	max := &MaxHeap{}
	heap.Init(min)
	heap.Init(max)
	return &MedianOfAges{minHeap: min, maxHeap: max}
}

func (med *MedianOfAges) FindMedian() float64 {
	// if the max and min size is equal
	// then take largest of small and smallest of large list
	// divide by two
	if med.maxHeap.Len() == med.minHeap.Len() {
		return float64(float64(med.maxHeap.Top() + med.minHeap.Top()) / 2.0)
	}
	// since max list will have the bigger list than the min list
	return float64(med.maxHeap.Top())
}

func (med *MedianOfAges) InsertNum(num int) {
	if med.maxHeap.Empty || med.maxHeap.Top() >= num {
		heap.Push(med.maxHeap, num)
	} else {
		heap.Push(med.minHeap, num)
	}

	// max will have one large list or it will be equal list
	if med.maxHeap.Len() > med.minHeap.len() + 1 {
		heap.Push(med.minHeap, heap.Pop(med.maxHeap).(int))
	} else if med.maxHeap.len() < med.minHeap.Len() {
		heap.Push(med.maxHeap, heap.Pop(med.minHeap).(int))
	}
}

func main() {
	medianOfAges := new()
	medianOfAges.InsertNum(22)
	medianOfAges.InsertNum(35)
	fmt.Printf("Recommendation for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(30)
	fmt.Printf("Recommendation for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(325)
	fmt.Printf("Recommendation for ages under: %f\n", medianOfAges.FindMedian())
}