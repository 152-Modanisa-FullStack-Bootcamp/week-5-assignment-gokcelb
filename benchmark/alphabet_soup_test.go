package benchmark

import (
	"bootcamp/assignment"
	"testing"
)

var testStringFive = genTestString(5)
var testStringHundred = genTestString(100)

func genTestString(n int) (testString string) {
	random := "cbadjfklasjdfkvmewjztytjsaagjeteoqrewpffaskdvmpjuyxcs"
	for i := 0; i < n; i++ {
		testString += random
	}
	return
}

func benchmarkCountingSortLimited(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.AlphabetSoupCountingSortLimited(s)
	}
}

func BenchmarkAlphabetSoupCountingSortLimited1(b *testing.B) {
	benchmarkCountingSortLimited(testStringFive, b)
}

func BenchmarkAlphabetSoupCountingSortLimited2(b *testing.B) {
	benchmarkCountingSortLimited(testStringHundred, b)
}

func benchmarkCountingSortLimitless(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.AlphabetSoupCountingSortLimitless(s)
	}
}

func BenchmarkAlphabetsoupCountingSortLimitless1(b *testing.B) {
	benchmarkCountingSortLimitless(testStringFive, b)
}

func BenchmarkAlphabetsoupCountingSortLimitless2(b *testing.B) {
	benchmarkCountingSortLimitless(testStringHundred, b)
}

func benchmarkBubbleSort(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.AlphabetSoupBubbleSort(s)
	}
}

func BenchmarkAlphabetSoupBubbleSort1(b *testing.B) { benchmarkBubbleSort(testStringFive, b) }

func BenchmarkAlphabetSoupBubbleSort2(b *testing.B) { benchmarkBubbleSort(testStringHundred, b) }

func benchmarkMergeSort(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		assignment.AlphabetSoupMergeSort(s)
	}
}
func BenchmarkAlphabetSoupMergeSort1(b *testing.B) { benchmarkMergeSort(testStringFive, b) }

func BenchmarkAlphabetSoupMergeSort2(b *testing.B) { benchmarkMergeSort(testStringHundred, b) }
