package linkedlist

// 如何判断一个字符串是否是回文字符串的问题
// 解题思路
// 1. 快慢指针同时前进
// 2. 慢指针前进的时候逆序慢指针
// 3. 此时从中间位置开始进行比较
func (ll *SinglyLinkedList) IsPalindrome() bool {
	if ll.head == nil || ll.head.Next == nil {
		return true
	}

	var prevNode *Node
	slowNode := ll.head
	fastNode := ll.head

	for fastNode != nil && fastNode.Next != nil {
		fastNode = fastNode.Next.Next

		next := slowNode.Next
		slowNode.Next = prevNode

		prevNode = slowNode
		slowNode = next
	}

	if fastNode != nil {
		slowNode = slowNode.Next
	}

	for slowNode != nil {
		if slowNode.Data != prevNode.Data {
			return false
		}

		slowNode = slowNode.Next
		prevNode = prevNode.Next
	}

	return true
}
