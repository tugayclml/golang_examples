package main

import "fmt"

var g int = 20

func main()  {
	fmt.Println(g)
	g -= 10
	fmt.Println(g)
	changeGValue()
	fmt.Println(g)
}

func changeGValue()  {
	g += 100
}
