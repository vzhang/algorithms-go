package linkedlist

type ElementType interface {}

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

type LinkedList struct {
	Capacity int
	Head *Node
}

func NewLinkedList(capacity int, head *Node) *LinkedList {
	return &LinkedList{
		Capacity: capacity,
		Head: head,
	}
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}

func (ll *LinkedList) Length() int {
	if ll.Head == nil {
		return 0
	}
	length := 1
	curNode := ll.Head
	for {
		if curNode.Next == nil {
			return length
		}
		curNode = curNode.Next
		length = length + 1
	}
}

func (ll *LinkedList) Append(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		return
	}

	curNode := ll.Head
	for {
		if curNode.Next == nil {
			break
		} else {
			curNode = curNode.Next
		}
	}

	curNode.Next = node
}
