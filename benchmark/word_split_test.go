package benchmark

import (
	"bootcamp/assignment"
	"strings"
	"testing"
)

// Normal cases
var word = "notcat"
var words = "apple,bat,cat,goodbye,hello,yellow,why,someone,noone,everyone,anyone,youredone"
var wordsNThousand = genNormal(1000)
var wordsNTenThousand = genNormal(10000)

func genNormal(n int) string {
	var temp string
	for i := 0; i < n; i++ {
		if i == n-1 {
			temp += words
		} else {
			temp += words + ","
		}
	}
	return temp
}
func BenchmarkWordSplitDict1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitDict([2]string{word, words})
	}
}

func BenchmarkWordSplitDict2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitDict([2]string{word, wordsNThousand})
	}
}

func BenchmarkWordSplitDict3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitDict([2]string{word, wordsNTenThousand})
	}
}

func BenchmarkWordSplitList1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitList([2]string{word, words})
	}
}

func BenchmarkWordSplitList2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitList([2]string{word, wordsNThousand})
	}
}

func BenchmarkWordSplitList3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WordSplitList([2]string{word, wordsNTenThousand})
	}
}

// Isolated cases
var wordsDTen = genD(1)
var wordsDHundred = genD(100)
var wordsDTenThousand = genD(10000)

var wordsLTen = genL(1)
var wordsLHundred = genL(100)
var wordsLTenThousand = genL(10000)

func genD(n int) map[string]int {
	var temp string
	for i := 0; i < n; i++ {
		temp += words
	}
	var wordsD = make(map[string]int)
	for _, word := range strings.Split(words, ",") {
		wordsD[word] += 1
	}
	return wordsD
}

func genL(n int) []string {
	var temp string
	for i := 0; i < n; i++ {
		temp += words
	}
	return strings.Split(temp, ",")
}

func benchmarkD(word string, words map[string]int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WithDict(word, words)
	}
}

func benchmarkL(word string, words []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.WithList(word, words)
	}
}

func BenchmarkWithDict1(b *testing.B) { benchmarkD(word, wordsDTen, b) }

func BenchmarkWithDict2(b *testing.B) { benchmarkD(word, wordsDHundred, b) }

func BenchmarkWithDict3(b *testing.B) { benchmarkD(word, wordsDTenThousand, b) }

func BenchmarkWithList1(b *testing.B) { benchmarkL(word, wordsLTen, b) }

func BenchmarkWithList2(b *testing.B) { benchmarkL(word, wordsLHundred, b) }

func BenchmarkWithList3(b *testing.B) { benchmarkL(word, wordsLTenThousand, b) }
