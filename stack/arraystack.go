package stack

import "sync"

type ArrayStack struct {
	items []interface{}
	size  int
	sync.Mutex
}

func NewStack(capacity int) *ArrayStack {
	return &ArrayStack{
		items: make([]interface{}, 0, 32),
	}
}

func (s *ArrayStack) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()

	s.items = append(s.items, v)
	s.size += 1
}

func (s *ArrayStack) Pop() (v interface{}) {
	s.Lock()
	defer s.Unlock()

	if s.size == 0 {
		panic("empty")
	}

	v = s.items[s.size-1]
	s.items = s.items[:s.size-1]

	// 移动压缩空间
	newItems := make([]interface{}, s.size-1, s.size-1)
	for i := 0; i < s.size-1; i++ {
		newItems[i] = s.items[i]
	}
	s.items = newItems
	s.size -= 1
	return
}

func (s *ArrayStack) Peek() (v interface{}) {
	if s.size == 0 {
		panic("empty")
	}
	v = s.items[:s.size-1]
	return
}

func (s *ArrayStack) Size() int {
	return s.size
}

func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}
