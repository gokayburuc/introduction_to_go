package main

import "fmt"

func main() {
	arr := []float64{1, 2, 3, 4, 5} //without [5] also we can write [] emtpy square brackets
	x := arr[0:5]
	fmt.Println(x)

	y := arr[0:]
	fmt.Println(y)

	z := arr[0:len(arr)]
	fmt.Println(z)

	t := arr[:5]
	fmt.Println(t)

	p := arr[:]
	fmt.Println(p)

}
