package main

import "fmt"

func main() {
	fmt.Println(ConcatAlternate([]int{4, 5, 6, 1, 2, 3}, []int{1, 3, 7, 8, 10, 4, 5, 6}))
}

func ConcatAlternate(t, s []int) []int {
	var nb []int
	if len(t) > len(s) {
		out := len(t) - len(s)
		nb = append(nb, t[:out]...)
		for i := 0; i < len(s); i++ {
			nb = append(nb, t[out+i])
			nb = append(nb, s[i])

		}
		return nb

	} else if len(s) > len(t) {
		out := len(s) - len(t)
		nb = append(nb, s[:out]...)
		for i := 0; i < len(t); i++ {
			nb = append(nb, s[out+i])
			nb = append(nb, t[i])

		}
		return nb

	} else {
		times := len(s)
		for i := 0; i < times; i++ {
			nb = append(nb, s[i])
			nb = append(nb, t[i])

		}
		return nb
	}
}
