package main

import "fmt"

func main()  {

	numbers := []int{0,1,2,3,4,5,6,7,8,9}

	for i := range numbers{
		fmt.Println("Slice item",i,"is",numbers[i])
	}

	countryCapitalMap := map[string] string {"France": "Paris", "Japan": "Tokyo", "Turkey": "Ankara"}
	items := map[string] int {"Item1": 5, "Item2": 10}

	for country := range countryCapitalMap{
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	for country, capital := range countryCapitalMap{
		fmt.Println("Capital of",country,"is",capital)
	}

	for item, count := range items{
		fmt.Println(item,":",count)
	}
}
