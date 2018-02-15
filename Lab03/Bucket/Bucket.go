package Bucket

type Bucket struct {
	value int
}

func NewBucket() Bucket {
	return Bucket{value: -2}
}

func (b *Bucket) GetValue() int {
	return b.value
}

func (b *Bucket) SetValue(value int) {
	b.value = value
}

func (b *Bucket) Remove() {
	b.value = -1
}

func (b *Bucket) Empty() bool {
	return b.value == -1 || b.value == -2
}

func (b *Bucket) BeenSet() bool {
	return b.value != -2
}
