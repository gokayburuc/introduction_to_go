package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			//divisible by 2
			fmt.Println(i, "divisible by 2")
		} else if i%3 == 0 {
			//divisible by 3
			fmt.Println(i, "divisible by 3")
		} else if i%5 == 0 {
			//divisible by 5
			fmt.Println(i, "divisible by 5")
		} else if i%7 == 0 {
			fmt.Println(i, "divsible by 7")
		} else {
			fmt.Println(i, "other")
		}
	}
}
