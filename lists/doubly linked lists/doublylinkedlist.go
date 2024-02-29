package main

import (
	"fmt"
)

func New[T any](values ...T) *DoublyLinkedList[T] {
	dll := DoublyLinkedList[T]{count: 0}
	for i := 0; i < len(values); i++ {
		fmt.Println("Adding ", values[i])
		dll.Add(values[i])
	}

	return &dll
}

type DLLNode[T any] struct {
	value       T
	ptrPrevious *DLLNode[T]
	ptrNext     *DLLNode[T]
}

type DoublyLinkedList[T any] struct {
	count int
	head  *DLLNode[T]
	tail  *DLLNode[T]
}

func (d *DoublyLinkedList[T]) Size() int {
	return d.count
}

func (d *DoublyLinkedList[T]) IsEmpty() bool {
	return d.Size() == 0
}

func (d *DoublyLinkedList[T]) Add(value T) {
	currentnode := DLLNode[T]{value: value}
	currentnode.ptrNext = nil
	currentnode.ptrPrevious = nil

	if d.Size() > 0 {
		fmt.Println("Size is greater than zero")
		d.head.ptrPrevious = &currentnode
		d.tail = &currentnode
		currentnode.ptrNext = d.Head()
		currentnode.ptrPrevious = d.Tail()
		d.count++
	}

	if d.Size() == 0 {
		d.head = &currentnode
		d.tail = &currentnode
		d.count++
	}
}

func (d *DoublyLinkedList[T]) Head() *DLLNode[T] {
	fmt.Println("Current head value", d.head.value)
	return d.head
}

func (d *DoublyLinkedList[T]) Tail() *DLLNode[T] {
	fmt.Println("Current tail value", d.tail.value)
	return d.tail
}

func (d *DoublyLinkedList[T]) GetAllValues() []T {
	values := make([]T, 0, d.Size())
	current := d.tail
	for current != nil {
		values = append(values, current.value)
		current = current.ptrNext
	}

	return values
}

func main() {
	dll := New(90, 78, 89)
	dll.Add(57)

	fmt.Println("currently", dll.GetAllValues())
	fmt.Println("count", dll.Size())
}
