package stack

import (
	"sync"

	"github.com/vzhang/algorithms/linkedlist"
)

// https://zhuanlan.zhihu.com/p/115282977
type LinkedStack struct {
	root *linkedlist.Node
	size int
	sync.Mutex
}

func (s *LinkedStack) Push(v linkedlist.ElementType) {
	s.Lock()
	defer s.Unlock()

	if s.root == nil {
		s.root = new(linkedlist.Node)
		s.root.Data = v
	} else {
		preNode := s.root
		newNode := new(linkedlist.Node)
		newNode.Data = v
		newNode.Next = preNode

		s.root = newNode
	}

	s.size += 1
}
