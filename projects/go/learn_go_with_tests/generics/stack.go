// Package generics implements a simple stack data structure using Go's generics.
package generics

// Stack is a simple stack data structure that can hold values of any type.
type Stack[T any] struct {
	values []T
}

// NewStack creates and returns a new Stack instance for the specified type.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a new value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// IsEmpty checks if the stack is empty and returns true if it is, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Pop removes and returns the value at the top of the stack.
// If the stack is empty, it returns the zero value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
