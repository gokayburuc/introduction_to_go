package main

import (
	"fmt"
)

var x string = "hello"
var y string = "world"

func main() {
	x += y // x = x + y
	fmt.Println(x)

	fmt.Println( x == y )
	fmt.Println (x != y )


	z := "bye bye world"
	fmt.Println(z)

	var t = "Mahmut abi"
	fmt.Println(t)
}