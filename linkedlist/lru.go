package linkedlist

type LRUCache struct {
	*SinglyLinkedList
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		SinglyLinkedList: NewLinkedList(),
		capacity:         capacity,
	}
}

func (lru *LRUCache) Append(value ElementType) {
	// 1 查找元素是否存在
	index := lru.IndexOf(value)
	// 2. 未找到对应的元素
	if index == -1 {
		// 2.1 如果容量满了，删除最后一个元素
		if lru.Size() >= lru.capacity {
			lru.RemoveAt(lru.Size() - 1)
		}
	} else {
		// 3. 找到对应的元素
		// 3.1 删除这个元素
		lru.RemoveAt(index)
	}
	// 4 添加元素到第一个未知
	lru.InsertAt(0, value)
}
