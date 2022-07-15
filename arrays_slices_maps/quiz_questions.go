// how do you access the 4th element of an array or slice

package main

import "fmt"

func main() {
	var x[10]int
	x[5]=25
	fmt.Println(x)

	//4th element of x
	fmt.Println(x[4])
}