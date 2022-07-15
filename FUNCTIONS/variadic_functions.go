package main

import "fmt"

// ... means this function can take zero or more parameters
func add(args ...int) int {
	total := 0
	// TODO: for _,v expression
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	fmt.Println(add(1, 2, 3))

	// TODO : usage of slices with variadic functions
	xs := []int{1,2,3,4,5}
	fmt.Println(add(xs...))
}
