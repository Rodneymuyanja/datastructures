package main

import (
	"errors"
	"fmt"
)

func New[T any](values ...T) *Queue[T] {
	queue := Queue[T]{make([]T, 0, len(values))}
	queue.Push(values...)
	return &queue
}

type Queue[T any] struct {
	array []T
}

func (q *Queue[T]) Size() int {
	return len(q.array)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Push(values ...T) {
	q.array = append(q.array, values...)
}

func (q *Queue[T]) GetAllValues() []T {
	values := make([]T, 0, q.Size())
	values = append(values, q.array...)
	return values
}

func (q *Queue[T]) Pop() (result T, err error) {
	if q.IsEmpty() {
		return result, errors.New("Queue is empty")
	}

	value := q.array[0]
	q.array = q.array[1:]
	return value, nil
}

func (q *Queue[T]) Peek() (result T, err error) {
	if q.IsEmpty() {
		return result, errors.New("Queue is empty")
	}

	value := q.array[q.Size()-1]
	return value, nil
}

func main() {
	q := New(90)

	i := 10
	for i != 0 {
		p, _ := q.Peek()
		q.Push(p - 1)
		i--
	}

	fmt.Println("all ", q.GetAllValues())

	k, _ := q.Pop()

	fmt.Println("popped value ", k)
	fmt.Println("after pop ", q.GetAllValues())

}
