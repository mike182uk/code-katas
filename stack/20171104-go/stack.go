package stack

import "errors"

var (
	// ErrUnderflow represents a stack underflow error
	ErrUnderflow = errors.New("Stack empty")
	// ErrOverflow represents a stack overflow error
	ErrOverflow = errors.New("Stack full")
)

// Stack represents a stack
type Stack struct {
	elements []int
	capacity int
}

// NewStack returns a new Stack instance
func NewStack(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
	}
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

// Push adds an element to the stack
func (s *Stack) Push(element int) error {
	if len(s.elements) == s.capacity {
		return ErrOverflow
	}

	s.elements = append(s.elements, element)

	return nil
}

// Pop removes an element from the stack
func (s *Stack) Pop() (int, error) {
	size := len(s.elements)
	lastElemIdx := size - 1

	if size == 0 {
		return -1, ErrUnderflow
	}

	e := s.elements[lastElemIdx]

	s.elements = s.elements[:lastElemIdx]

	return e, nil
}

// Peek returns the element at the top of the stack without removing it from the stack
func (s *Stack) Peek() (int, error) {
	size := len(s.elements)

	if size == 0 {
		return -1, ErrUnderflow
	}

	return s.elements[size-1], nil
}
