package main
import "fmt"

func main() {

	question := (true && false ) || (false && true) || (false && false)
	fmt.Println(question)
}