package DHash

import (
	"github.com/Amertz08/EECS560-go/Lab03/Bucket"
	"fmt"
)

type DHash struct {
	mod, k, p int
	buckets []Bucket.Bucket
}

func NewDHash(mod int, k int, p int) DHash {
	return DHash{mod: mod, k: k, p: p, buckets: make([]Bucket.Bucket, mod)}
}

func (h *DHash) Insert(value int) {
	if h.Find(value) {
		fmt.Println("Value already exists in table")
		return
	}

	hash := h.Hash(value)
	if hash == -1 {
		fmt.Println("Hash could not be found")
		return
	}

	h.buckets[hash].SetValue(value)
}

func (h *DHash) Remove(value int) int {
	for i := 0; i < h.mod; i++ {
		hash := h.hash(value, i)
		if hash > h.mod {
			hash = hash % h.mod
		}
		if h.buckets[hash].Empty() {
			return hash
		}
	}
	return -1
}

func (h *DHash) hash(value int, i int) int {
	return (value % h.mod + (i * (h.p - (value % h.p)))) % h.mod
}

func (h *DHash) Hash(value int) int {
	for i := 0; i < h.k; i++ {
		hash := h.hash(value, i)
		if hash > h.mod {
			hash = hash % h.mod
		}
		if h.buckets[hash].Empty() {
			return hash
		}
	}
	return -1
}

func (h *DHash) Print() {
	for i, b := range h.buckets {
		if !b.Empty() {
			fmt.Printf("Bucket %d: %d\n", i, b.GetValue())
		}
	}
}

func (h *DHash) Find(value int) bool {
	for i := 0; i < h.k; i++ {
		hash := h.hash(value, i)
		if h.buckets[hash].GetValue() == value {
			return true
		} else if h.buckets[hash].BeenSet() {
			return false
		}
	}
	return false
}

func (h *DHash) LoadFactor() float64 {
	n := 0.0
	for i := 0; i < h.mod; i++ {
		if !h.buckets[i].Empty() {
			n += 1.0
		}
	}
	return n / float64(h.mod)
}
