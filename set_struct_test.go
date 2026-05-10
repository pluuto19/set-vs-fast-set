package setvsfastset_test

import (
	"math/rand/v2"
	s "sets-comparision"
	"strings"
	"testing"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const strlen = 16

func generateRandomString(n int, prefix string) string {
	var sb = strings.Builder{}

	sb.WriteString(prefix)

	for j := 0; j < n; j++ {
		sb.WriteByte(charset[rand.IntN(len(charset))])
	}

	return sb.String()
}

func seedStructSet(n int) (*s.SetStruct[string], []string) {
	set := s.NewSetStruct[string](n)
	validStrings := make([]string, 0, n)
	for i := 0; i < n; i++ {
		str := generateRandomString(strlen, "hit_")
		set.Add(str)
		validStrings = append(validStrings, str)
	}

	return set, validStrings
}

func BenchmarkStructSetAdd(b *testing.B) {
	b.StopTimer()
	set := s.NewSetStruct[string](b.N)
	validStr := make([]string, 0, b.N)

	for i := 0; i < b.N; i++ {
		validStr = append(validStr, generateRandomString(strlen, "hit_"))
	}
	
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Add(validStr[i])
	}
}

func BenchmarkStructSetDelete(b *testing.B) {
	b.StopTimer()
	set, validStr := seedStructSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Delete(validStr[i])
	}
}

func BenchmarkStructSetContainsBest(b *testing.B) {
	b.StopTimer()
	set, validStr := seedStructSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(validStr[i])
	}
}

func BenchmarkStructSetContainsWorst(b *testing.B) {
	b.StopTimer()
	set, _ := seedStructSet(b.N)

	invalidStr := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		invalidStr = append(invalidStr, generateRandomString(strlen, "miss_"))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.Contains(invalidStr[i])
	}
}

func BenchmarkStructSetContainsAvg(b *testing.B) {
	b.StopTimer()
	set, validStr := seedStructSet(b.N)

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
