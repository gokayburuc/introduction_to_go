package main

import "fmt"

func main(){
	x:=make(map[int]int)

	x[1]=25
	fmt.Println(x[1])
	fmt.Println(len(x)) //length of map
	fmt.Println(len(x)) //length of map
}