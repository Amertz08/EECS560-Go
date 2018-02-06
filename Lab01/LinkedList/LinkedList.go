package LinkedList

import (
	"fmt"
	"github.com/Amertz08/EECS560-go/Lab01/Node"
)

type LinkedList struct {
	head *Node.Node
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (l *LinkedList) insert(val int) {
	fmt.Printf("Inserting value %v\n", val)
	if l.isEmpty() {
		l.head = Node.NewNode(val)
	} else {
		tmp := l.head
		l.insertHelper(tmp, val)
	}
}

func (l *LinkedList) insertHelper(tmp *Node.Node, val int) {
	if tmp.next == nil {
		tmp.next = Node.NewNode(val)
		return
	}
	l.insertHelper(tmp.next, val)
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) print() {
	fmt.Println("Printing list")
	if l.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	tmp := l.head
	for tmp != nil {
		fmt.Println(tmp.getValue())
		tmp = tmp.next
	}
}

func (l *LinkedList) find(val int) bool {
	tmp := l.head

	for tmp != nil {
		if tmp.value == val {
			return true
		}
		tmp = tmp.next
	}
	return false
}

func (l *LinkedList) eraseHelper(tmp *Node.Node, val int) {
	if tmp.next == nil {
		fmt.Println("Value not in list")
	} else if tmp.next.value == val {
		tmp.next = tmp.next.next
	} else {
		l.eraseHelper(tmp.next, val)
	}
}

func (l *LinkedList) erase(val int) {
	fmt.Printf("Erasing value %v\n", val)
	if l.isEmpty() {
		fmt.Println("Empty list cannot erase value")
		return
	}

	if l.head.value == val {
		l.head = l.head.next
		return
	}

	tmp := l.head
	l.eraseHelper(tmp, val)
}