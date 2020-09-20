package main

import "github.com/vzhang/algorithms/linkedlist"

func main() {
	ll := linkedlist.NewLinkedList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.String()
	ll.Reverse()
	ll.String()

	ll2 := linkedlist.NewLinkedList()
	ll2.Append(4)
	ll2.Append(5)
	ll2.Append(6)

	ll.Merge(ll2)
	ll.String()
	ll.RemoveAtFromLast(3)
	ll.String()
}
