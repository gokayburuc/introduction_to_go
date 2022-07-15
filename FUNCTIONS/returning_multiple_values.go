package main

import "fmt"

func main() {
	x, y := f()

	fmt.Println(x)
	fmt.Println(y)
}

// TODO: function writing syntax
func f() (int, int) {
	return 5, 6
}
