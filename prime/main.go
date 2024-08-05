package main

import (
	"fmt"
	"os"
	"strconv"
)

// $ go run . 225225
// 3*3*5*5*7*11*13
// $ go run . 8333325
// 3*3*5*5*7*11*13*37
// $ go run . 9539
// 9539
// $ go run . 804577
// 804577
// $ go run . 42
// 2*3*7
// $ go run . a
// $ go run . 0
// $ go run . 1
// $

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	if n < 4 {
		return
	}
	result := ""
	value := 2
	for n > 1 {
		if n%value == 0 {
			result += strconv.Itoa(value)
			n = n / value

		} else {
			value++
		}
	}
	fmt.Println(result)
}
