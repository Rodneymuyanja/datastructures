package main

import "fmt"

type Node[T any] struct {
	value       T
	ptrPrevious *Node[T]
	ptrNext     *Node[T]
}

type DoublyLinkedList[T any] struct {
	count int
	head  *Node[T]
	tail  *Node[T]
}

func New[T any](values ...T) *DoublyLinkedList[T] {
	dll := DoublyLinkedList[T]{count: 0}
	for i := 0; i < len(values); i++ {
		dll.Add(values[i])
	}

	return &dll
}

func (d *DoublyLinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}

	if d.IsEmpty() {
		//head can't have a next
		d.head = newNode
		d.head.ptrPrevious = newNode
		d.head.ptrNext = nil

		//tail can't have a previous
		d.tail = newNode
		d.tail.ptrNext = newNode
		d.tail.ptrPrevious = nil
	}

	if !d.IsEmpty() {
		//get the current tail
		tailNode := d.tail
		//set it's previous pointer to the new node
		d.tail.ptrPrevious = newNode
		//set the new node's next pointer to current tail node
		newNode.ptrNext = tailNode
		newNode.ptrPrevious = nil

		d.tail = newNode
		//increase the count in our collection

	}

	d.count++
}

func (d *DoublyLinkedList[T]) Tail() *Node[T] {
	return d.tail
}

func (d *DoublyLinkedList[T]) IsEmpty() bool {
	return d.count == 0
}

func (d *DoublyLinkedList[T]) Size() int {
	return d.count
}

func (d *DoublyLinkedList[T]) GetAllValues() []T {
	values := make([]T, 0, d.count)
	fmt.Println("dll count", d.count)
	currentHead := d.head

	for currentHead != nil {
		values = append(values, currentHead.value)
		currentHead = currentHead.ptrPrevious
	}

	return values
}

func (d *DoublyLinkedList[T]) Clear() {
	currentTail := d.tail

	for currentTail != nil {
		temp := currentTail
		currentTail = currentTail.ptrNext
		temp.ptrNext = nil
		temp.ptrPrevious = nil
		temp = nil
	}

	currentTail = nil
	d.head = nil
	d.count = 0
}

func (d *DoublyLinkedList[T]) GetNodeAt(index int) *Node[T] {

	if index >= d.Size() {
		panic("index out of bounds")
	}

	currenthead := d.head
	for i := 1; i <= index; i++ {
		currenthead = currenthead.ptrPrevious
	}

	return currenthead
}

func (d *DoublyLinkedList[T]) InsertAt(value T, index int) {
	if index < 0 || index > d.Size() {
		panic("index out of bounds")
	}

	if index == 0 {
		d.replaceAtZero(value)
		return
	}

	newNode := &Node[T]{value: value}

	//get the node being replaced
	nodeBeingReplaced := d.GetNodeAt(index)

	//set the new node's previous pointer to the node
	//being replaced
	newNode.ptrPrevious = nodeBeingReplaced

	//sabotage the nodeBeingReplaced.ptrNext and give it
	//to the new node
	newNode.ptrNext = nodeBeingReplaced.ptrNext

	//get the next node's previous pointer and
	//assign it to the new node
	newNode.ptrNext.ptrPrevious = newNode

	//inform the nodeBeingReplaced that it has a new next pointer
	nodeBeingReplaced.ptrNext = newNode

	d.count++
}

func (d *DoublyLinkedList[T]) replaceAtZero(value T) {
	newNode := Node[T]{value: value}

	newNode.ptrNext = nil

	d.head.ptrNext = &newNode
	newNode.ptrPrevious = d.head

	d.head = &newNode
	d.count++
}

func main() {
	dll := New(90, 100, 200, 300, 400)
	dll.Add(78)
	dll.Add(9000)

	fmt.Println("currently ", dll.GetAllValues())

	// dll.InsertAt(560, 2)
	dll.InsertAt(36645, 0)

	dll.Clear()
	fmt.Println("currently ", dll.GetAllValues())

	dll.Add(670)
	fmt.Println("currently ", dll.GetAllValues())
}
