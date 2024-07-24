package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 { // checks the length of argument being passed
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	j := 0 // keep track of index j where rune of str1 == rune of str2
	for i := 0; i < len(str2); i++ {
		if str2[i] == str1[j] {
			j++
		}
		if j == len(str1) {
			fmt.Println(str1)
			break
		}

	}
}
