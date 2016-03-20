package generator

import "testing"

func TestStackCreation(t *testing.T) {
	s := newStack()
	if s != nil {
		t.Fatal("stack should be nil, but wasn't")
	}
	if !s.empty() {
		t.Fatal("stack should report empty, but didn't")
	}
}

func TestPushingItem(t *testing.T) {
	s := newStack()
	s = s.push(position{1, 2})
	if s == nil {
		t.Fatal("stack should not be nil after push, but was")
	}
	if !(s.pos.row == 1 && s.pos.col == 2) {
		t.Fatal("correct value was not present on stack")
	}
	if s.next != nil {
		t.Fatal("next value on stack should have been nil, but wasn't")
	}
}

func TestPoppingItems(t *testing.T) {
	s := newStack()
	s = s.push(position{1, 2})
	var pos position
	s, pos = s.pop()
	if s != nil {
		t.Fatal("stack is not nil, it should be")
	}
	if !(pos.row == 1 && pos.col == 2) {
		t.Fatal("incorrect position returned")
	}
}

func TestPushingTwice(t *testing.T) {
	s := newStack()
	s = s.push(position{1, 2}).push(position{3, 4})
	var pos position
	s, pos = s.pop()
	if s == nil {
		t.Fatal("stack is nil, but should not be")
	}
	if !(pos.row == 3 && pos.col == 4) {
		t.Fatal("incorrect position returned first")
	}
	s, pos = s.pop()
	if s != nil {
		t.Fatal("stack is nil, but should not be")
	}
	if !(pos.row == 1 && pos.col == 2) {
		t.Fatal("incorrect position returned second")
	}
}
