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

func (l *LinkedList) Insert(val int) {
	fmt.Printf("Inserting value %v\n", val)
	if l.isEmpty() {
		l.head = Node.NewNode(val)
	} else {
		tmp := l.head
		l.insertHelper(tmp, val)
	}
}

func (l *LinkedList) insertHelper(tmp *Node.Node, val int) {
	if tmp.Next == nil {
		tmp.Next = Node.NewNode(val)
		return
	}
	l.insertHelper(tmp.Next, val)
}

func (l *LinkedList) InsertFront(val int) {
	newNode := Node.NewNode(val)
	if l.isEmpty() {
		l.head = newNode
	} else {
		tmp := l.head
		l.head = newNode
		l.head.Next = tmp
	}
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) Print() {
	fmt.Println("Printing list")
	if l.isEmpty() {
		fmt.Println("List is empty")
		return
	}
	tmp := l.head
	for tmp != nil {
		fmt.Println(tmp.GetValue())
		tmp = tmp.Next
	}
}

func (l *LinkedList) Find(val int) bool {
	tmp := l.head

	for tmp != nil {
		if tmp.Value == val {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *LinkedList) eraseHelper(tmp *Node.Node, val int) {
	if tmp.Next == nil {
		fmt.Println("Value not in list")
	} else if tmp.Next.Value == val {
		tmp.Next = tmp.Next.Next
	} else {
		l.eraseHelper(tmp.Next, val)
	}
}

func (l *LinkedList) Erase(val int) {
	fmt.Printf("Erasing value %v\n", val)
	if l.isEmpty() {
		fmt.Println("Empty list cannot erase value")
		return
	}

	if l.head.Value == val {
		l.head = l.head.Next
		return
	}

	tmp := l.head
	l.eraseHelper(tmp, val)
}