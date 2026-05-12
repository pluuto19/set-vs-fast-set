package setvsbitset_test

import (
	"math/rand/v2"
	s "sets-comparision"
	"testing"
)

func seedStructSet(n int) (*s.SetStruct[int64], []int64) {
	set := s.NewSetStruct[int64](n)
	validInts := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		v := int64(rand.IntN(universe))
		set.Add(v)
		validInts = append(validInts, v)
	}

	return set, validInts
}

func BenchmarkStructSetAdd(b *testing.B) {
	b.StopTimer()
	set := s.NewSetStruct[int64](b.N)
	validInts := make([]int64, 0, b.N)

	for i := 0; i < b.N; i++ {
		validInts = append(validInts, int64(rand.IntN(universe)))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Add(validInts[i])
	}
}

func BenchmarkStructSetDelete(b *testing.B) {
	b.StopTimer()
	set, validInts := seedStructSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Delete(validInts[i])
	}
}

func BenchmarkStructSetContainsBest(b *testing.B) {
	b.StopTimer()
	set, validInts := seedStructSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(validInts[i])
	}
}

func BenchmarkStructSetContainsWorst(b *testing.B) {
	b.StopTimer()
	set := s.NewSetStruct[int64](b.N)

	invalidInts := make([]int64, 0, b.N)
	for i := 0; i < b.N; i++ {
		invalidInts = append(invalidInts, int64(rand.IntN(universe)))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(invalidInts[i])
	}
}

func BenchmarkStructSetContainsAvg(b *testing.B) {
	b.StopTimer()
	set, validInts := seedStructSet(b.N)

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
