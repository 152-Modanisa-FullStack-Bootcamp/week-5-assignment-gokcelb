package assignment

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	testCases := []struct {
		desc             string
		givenNum1        uint32
		givenNum2        uint32
		expectedSum      uint32
		expectedOverflow bool
	}{
		{
			desc:             "given maxuint32 and 1, expect 0 and overflow true",
			givenNum1:        math.MaxUint32,
			givenNum2:        1,
			expectedSum:      0,
			expectedOverflow: true,
		},
		{
			desc:             "given 1 and 1, expect 2 and overflow false",
			givenNum1:        1,
			givenNum2:        1,
			expectedSum:      2,
			expectedOverflow: false,
		},
		{
			desc:             "given 42 and 2701, expect 2743 and overflow false",
			givenNum1:        42,
			givenNum2:        2701,
			expectedSum:      2743,
			expectedOverflow: false,
		},
		{
			desc:             "given 42 and maxuint32, expect 42 and overflow true",
			givenNum1:        42,
			givenNum2:        math.MaxUint32,
			expectedSum:      41,
			expectedOverflow: true,
		},
		{
			desc:             "given 4294967290 and 5, expect 4294967295 and overflow false",
			givenNum1:        4294967290,
			givenNum2:        5,
			expectedSum:      4294967295,
			expectedOverflow: false,
		},
		{
			desc:             "given 4294967290 and 6, expect 0 and overflow true",
			givenNum1:        4294967290,
			givenNum2:        6,
			expectedSum:      0,
			expectedOverflow: true,
		},
		{
			desc:             "given 4294967290 and 10, expect 4 and overflow true",
			givenNum1:        4294967290,
			givenNum2:        10,
			expectedSum:      4,
			expectedOverflow: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			sum, overflow := AddUint32(tC.givenNum1, tC.givenNum2)

			assert.Equal(t, tC.expectedSum, sum)
			assert.Equal(t, tC.expectedOverflow, overflow)
		})
	}
}

func TestCeilNumberStringConversion(t *testing.T) {
	testCases := formCeilNumTestCases()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result1 := CeilNumberStringConversion(tC.givenNum)

			assert.Equal(t, tC.expectedNum, result1)
		})
	}
}

func TestCeilNumberNoStringConversion(t *testing.T) {
	testCases := formCeilNumTestCases()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CeilNumberNoStringConversion(tC.givenNum)

			assert.Equal(t, tC.expectedNum, result)
		})
	}
}

func TestAlphabetSoupCountingSortLimited(t *testing.T) {
	testCases := formAlphabetSoupTestCases()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := AlphabetSoupCountingSortLimited(tC.givenWord)

			assert.Equal(t, tC.expectedWord, result)
		})
	}
}

func TestAlphabetSoupCoutingSortLimitless(t *testing.T) {
	testCases := formAlphabetSoupTestCases()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := AlphabetSoupCountingSortLimitless(tC.givenWord)

			assert.Equal(t, tC.expectedWord, result)
		})
	}
}

func TestAlphabetSoupBubbleSort(t *testing.T) {
	testCases := formAlphabetSoupTestCases()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := AlphabetSoupBubbleSort(tC.givenWord)

			assert.Equal(t, tC.expectedWord, result)
		})
	}
}

func TestStringMask(t *testing.T) {
	testCases := []struct {
		desc              string
		givenString       string
		givenReferenceNum uint
		expectedString    string
	}{
		{
			desc:              "given !mysecret* and 2, expect !m********",
			givenString:       "!mysecret*",
			givenReferenceNum: 2,
			expectedString:    "!m********",
		},
		{
			desc:              "given empty string and a positive integer, expect *",
			givenString:       "",
			givenReferenceNum: uint(rand.Int()),
			expectedString:    "*",
		},
		{
			desc:              "given string and 0, expect ******",
			givenString:       "string",
			givenReferenceNum: 0,
			expectedString:    "******",
		},
		{
			desc:              "given string and 3, expect str***",
			givenString:       "string",
			givenReferenceNum: 3,
			expectedString:    "str***",
		},
		{
			desc:              "given string and 5, expect strin*",
			givenString:       "string",
			givenReferenceNum: 5,
			expectedString:    "strin*",
		},
		{
			desc:              "given string and 6, expect ******",
			givenString:       "string",
			givenReferenceNum: 6,
			expectedString:    "******",
		},
		{
			desc:              "given string and 7, expect ******",
			givenString:       "string",
			givenReferenceNum: 7,
			expectedString:    "******",
		},
		{
			desc:              "given s*r*n* and 3, expect s*r***",
			givenString:       "s*r*n*",
			givenReferenceNum: 3,
			expectedString:    "s*r***",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := StringMask(tC.givenString, tC.givenReferenceNum)

			assert.Equal(t, tC.expectedString, result)
		})
	}
}

func TestWordSplitDict(t *testing.T) {
	testCases := formWordSplitTestCases()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := WordSplitDict([2]string{tC.givenWord, tC.givenWordList})

			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestWordSplitList(t *testing.T) {
	testCases := formWordSplitTestCases()
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := WordSplitList([2]string{tC.givenWord, tC.givenWordList})

			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestVariadicSet(t *testing.T) {
	testCases := []struct {
		desc     string
		args     []interface{}
		expected []interface{}
	}{
		{
			desc:     "given 4, 2, 5, 4, 2 and 4 as parameters, expect []interface{4,2,5} []interface{bootcamp,rocks!,really}",
			args:     []interface{}{4, 2, 5, 4, 2},
			expected: []interface{}{"bootcamp", "rocks!", "really"},
		},
		{
			desc:     "given bootcamp,rocks!,really,rocks! as parameters, expect",
			args:     []interface{}{"bootcamp", "rocks!", "really", "rocks!"},
			expected: []interface{}{"bootcamp", "rocks!", "really"},
		},
		{
			desc:     "given 1,uint32(1),first,2,uint32(2),second,1,uint32(2),first as parameters, expect []interface{1,uint32(1),first,2,uint32(2),second",
			args:     []interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"},
			expected: []interface{}{1, uint32(1), "first", 2, uint32(2), "second"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			set := VariadicSet(4, 2, 5, 4, 2, 4)

			assert.Equal(t, []interface{}{4, 2, 5}, set)
		})
	}
}

type CeilNum struct {
	desc        string
	givenNum    float64
	expectedNum float64
}

func formCeilNumTestCases() []CeilNum {
	return []CeilNum{
		{
			desc:        "given 42.42, expect 42.50",
			givenNum:    42.42,
			expectedNum: 42.50,
		},
		{
			desc:        "given 42, expect 42",
			givenNum:    42,
			expectedNum: 42,
		},
		{
			desc:        "given 42.01, expect 42.25",
			givenNum:    42.01,
			expectedNum: 42.25,
		},
		{
			desc:        "given 42.24, expect 42.25",
			givenNum:    42.24,
			expectedNum: 42.25,
		},
		{
			desc:        "given 42.25, expect 42.25",
			givenNum:    42.25,
			expectedNum: 42.25,
		},
		{
			desc:        "given 42.26, expect 42.50",
			givenNum:    42.26,
			expectedNum: 42.50,
		},
		{
			desc:        "given 42.55, expect 42.75",
			givenNum:    42.55,
			expectedNum: 42.75,
		},
		{
			desc:        "given 42.75, expect 42.75",
			givenNum:    42.75,
			expectedNum: 42.75,
		},
		{
			desc:        "given 42.76, expect 43",
			givenNum:    42.76,
			expectedNum: 43,
		},
		{
			desc:        "given 42.99, expect 43",
			givenNum:    42.99,
			expectedNum: 43,
		},
		{
			desc:        "given 43.13, expect 43.25",
			givenNum:    43.13,
			expectedNum: 43.25,
		},
	}
}

type AlphabetSoup struct {
	desc         string
	givenWord    string
	expectedWord string
}

func formAlphabetSoupTestCases() []AlphabetSoup {
	return []AlphabetSoup{
		{
			desc:         "given hello, expect ehllo",
			givenWord:    "hello",
			expectedWord: "ehllo",
		},
		{
			desc:         "given empty string, expect empty string",
			givenWord:    "",
			expectedWord: "",
		},
		{
			desc:         "given h, expect h",
			givenWord:    "h",
			expectedWord: "h",
		},
		{
			desc:         "given ab, expect ab",
			givenWord:    "ab",
			expectedWord: "ab",
		},
		{
			desc:         "given ba, expect ab",
			givenWord:    "ba",
			expectedWord: "ab",
		},
		{
			desc:         "given bac, expect abc",
			givenWord:    "bac",
			expectedWord: "abc",
		},
		{
			desc:         "given cba, expect abc",
			givenWord:    "cba",
			expectedWord: "abc",
		},
		{
			desc:         "given cbafda, expect aabcdf",
			givenWord:    "cbafda",
			expectedWord: "aabcdf",
		},
	}
}

type WordSplit struct {
	desc           string
	givenWord      string
	givenWordList  string
	expectedResult string
}

func formWordSplitTestCases() []WordSplit {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	return []WordSplit{
		{
			desc:           "given hellocat and words, expect hello,cat",
			givenWord:      "hellocat",
			givenWordList:  words,
			expectedResult: "hello,cat",
		},
		{
			desc:           "given catbat and words, expect cat,bat",
			givenWord:      "catbat",
			givenWordList:  words,
			expectedResult: "cat,bat",
		},
		{
			desc:           "given yellowapple and words, expect yellow,words",
			givenWord:      "yellowapple",
			givenWordList:  words,
			expectedResult: "yellow,apple",
		},
		{
			desc:           "given empty string and words, expect not possible",
			givenWord:      "",
			givenWordList:  words,
			expectedResult: "not possible",
		},
		{
			desc:           "given notcat and words, expect not possible",
			givenWord:      "notcat",
			givenWordList:  words,
			expectedResult: "not possible",
		},
		{
			desc:           "given bootcamprocks! and words, expect not possible",
			givenWord:      "bootcamprocks!",
			givenWordList:  words,
			expectedResult: "not possible",
		},
	}
}
