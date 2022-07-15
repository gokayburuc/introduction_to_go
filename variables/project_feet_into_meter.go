package main
import "fmt"

const feet = 0.3048

func main() {

	var feet_input float64
	fmt.Println("Enter Feet : ")
	fmt.Scanf("%f", &feet_input)

	feet_to_meter := feet * feet_input
	fmt.Println(feet_to_meter)
}