package main

import (
	"fmt"

	"github.com/maxclav/heap"
)

func main() {
	examples()
}

func examples() {
	fmt.Printf("==============[ Min HEAP ]==============\n")
	exampleMinHeap()
	fmt.Printf("==============[ Max HEAP ]==============\n")
	exampleMaxHeap()
	fmt.Println("========================================")
}

func exampleMinHeap() {
	minHeap := heap.NewMin()
	example(minHeap)
}

func exampleMaxHeap() {
	maxHeap := heap.NewMax()
	example(maxHeap)
}

func example(h heap.Heap) {
	values := []int{5, 1, 10, 3, 0, 6}

	fmt.Printf("Inserting values=%v\n", values)
	for _, value := range values {
		h.Insert(value)
		fmt.Printf("\nInserting %v\n", value)
		h.Println()
	}

	fmt.Printf("\nRemoving all values one by one\n")

	for !h.IsEmpty() {
		val := h.Extract()
		fmt.Printf("\nExtracted: %v\n", val)
		h.Println()
	}
}
