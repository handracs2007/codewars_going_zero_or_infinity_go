package main

import (
	"fmt"
	"strconv"
	"strings"
)

func process(n int) float64 {
	if n == 0 {
		return float64(0)
	} else {
		return 1 + process(n-1)/float64(n)
	}
}

func Going(n int) float64 {
	result := process(n)
	resultStr := fmt.Sprintf("%.10f", result)

	dotIndex := strings.Index(resultStr, ".")
	maxIndex := dotIndex + 7

	if maxIndex > len(resultStr) {
		maxIndex = len(resultStr)
	}

	scaledStr := resultStr[:maxIndex]
	scaled, _ := strconv.ParseFloat(scaledStr, 10)

	return scaled
}

func main() {
	fmt.Println(Going(8))
}
