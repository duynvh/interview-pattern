package main

type Pair []int

type Bucket struct {
	b []Pair
}

func BucketInit() *Bucket {
	bucket := &Bucket{
		b: make([]Pair, 0),
	}

	return bucket
}

func (b *Bucket) Get(hashKey int) int {
	for _, pair := range b.b {
		if pair[0] == hashKey {
			return pair[1]
		}
	}

	return -1
}

func (b *Bucket) Update(hashKey int, value int) {
	for i, pair := range b.b {
		if pair[0] == hashKey {
			b.b[i][1] = value
		}
	}

	b.b = append(b.b, Pair{hashKey, value})
}

func (b *Bucket) Remove(hashKey int) {
	for i, pair := range b.b {
		if pair[0] == hashKey {
			b.b = append(b.b[:i], b.b[i+1:]...)
			return
		}
	}
}
