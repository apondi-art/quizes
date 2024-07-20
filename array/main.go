package main

import "fmt"

func main() {
	s := []int{3, 4, 7, 2, 8, 1}
	//number := 4
	fmt.Println(ArrangeArrayInOrder(s))
}

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i <= len(numbers); i++ {
		for j := i + 1; j <= len(numbers)-1; j++ {
			fmt.Println(numbers[i], numbers[j]) //how to access  values of array
			v := numbers[i] + numbers[j]
			if v == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}
func ArrangeArrayInOrder(t []int) []int {
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i] > t[j] {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
	return t
}
func CheckForFactor(base int, factor int) bool {

	return base%factor == 0
}
