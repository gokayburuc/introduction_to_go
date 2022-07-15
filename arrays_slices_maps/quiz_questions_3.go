package main

import "fmt"

func main() {
	x:=[6]string{"a","b","c","d","e","f"}

	// x[2:5]
	t := x[2:5]
	fmt.Println(t) // from second element which is included to fifth element 
}