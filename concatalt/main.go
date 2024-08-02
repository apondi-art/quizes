package main

import "fmt"

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {
	var nb []int
	if len(slice1) > len(slice2) {
		for i := 0; i < len(slice1); i++ {
			if i > len(slice2)-1 {
				nb = append(nb, slice1[i])
			} else {
				nb = append(nb, slice1[i])
				nb = append(nb, slice2[i])

			}
		}
		return nb

	} else if len(slice2) > len(slice1) {
		for i := 0; i < len(slice2); i++ {
			if i > len(slice1)-1 {
				nb = append(nb, slice2[i])
			} else {
				nb = append(nb, slice2[i])
				nb = append(nb, slice1[i])

			}
		}
		return nb

	} else {
		times := len(slice2)
		for i := 0; i < times; i++ {
			nb = append(nb, slice1[i])
			nb = append(nb, slice2[i])

		}
		return nb
	}
}
