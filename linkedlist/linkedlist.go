package linkedlist

import (
	"fmt"
	"sync"
)

type ElementType interface{}

type Node struct {
	Data ElementType
	Next *Node
}

func NewNode(value ElementType) *Node {
	return &Node{
		Data: value,
		Next: nil,
	}
}

type SinglyLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func NewLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (ll *SinglyLinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *SinglyLinkedList) Size() int {
	if ll.head == nil {
		return 0
	}
	length := 1
	curNode := ll.head
	for {
		if curNode.Next == nil {
			return length
		}
		curNode = curNode.Next
		length = length + 1
	}
}

func (ll *SinglyLinkedList) Append(value ElementType) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	node := &Node{Data: value}
	if ll.head == nil {
		ll.head = node
		ll.size = 1
		return
	}

	curNode := ll.head
	for {
		if curNode.Next == nil {
			break
		} else {
			curNode = curNode.Next
		}
	}
	ll.size++
	curNode.Next = node
}

func (ll *SinglyLinkedList) InsertAt(i int, value ElementType) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return fmt.Errorf("index out of bounds")
	}
	newNode := &Node{Data: value}
	if i == 0 {
		newNode.Next = ll.head
		ll.head = newNode
		return nil
	}

	curNode := ll.head
	for j := 0; j < i-1; j++ {
		curNode = curNode.Next
	}
	newNode.Next = curNode.Next
	curNode.Next = newNode
	ll.size++
	return nil
}

func (ll *SinglyLinkedList) RemoveAt(i int) (*ElementType, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("index out of bounds")
	}

	curNode := ll.head
	for j := 0; i < i-1; j++ {
		curNode = curNode.Next
	}

	remove := curNode.Next
	curNode.Next = remove.Next
	ll.size--
	return &remove.Data, nil
}

func (ll *SinglyLinkedList) IndexOf(value ElementType) int {
	curNode := ll.head
	i := 0
	for {
		if curNode.Data == value {
			return i
		}

		if curNode.Next == nil {
			return -1
		}

		curNode = curNode.Next
		i++
	}
}

func (ll *SinglyLinkedList) String() {
	curNode := ll.head
	for {
		if curNode == nil {
			break
		}

		fmt.Print(curNode.Data)
		fmt.Print("")
		curNode = curNode.Next
	}
	fmt.Println()
}

func (ll *SinglyLinkedList) Head() *Node {
	return ll.head
}
