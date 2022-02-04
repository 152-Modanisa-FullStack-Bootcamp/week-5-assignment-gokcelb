package benchmark

import (
	"bootcamp/assignment"
	"testing"
)

func BenchmarkWordSplitDict(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitDict([2]string{"notcat", "apple,bat,cat,goodbye,hello,yellow,why,someone,noone,everyone,anyone,youredone"})
	}
}

func BenchmarkWordSplitList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitList([2]string{"notcat", "apple,bat,cat,goodbye,hello,yellow,why,someone,noone,everyone,anyone,youredone"})
	}
}
