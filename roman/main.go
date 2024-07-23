package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	var calculate string
	var sign string
	type Data struct {
		Numericals  int
		Calculation string
		Signs       string
	}
	Calculation := []Data{
		{
			Numericals:  1000,
			Calculation: "M",
			Signs:       "M",
		},
		{
			Numericals:  900,
			Calculation: "(M-C)",
			Signs:       "CM",
		},

		{
			Numericals:  500,
			Calculation: "D",
			Signs:       "D",
		},
		{
			Numericals:  400,
			Calculation: "(D-C)",
			Signs:       "CD",
		},
		{
			Numericals:  100,
			Calculation: "C",
			Signs:       "C",
		},
		{
			Numericals:  90,
			Calculation: "(C-X)",
			Signs:       "XC",
		},
		{
			Numericals:  50,
			Calculation: "L",
			Signs:       "L",
		},
		{
			Numericals:  40,
			Calculation: "(L-X)",
			Signs:       "XL",
		},
		{
			Numericals:  10,
			Calculation: "X",
			Signs:       "X",
		},
		{
			Numericals:  9,
			Calculation: "(X-I)",
			Signs:       "IX",
		},
		{
			Numericals:  5,
			Calculation: "V",
			Signs:       "V",
		},
		{
			Numericals:  4,
			Calculation: "(V-I)",
			Signs:       "IV",
		},
		{
			Numericals:  1,
			Calculation: "I",
			Signs:       "I",
		},
	}
	for _, value := range Calculation {
		for n >= value.Numericals {
			if calculate == "" {
				calculate += value.Calculation
			} else {
				calculate = calculate + "+" + value.Calculation
			}
			sign += value.Signs
			n -= value.Numericals

		}
	}
	fmt.Println(calculate)
	fmt.Println(sign)

}
