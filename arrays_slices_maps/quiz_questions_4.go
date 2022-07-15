package main

import "fmt"

func main() {
	x := []int{22, 13, 45, 7, 10, 3, 5, 39, 97}
	min, max := MinMaxFind(x)
	fmt.Println(x)
	fmt.Println(min)
	fmt.Println(max)
}

// find the min ,max elements in array

func MinMaxFind(array []int) (int, int) {
	min := array[0]
	max := array[0]

	// TODO: For loop
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
