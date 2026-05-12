package setvsbitset

import "math"

type BitSet struct {
	b []int64
}

func NewBitSet(universe int) *BitSet {
	return &BitSet{
		b: make([]int64, int(math.Ceil(float64(universe)/64))),
	}
}

func (bs *BitSet) Add(x int64) {
	bs.b[x/64] |= 1 << (x % 64)
}

func (bs *BitSet) Clear(x int64) {
	bs.b[x/64] &= ^(1 << (x % 64))
}

func (bs *BitSet) Contains(x int64) bool {
	return bs.b[x/64]>>(x%64)&1 == 1
}
