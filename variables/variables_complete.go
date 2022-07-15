package main

import "fmt"

// variables

var name string
var surname string = "Buruc"

// age := 25 statement only works in functions
// name = "Gokay" also

var (
	country = "Turkey"
	city    = "Bursa"
)

func main() {
	age := 35
	name = "GOKAY"
	fmt.Println(country, "/", city)
	fmt.Println(name, " ", surname)
	fmt.Println(age)
}
