package main

import (
	"github.com/vzhang/algorithms/linkedlist"
)

func main() {
	ll := linkedlist.NewLinkedList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(5)
	ll.InsertAt(3, 4)
	ll.RemoveAt(0)
}
