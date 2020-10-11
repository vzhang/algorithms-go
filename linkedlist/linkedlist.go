package linkedlist

import (
	"fmt"
	"sync"
)

type ElementType int //interface{}

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
		ll.size++
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
	for j := 0; j < i-1; j++ {
		curNode = curNode.Next
	}

	remove := curNode.Next
	curNode.Next = remove.Next
	ll.size--
	return &remove.Data, nil
}

func (ll *SinglyLinkedList) IndexOf(value ElementType) int {
	if ll.head == nil {
		return -1
	}
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

// 单链表反转
func (ll *SinglyLinkedList) Reverse() {
	if ll.head == nil || ll.head.Next == nil {
		return
	}
	var reversedNode *Node
	curNode := ll.head
	for curNode != nil {
		next := curNode.Next

		var temp = curNode
		temp.Next = reversedNode
		reversedNode = temp

		curNode = next
	}

	ll.head = reversedNode
}

// 链表中的环检测
// 快慢指针，如果相遇了，就说明有环，如果结束了并且没相遇，就证明没有环
func (ll *SinglyLinkedList) LoopDetected() bool {
	if ll.head == nil || ll.head.Next == nil {
		return false
	}

	slowNode := ll.head
	fastNode := ll.head

	for fastNode != nil && fastNode.Next != nil {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next

		if slowNode == fastNode {
			return true
		}
	}

	return false
}

// 两个有序的链表合并
func (ll *SinglyLinkedList) Merge(ll2 *SinglyLinkedList) {
	head := NewNode(-1)
	curNode := head

	curNode1 := ll.head
	curNode2 := ll2.head
	for curNode1 != nil && curNode2 != nil {
		if curNode1.Data < curNode2.Data {
			curNode.Next = curNode1
			curNode1 = curNode1.Next
		} else {
			curNode.Next = curNode2
			curNode2 = curNode2.Next
		}

		curNode = curNode.Next
	}

	if curNode1 != nil {
		curNode.Next = curNode1
	}

	if curNode2 != nil {
		curNode.Next = curNode2
	}

	ll.head = head.Next
}

// 删除链表倒数第n个结点

/*
	解题思路：
	定义两个指针A和B,A和B之间前进的步数相差N
	当B走完的时候，A所在的位置就是结果
*/
func (ll *SinglyLinkedList) RemoveAtFromLast(i int) (*ElementType, error) {
	var slowNode = ll.head
	var fastNode = ll.head

	for j := i; j >= 0; j-- {
		if fastNode == nil {
			return nil, fmt.Errorf("i is bigger than lenght")
		}

		fastNode = fastNode.Next
	}

	for fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}

	ret := slowNode.Next

	slowNode.Next = slowNode.Next.Next

	return &ret.Data, nil
}

// 求链表的中间结点
// 快慢指针，当快指针结束的时候，就是中间结点
func (ll *SinglyLinkedList) Middle() *Node {
	if ll.head == nil {
		return nil
	}

	fastNode := ll.head
	slowNode := ll.head

	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	return slowNode
}
