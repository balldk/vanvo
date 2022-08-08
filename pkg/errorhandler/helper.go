package errorhandler

import (
	"fmt"
	"strings"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findNumDigit(n int) int {
	digit := 0

	for n != 0 {
		n /= 10
		digit += 1
	}

	return digit
}

func padSpaceNum(num, maxDigit int) string {
	digit := findNumDigit(num)
	s := fmt.Sprint(num)
	s += strings.Repeat(" ", maxDigit-digit)

	return s
}
