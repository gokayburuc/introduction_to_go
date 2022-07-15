package main

import "fmt"

func main(){
	capitals:=make(map[string]string)

	capitals["Albania"]= "Tirana"
	capitals["Andorra"]= "Andorra la Vella"
	capitals["Austria"]= "Vienna"
	capitals["Belarus"]= "Minsk"
	capitals["Belgium"]= "Brussels"
	capitals["Bosnia and Herzegovina"]= "Sarajevo"
	capitals["Bulgaria"]= "Sofia"
	capitals["Croatia"]= "Zagreb"
	capitals["Czechia"]= "Prague"
	capitals["Denmark"]= "Copenhagen"
	capitals["Estonia"]= "Tallinn"
	capitals["Finland"]= "Helsinki"
	capitals["France"]= "Paris"
	capitals["Germany"]= "Berlin"
	capitals["Greece"]= "Athens"
	capitals["Hungary"]= "Budapest"
	capitals["Iceland"]= "Reykjavik"
	capitals["Ireland"]= "Dublin"
	capitals["Italy"]= "Rome"
	capitals["Latvia"]= "Riga"
	capitals["Liechtenstein"]= "Vaduz"
	capitals["Lithuania"]= "Vilnius"
	capitals["Luxembourg"]= "Luxembourg"
	capitals["Malta"]= "Valletta"
	capitals["Moldova"]= "Chisinau"
	capitals["Monaco"]= "Monaco"
	capitals["Montenegro"]= "Podgorica"
	capitals["Netherlands"]= "Amsterdam / Haag"
	capitals["North Macedonia"]= "Skopje"
	capitals["Norway"]= "Oslo"
	capitals["Poland"]= "Warsaw"
	capitals["Portugal"]= "Lisbon"
	capitals["Romania"]= "Bucharest"
	capitals["Russia"]= "Moscow"
	capitals["San Marino	San "]= "Marino"
	capitals["Serbia"]= "Belgrade"
	capitals["Slovakia"]= "Bratislava"
	capitals["Slovenia"]= "Ljubljana"
	capitals["Spain"]= "Madrid"
	capitals["Sweden"]= "Stockholm"
	capitals["Switzerland"]= "Bern"
	capitals["Ukraine"]= "Kiev"
	capitals["United Kingdom"]= "London"


	fmt.Println(capitals["Spain"])

	// check for zero values
	// (capitals["Spain"] == )
	name,ok := capitals["Spain"]
	fmt.Println(name,ok)


	//if condition zero value check
	if name,ok := capitals["Spain"]; ok {
		fmt.Println(name,ok)
	}
}
