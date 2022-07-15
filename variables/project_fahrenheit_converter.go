package main

import (
	"fmt"
)


// kullanicidan veri al

func main() {
	var celcius float64
	fmt.Println("CELCIUS : ")
	fmt.Scanln(&celcius)
	fahrenheit := (celcius * 1.8) + 32
	fmt.Println(celcius, "C = ",fahrenheit, " F")
}