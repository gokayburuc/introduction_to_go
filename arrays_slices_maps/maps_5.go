package main

import "fmt"

func main(){
	city := map[string]map[string]string{
		"Istanbul":map[string]string{
			"Area":"Marmara",
			"Mayor":"Ekrem İmamoglu",
	},
		"Bursa":map[string]string{
			"Area":"Marmara",
			"Mayor":"Alinur Aktas",
		},
	}

	fmt.Println(city)
}