package main

import "fmt"

func main() {
	x := 5
	f(x)
}

// we have to define
func f(x int) {
	fmt.Println(x)
}
