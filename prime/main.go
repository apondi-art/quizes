package main

import (
	"fmt"
	"os"
	"strconv"
)

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
			if result == "" {
				result += strconv.Itoa(value)
			} else {
				result += "*" + strconv.Itoa(value)
			}

			n = n / value

		} else {
			value++
		}
	}
	fmt.Println(result)
}
