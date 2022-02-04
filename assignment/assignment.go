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
		// append the same letter for as much as its quantity
		if quantity > 0 {
			for k := 0; k < quantity; k++ {
				newWord = append(newWord, byte(i)+smallest)
			}
		}
	}
	return string(newWord)
}

func StringMask(s string, n uint) string {
	return ""
}

func WordSplit(arr [2]string) string {
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
