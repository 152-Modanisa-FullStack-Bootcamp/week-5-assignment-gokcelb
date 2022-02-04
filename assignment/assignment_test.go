package assignment

import (
	"math"
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

func TestAlphabetSoup(t *testing.T) {
	testCases := []struct {
		desc         string
		givenWord    string
		expectedWord string
	}{
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
	}
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := AlphabetSoup(tC.givenWord)

			assert.Equal(t, tC.expectedWord, result)
		})
	}
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/
	result := StringMask("!mysecret*", 2)

	assert.Equal(t, "!m********", result)
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/
	result := WordSplit([2]string{"hellocat", words})

	assert.Equal(t, "hello,cat", result)
}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
	set := VariadicSet(4, 2, 5, 4, 2, 4)

	assert.Equal(t, []interface{}{4, 2, 5}, set)
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
