package main

import (
	"fmt"
	"os"
)

func splitter(input string) []string {
	var res []string
	var temp string

	for _, r := range input {
		if r == ' ' && temp != "" {
			res = append(res, temp)
			temp = ""
		}
		if r != ' ' {
			temp += string(r)
		}
	}

	if temp != "" {
		res = append(res, temp)
	}
	return res
}

func CleanStr(input string) string {
	str := splitter(input)
	res := ""

	for i, word := range str {
		if i != len(str)-1 {
			res = word + " " + res
			continue
		}
		res = word + res
	}

	return res
}

func main() {
	if len(os.Args) == 2 {
		res := CleanStr(os.Args[1])
		fmt.Println(res)
		return
	}
	fmt.Println()
}
