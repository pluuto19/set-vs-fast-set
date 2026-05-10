package setvsfastset_test

import (
	"math/rand/v2"
	s "sets-comparision"
	"testing"
)

func seedBoolSet(n int) (*s.SetBool[string], []string) {
	set := s.NewSetBool[string](n)
	validStrings := make([]string, 0, n)
	for i := 0; i < n; i++ {
		str := generateRandomString(strlen, "hit_")
		set.Add(str)
		validStrings = append(validStrings, str)
	}

	return set, validStrings
}

func BenchmarkBoolSetAdd(b *testing.B) {
	b.StopTimer()
	set := s.NewSetBool[string](b.N)
	validStr := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {
		validStr = append(validStr, generateRandomString(strlen, "hit_"))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Add(validStr[i])
	}
}

func BenchmarkBoolSetDelete(b *testing.B) {
	b.StopTimer()
	set, validStr := seedBoolSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Delete(validStr[i])
	}
}

func BenchmarkBoolSetContainsBest(b *testing.B) {
	b.StopTimer()
	set, validStr := seedBoolSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(validStr[i])
	}
}

func BenchmarkBoolSetContainsWorst(b *testing.B) {
	b.StopTimer()
	set, _ := seedBoolSet(b.N)

	invalidStr := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		invalidStr = append(invalidStr, generateRandomString(strlen, "miss_"))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(invalidStr[i])
	}
}

func BenchmarkBoolSetContainsAvg(b *testing.B) {
	b.StopTimer()
	set, validStr := seedBoolSet(b.N)

	invalidStr := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		invalidStr = append(invalidStr, generateRandomString(strlen, "miss_"))
	}

	bools := make([]bool, 0, b.N)
	for i := 0; i < b.N; i++ {
		bools = append(bools, rand.IntN(2) == 0)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if bools[i] {
			set.Contains(validStr[i])
		} else {
			set.Contains(invalidStr[i])
		}
	}
}
