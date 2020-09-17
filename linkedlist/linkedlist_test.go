package linkedlist

import (
	"fmt"
	"testing"
)

var ll SinglyLinkedList

func TestAppend(t *testing.T) {
	if !ll.IsEmpty() {
		t.Errorf("list should be empty")
	}

	ll.Append("first")
	if ll.IsEmpty() {
		t.Errorf("list should not be empty")
	}

	if size := ll.size; size != 1 {
		t.Errorf("wrong count, expected 1 and got %v", size)
	}

	ll.Append("second")
	ll.Append("third")
	if size := ll.size; size != 3 {
		t.Errorf("wrong count, expected 3 and got %v", size)
	}
}

func TestRemoveAt(t *testing.T) {
	_, err := ll.RemoveAt(1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := ll.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %v", size)
	}
}

func TestInsert(t *testing.T) {
	err := ll.InsertAt(1, "second2")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := ll.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	err = ll.InsertAt(0, "zero")
	if err != nil {
		t.Errorf("unexpected errpr %s", err)
	}
}

func TestIndexOf(t *testing.T) {
	if i := ll.IndexOf("zero"); i != 0 {
		t.Errorf("expected position 0 but got %d", i)
	}
	if i := ll.IndexOf("first"); i != 1 {
		t.Errorf("expected position 1 but got %d", i)
	}
	if i := ll.IndexOf("second2"); i != 2 {
		t.Errorf("expected position 2 but got %d", i)
	}
	if i := ll.IndexOf("third"); i != 3 {
		t.Errorf("expected position 2 but got %d", i)
	}
}

func TestHead(t *testing.T) {
	ll.Append("zero")
	h := ll.Head()
	if "zero" != fmt.Sprint(h.Data) {
		t.Errorf("Expected `zero` but got %s", fmt.Sprint(h.Data))
	}
}
