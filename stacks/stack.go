package main

import (
	"errors"
	"fmt"
)

func New[T any](values ...T) *Stack[T] {
	stack := Stack[T]{make([]T, 0, len(values))}
	stack.Push(values...)
	return &stack
}

// This is a Stack definition
// it's constraint is any, ie any type
type Stack[T any] struct {
	array []T
}

func (s *Stack[T]) Pop() (result T, err error) {
	if s.IsEmpty() {
		return result, errors.New("Stack is empty")
	}

	item := s.array[s.Size()-1]
	s.array = s.array[:s.Size()-1]
	return item, nil
}

func (s *Stack[T]) Size() int {
	return len(s.array)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(values ...T) {
	s.array = append(s.array, values...)
}

func (s *Stack[T]) Clear() {
	s.array = nil
}

func (s *Stack[T]) GetAllValues() []T {
	values := make([]T, 0, s.Size())
	values = append(values, s.array...)
	return values
}

func (s *Stack[T]) Peek() (result T, err error) {
	if s.IsEmpty() {
		return result, errors.New("Stack is empty")
	}

	value := s.array[s.Size()-1]
	return value, nil
}

func main() {
	s := New("Hi")
	s.Push("there")

	i := 10
	intStack := New(89)
	for i != 0 {
		k, _ := intStack.Peek()
		intStack.Push(k - 1)
		i--
	}

	intStack.Pop()
	fmt.Println("My string values", s.GetAllValues())

	fmt.Println("My int values", intStack.GetAllValues())

}
