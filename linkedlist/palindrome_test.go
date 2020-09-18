package linkedlist

import "testing"

func TestPalindrome(t *testing.T) {
	ll1 := NewLinkedList()
	if ok := ll1.IsPalindrome(); !ok {
		t.Errorf("nil is palindrome, but got false")
	}

	ll2 := NewLinkedList()
	ll2.Append(1)
	if ok := ll2.IsPalindrome(); !ok {
		t.Errorf("one is palindrome, but got false")
	}

	ll3 := NewLinkedList()
	ll3.Append(1)
	ll3.Append(2)
	ll3.Append(3)
	ll3.Append(2)
	ll3.Append(1)
	if ok := ll3.IsPalindrome(); !ok {
		t.Errorf("12321 is palindrome, but got false")
	}

	ll4 := NewLinkedList()
	ll4.Append(1)
	ll4.Append(2)
	ll4.Append(3)
	ll4.Append(3)
	ll4.Append(2)
	ll4.Append(1)
	if ok := ll4.IsPalindrome(); !ok {
		t.Errorf("123321 is palindrome, but got false")
	}

	ll5 := NewLinkedList()
	ll5.Append(1)
	ll5.Append(2)
	ll5.Append(3)
	ll5.Append(1)
	if ok := ll5.IsPalindrome(); ok {
		t.Errorf("1231 is palindrome, but got true")
	}

}
