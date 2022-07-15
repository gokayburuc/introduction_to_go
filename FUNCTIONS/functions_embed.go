package main

import "fmt"

func main() {
	fmt.Println(one())
}

func one() int {
	return two()
}

// func two() int {
// 	return 1
// }

func two() (r int) {
	r = 1
	return
}
