package quizes

import "fmt"

func PrintMemory(arr [10]byte) {
	
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
			bs = string((v%16)+ '0') + bs
			v /= 16
		}
		bsArr = append(bsArr, bs)
		bs = ""
	}
	fmt.Println(bsArr)
}
