
package main
import "fmt"

func main(){
	//TODO: Find the usage of 3,9 in golang slices
	test_slice := make([]int,3,9) 
	// length of test slice 
	fmt.Println(len(test_slice))
	//slice view 
	fmt.Println(test_slice)
}