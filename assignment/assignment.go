package assignment

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y
	overflow := x > math.MaxUint32-y
	return sum, overflow
}

// I wrote two implementations of CeilNumber function to see which was better
// and the benchmark tests showed me that no string conversion was much more performant

func CeilNumberStringConversion(f float64) float64 {
	stringifiedFloat := fmt.Sprintf("%v", f)
	separatedFloat := strings.Split(stringifiedFloat, ".")
	if len(separatedFloat) == 1 {
		return f
	}

	integer, _ := strconv.ParseFloat(separatedFloat[0], 0)
	decimalPoint, _ := strconv.ParseFloat(separatedFloat[1], 0)

	switch {
	case decimalPoint > 75:
		return integer + 1
	case decimalPoint > 50:
		return integer + (75 * math.Pow(10, -2))
	case decimalPoint > 25:
		return integer + (50 * math.Pow(10, -2))
	default:
		return integer + (25 * math.Pow(10, -2))
	}
}

func CeilNumberNoStringConversion(f float64) float64 {
	floored := math.Floor(f)
	ceiled := math.Ceil(f)

	// return the number as it is if decimal point is equal to
	// any threshold (i.e. 0.00, 0.25, 0.50, 0.75)
	if f == floored || f+0.25 == ceiled || f+0.5 == ceiled || f+0.75 == ceiled {
		return f
	}

	if f+0.25 > ceiled {
		return ceiled
	} else if f+0.5 > ceiled {
		return floored + 0.75
	} else if f+0.75 > ceiled {
		return floored + 0.50
	} else {
		return floored + 0.25
	}
}

func AlphabetSoupCountingSortLimited(s string) string {
	// intiialize zero valued fix-length array based on the constraint
	// that parameter s will consist of smallcase latin letters
	var byteList [26]byte
	// assuming smallcase a will be the smallest
	smallest := byte(97)
	// put "a" in the 0th index and form an already sorted array
	// store how many times the letter shows up as value
	for i := 0; i < len(s); i++ {
		byteList[s[i]-smallest] += 1
	}

	var newWord []byte
	for i := 0; i < len(byteList); i++ {
		quantity := int(byteList[i])
		if quantity == 0 {
			continue
		}
		// append the same letter for as much as its quantity
		for k := 0; k < quantity; k++ {
			newWord = append(newWord, byte(i)+smallest)
		}
	}
	return string(newWord)
}

func AlphabetSoupCountingSortLimitless(s string) string {
	// initialize a 128 bit long array to cover the whole ascii table
	var byteList [128]byte
	// fill byteList's indexes
	for i := 0; i < len(s); i++ {
		byteList[s[i]] += 1
	}

	var newWord []byte
	for i := 0; i < len(byteList); i++ {
		quantity := int(byteList[i])
		if quantity == 0 {
			continue
		}
		for k := 0; k < quantity; k++ {
			newWord = append(newWord, byte(i))
		}
	}
	return string(newWord)
}

func AlphabetSoupBubbleSort(s string) string {
	letters := strings.Split(s, "")
	for i := 0; i < len(letters)-1; i++ {
		for j := 0; j < len(letters)-1; j++ {
			if letters[j] > letters[j+1] {
				temp := letters[j]
				letters[j] = letters[j+1]
				letters[j+1] = temp
			}
		}
	}
	return strings.Join(letters, "")
}

func AlphabetSoupMergeSort(s string) string {
	letters := strings.Split(s, "")
	return strings.Join(mergesort(letters), "")
}

func mergesort(letters []string) []string {
	if len(letters) <= 1 {
		return letters
	}
	mid := len(letters) / 2
	left := mergesort(letters[:mid])
	right := mergesort(letters[mid:])
	return merge(left, right)
}

func merge(left, right []string) []string {
	var letters []string
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			letters = append(letters, left[i])
			i++
		} else {
			letters = append(letters, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		letters = append(letters, left[i])
	}
	for ; j < len(right); j++ {
		letters = append(letters, right[j])
	}
	return letters
}

func StringMask(s string, n uint) string {
	var masked string
	if len(s) == 0 {
		return "*"
	}
	if int(n) > len(s) || int(n) == len(s) {
		for i := 0; i < len(s); i++ {
			masked += "*"
		}
		return masked
	}

	toKeep := s[:n]
	toChangeLength := len(s[n:])
	for i := 0; i < toChangeLength; i++ {
		toKeep += "*"
	}
	return toKeep
}

// I wrote two implementations of WordSplit function to see if using a map
// would help me look up words faster; but using a list was significantly faster

// I compared WordSplitDict and WordSplitList with bigger inputs to see if using
// a dictionary would be better for bigger inputs; however, although the performance
// difference between the two decreased, using a list was still better, HERE IS WHY:
// After isolating the look up part, I noticed that the reason behind the poor
// performance of WordSplitDict is the first for loop I was using to create the dictionary.
// Since I am giving bigger inputs for the second element of the string array, the cost of
// looping and allocating map values gets higher and higher, and that's mainly why the WordSliptDict
// function which loops through the same big input twice and allocates dictionary values is less
// performant then the WordSplitList which loops just once

func WordSplitDict(arr [2]string) string {
	possibleMatch := arr[0]
	words := make(map[string]int)
	for _, word := range strings.Split(arr[1], ",") {
		words[word] += 1
	}

	for i := 0; i < len(possibleMatch); i++ {
		firstWord := possibleMatch[:i]
		secondWord := possibleMatch[i:]
		// for every combination of the two words, if any of the two
		// combinations are both in the words, then returns the two words
		_, ok1 := words[firstWord]
		_, ok2 := words[secondWord]
		if ok1 && ok2 {
			return fmt.Sprintf("%s,%s", firstWord, secondWord)
		}
	}
	return "not possible"
}

func WordSplitList(arr [2]string) string {
	possibleMatch := arr[0]
	words := strings.Split(arr[1], ",")

	for i := 0; i < len(possibleMatch); i++ {
		firstWord := possibleMatch[:i]
		secondWord := possibleMatch[i:]
		if scontains(firstWord, words) && scontains(secondWord, words) {
			return fmt.Sprintf("%s,%s", firstWord, secondWord)
		}
	}
	return "not possible"
}

// I isolated the WordSplitDict and WordSplitList funcitons' look ups
// As a result, while WhitList function was faster for very small inputs,
// WithDict funciton was much faster overall, and especially as the
// input got bigger.

func WithDict(possibleMatch string, words map[string]int) string {
	for i := 0; i < len(possibleMatch); i++ {
		firstWord := possibleMatch[:i]
		secondWord := possibleMatch[i:]
		_, ok1 := words[firstWord]
		_, ok2 := words[secondWord]
		if ok1 && ok2 {
			return fmt.Sprintf("%s,%s", firstWord, secondWord)
		}
	}
	return "not possible"
}

func WithList(possibleMatch string, words []string) string {
	for i := 0; i < len(possibleMatch); i++ {
		firstWord := possibleMatch[:i]
		secondWord := possibleMatch[i:]
		if scontains(firstWord, words) && scontains(secondWord, words) {
			return fmt.Sprintf("%s,%s", firstWord, secondWord)
		}
	}
	return "not possible"
}

// helper function to find out if string is in string list
func scontains(s string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if s == arr[i] {
			return true
		}
	}
	return false
}

func VariadicSet(args ...interface{}) []interface{} {
	// form a list to store unique values
	var set []interface{}
	for _, arg := range args {
		// if value is already in set, ignore and continue looping
		if icontains(arg, set) {
			continue
		}
		set = append(set, arg)
	}
	return set
}

// helper function to find out if interface is in interface list
func icontains(arg interface{}, arr []interface{}) bool {
	for i := 0; i < len(arr); i++ {
		if arg == arr[i] {
			return true
		}
	}
	return false
}
