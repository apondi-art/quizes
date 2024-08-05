package main

import "fmt"

func main() {
	str := "☕"
	for i := 0 ; i < len(str) ; i++ {
		fmt.Println(string(str[i]))
	}
	for _ , v := range str {
		fmt.Println(string(v))
	}
}
