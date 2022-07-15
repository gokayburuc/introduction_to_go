package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 3:
			fmt.Println("Three")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}
}
