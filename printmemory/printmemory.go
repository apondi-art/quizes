package quizes

import "fmt"

func PrintMemory(arr [10]byte) {
	hex := "0123456789ABCDEF"
	var str string
	var bs string
	var bsArr []string
	for _, v := range arr {
		if v >= 32 && v <= 126 {
			str += string(v)
		} else {
			str += "."
		}
		for v > 0 {
			bs = string(hex[v%16]) + bs
			v /= 16
		}
		bsArr = append(bsArr, bs)
		bs = ""
	}
	count := 0
	for i := 0; i < len(bsArr); i++ {
		if bsArr[i] == "" {
			bsArr[i] = "00"
		}
		count++
		if count%4 == 0 {
			fmt.Print(" ", bsArr[i])
			fmt.Println()
			count = 0
		} else if count == 1 {
			fmt.Print(bsArr[i])
		} else{
			fmt.Print(" ", bsArr[i])
		}
	}
	fmt.Println()
	fmt.Println(str)
}
