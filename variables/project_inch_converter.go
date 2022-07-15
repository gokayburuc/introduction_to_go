package main

import "fmt"

const inch = 2.54

func main() {
	var input float64
	fmt.Println("Enter inch: " )
	fmt.Scanf("%f", &input)
	inc_to_cm := inch * input
	fmt.Println(input , " inch = ", inc_to_cm)
}