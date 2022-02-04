package benchmark

import (
	"bootcamp/assignment"
	"testing"
)

func BenchmarkCeilNumberStringConversion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.CeilNumberStringConversion(57.35)
	}
}

func BenchmarkCeilNumberNoStringConversion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.CeilNumberNoStringConversion(57.35)
	}
}
