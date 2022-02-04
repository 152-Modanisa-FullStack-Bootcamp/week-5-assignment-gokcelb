package benchmark

import (
	"bootcamp/assignment"
	"testing"
)

func BenchmarkAlphabetSoupCountingSortLimited(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.AlphabetSoupCountingSortLimited("cbafda")
	}
}

func BenchmarkAlphabetsoupCountingSortLimitless(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.AlphabetSoupCountingSortLimitless("cbafda")
	}
}

func BenchmarkAlphabetSoupBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.AlphabetSoupBubbleSort("cbafda")
	}
}
