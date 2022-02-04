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

func AlphabetSoup(s string) string {
	// intiialize zero valued fix-length array based on the constraint
	// that parameter s will consist of smallcase latin letters
	var byteList [26]byte
	// assuming smallcase a will be the smallest
	smallest := byte(97)
	// put "a" in the 0th index to form an already sorted array
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
		// for every combination of the two words, if any of the two
		// combinations are both in the words, then returns the two words
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
		} else {
			set = append(set, arg)
		}
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
