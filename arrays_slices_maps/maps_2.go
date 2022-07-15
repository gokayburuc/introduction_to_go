package main
import "fmt"


func main(){
	// var x map[string]int
	x :=make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"]) // value of "key" key
	fmt.Println(x) // shows item enterily
}