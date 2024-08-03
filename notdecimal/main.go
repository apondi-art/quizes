package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(CountN("0.1"))
	fmt.Println(CountN("174.2"))
	fmt.Println(CountN("0.1255"))
	fmt.Println(CountN("1.20525856"))
	fmt.Println(CountN("-0.0f00d00"))
	fmt.Println(CountN(""))
	fmt.Println(CountN("-19.525856"))
	fmt.Println(CountN("1952"))
}

func CountN(s string) int {
	var str string

	for i := 1; i < len(s); i++ {
		if s[i] == '.' {
			str += s[i+1:]
		}
	}
	count := len(str)
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.00
	}
	result := math.Pow(10, float64(count))
	number := result * n
	return int(number)
}
