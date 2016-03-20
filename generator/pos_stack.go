package generator

type stack struct {
	pos  position
	next *stack
}

func newStack() *stack {
	return (*stack)(nil)
}

func (s *stack) empty() bool {
	return s == nil
}

func (s *stack) push(p position) *stack {
	return &stack{pos: p, next: s}
}

func (s *stack) pop() (*stack, position) {
	return s.next, s.pos
}
