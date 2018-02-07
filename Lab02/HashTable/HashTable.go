package HashTable

import (
	"github.com/Amertz08/EECS560-go/Lab01/LinkedList"
	"fmt"
)

type HashTable struct {
	list []LinkedList.LinkedList
	mod int
}

func NewHashTable(mod int) HashTable {
	table := HashTable{mod:mod}

	for i := 0; i < mod; i++ {
		table.list = append(table.list, LinkedList.NewLinkedList())
	}
	return table
}

func (t *HashTable) Print() {
	for i := 0; i < t.mod; i++ {
		fmt.Printf("%d: ", i)
		t.list[i].Print()
	}
}

func (t *HashTable) hash(val int) int {
	return val % t.mod
}

func (t *HashTable) Insert(val int) {
	if t.Find(val) {
		fmt.Println("Value already exists")
		return
	}
	t.list[t.hash(val)].InsertFront(val)
}

func (t *HashTable) Delete(val int) {
	t.list[t.hash(val)].Erase(val)
}

func (t *HashTable) Find(val int) bool {
	return t.list[t.hash(val)].Find(val)
}