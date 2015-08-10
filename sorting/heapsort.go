package sorting

import (
	"fmt"
	array "github.com/yikliu/golib/arraylist"
	"github.com/yikliu/golib/heap"
)

func HeapSort(arr *array.List) {
	fmt.Print("before heapsort: ")
	arr.Print()
	heap := heap.New(arr)
	for i := heap.Size(); i > 0; i-- {
		heap.Swap(0, i)
		heap.Heapify(0, i)
	}
	fmt.Print("after heapsort: ")
	arr.Print()
}
