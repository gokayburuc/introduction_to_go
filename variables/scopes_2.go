package main

import "fmt"

var x string = "Hello World"

func main() {
	fmt.Println(x)
	fmt.Println("main function")
	f()

}

func f() {
	fmt.Println(x)
	fmt.Println("f function")

}
