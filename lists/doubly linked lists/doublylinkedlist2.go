package main

import (
	"fmt"
)

func New[T any](values ...T) *DoublyLinkedList[T] {
	dll := DoublyLinkedList[T]{make([]Node[T], 0, len(values))}
	fmt.Println("length of input", len(values))
	for i := 0; i < len(values); i++ {
		fmt.Println("counting", i)
		dll.Add(values[i])
	}
	return &dll
}

type Node[T any] struct {
	value       T
	ptrPrevious *Node[T]
	ptrNext     *Node[T]
}

type DoublyLinkedList[T any] struct {
	array []Node[T]
}

func (d *DoublyLinkedList[T]) Size() int {
	return len(d.array)
}

func (d *DoublyLinkedList[T]) IsEmpty() bool {
	return d.Size() == 0
}

func (d *DoublyLinkedList[T]) Add(value T) {
	node := Node[T]{value: value}
	node.ptrNext = nil
	node.ptrPrevious = nil

	if d.IsEmpty() {
		fmt.Println("Initial size", d.Size())
		d.array = append(d.array, node)
		fmt.Println("After Initial append", d.Size())
	}

	if d.Size() == 1 {
		currentTailInDll := d.array[0]
		currentTailInDll.ptrPrevious = &node
		node.ptrNext = &currentTailInDll
		d.array = append(d.array, node)
	}

	if d.Size() > 1 {
		currentTailInDll := d.array[d.Size()-1]
		currentTailInDll.ptrPrevious = &node
		node.ptrNext = &currentTailInDll
		d.array = append(d.array, node)
	}
}

func (d *DoublyLinkedList[T]) GetAllValues() []T {
	values := make([]T, 0, d.Size())
	fmt.Println("size of values ", cap(values))
	fmt.Println("size ", d.Size())
	for i := 0; i < d.Size(); i++ {
		fmt.Println("in here", i)
		values = append(values, d.array[i].value)
	}

	return values
}

func (d *DoublyLinkedList[T]) Clear() {
	d.array = nil
}

func main() {
	dll := New(45)
	//dll.Clear()
	// dll.Add(46)
	// dll.Add(89)
	// dll.Add(139)
	// dll.Add(789)

	node := dll.array[len(dll.array)-1]

	fmt.Println("Last node ", node.value)
	// fmt.Println("Node next ", node.ptrNext.value)
	// fmt.Println("Node next ", node.ptrNext.ptrNext.value)

	fmt.Println("currently", dll.GetAllValues())

	// fmt.Println("count", dll.Size())
	// dll.Clear()
	// fmt.Println("count", dll.Size())
}
