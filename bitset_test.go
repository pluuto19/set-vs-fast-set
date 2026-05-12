package setvsbitset_test

import (
	"math/rand/v2"
	s "sets-comparision"
	"testing"
)

const universe = 1 << 24

func seedBitSet(n int) (*s.BitSet, []int64) {
	set := s.NewBitSet(universe)
	validInts := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		v := int64(rand.IntN(universe))
		set.Add(v)
		validInts = append(validInts, v)
	}

	return set, validInts
}

func BenchmarkBitSetAdd(b *testing.B) {
	b.StopTimer()
	set := s.NewBitSet(universe)
	validInts := make([]int64, 0, b.N)

	for i := 0; i < b.N; i++ {
		validInts = append(validInts, int64(rand.IntN(universe)))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Add(validInts[i])
	}
}

func BenchmarkBitSetDelete(b *testing.B) {
	b.StopTimer()
	set, validInts := seedBitSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Clear(validInts[i])
	}
}

func BenchmarkBitSetContainsBest(b *testing.B) {
	b.StopTimer()
	set, validInts := seedBitSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(validInts[i])
	}
}

func BenchmarkBitSetContainsWorst(b *testing.B) {
	b.StopTimer()
	set := s.NewBitSet(universe)

	// invalidInts are values guaranteed not in the set (empty set, all misses)
	invalidInts := make([]int64, 0, b.N)
	for i := 0; i < b.N; i++ {
		invalidInts = append(invalidInts, int64(rand.IntN(universe)))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(invalidInts[i])
	}
}

func BenchmarkBitSetContainsAvg(b *testing.B) {
	b.StopTimer()
	set, validInts := seedBitSet(b.N)

	invalidInts := make([]int64, 0, b.N)
	for i := 0; i < b.N; i++ {
		invalidInts = append(invalidInts, int64(rand.IntN(universe)))
	}

	bools := make([]bool, 0, b.N)
	for i := 0; i < b.N; i++ {
		bools = append(bools, rand.IntN(2) == 0)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if bools[i] {
			set.Contains(validInts[i])
		} else {
			set.Contains(invalidInts[i])
		}
	}
}
