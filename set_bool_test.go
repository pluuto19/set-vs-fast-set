package setvsbitset_test

import (
	"math/rand/v2"
	s "sets-comparison"
	"testing"
)

func seedBoolSet(n int) (*s.SetBool[int64], []int64) {
	set := s.NewSetBool[int64](n)
	validInts := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		v := int64(rand.IntN(universe))
		set.Add(v)
		validInts = append(validInts, v)
	}

	return set, validInts
}

func BenchmarkBoolSetAdd(b *testing.B) {
	b.StopTimer()
	set := s.NewSetBool[int64](b.N)
	validInts := make([]int64, 0, b.N)

	for i := 0; i < b.N; i++ {
		validInts = append(validInts, int64(rand.IntN(universe)))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Add(validInts[i])
	}
}

func BenchmarkBoolSetDelete(b *testing.B) {
	b.StopTimer()
	set, validInts := seedBoolSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Delete(validInts[i])
	}
}

func BenchmarkBoolSetContains(b *testing.B) {
	b.StopTimer()
	set, validInts := seedBoolSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(validInts[i])
	}
}
