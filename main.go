package main

import (
	"fmt"

	"github.com/vzhang/algorithms/sort"
)

func main() {
	//ll := linkedlist.NewLinkedList()
	//ll.Append(1)
	//ll.Append(2)
	//ll.Append(3)
	//ll.String()
	//ll.Reverse()
	//ll.String()
	//
	//ll2 := linkedlist.NewLinkedList()
	//ll2.Append(4)
	//ll2.Append(5)
	//ll2.Append(6)
	//
	//ll.Merge(ll2)
	//ll.String()
	//ll.RemoveAtFromLast(3)
	//ll.String()

	data := []int{5, 8, 5, 2, 9, 1, 1, 12, 3, 4, 5}
	sort.SelectionSort(data)
	fmt.Println(data)
}
