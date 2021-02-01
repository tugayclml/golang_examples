package main

import "fmt"

func main()  {

	var a int = 100
	var b int = 200
	var ret int

	ret = max(a,b)
	fmt.Printf("Max value is: %d\n", ret)
	
}

func max(a int, b int) int{
	/* Local variable declaration */
	var result int
	
	if a>b {
		result = a
	}else if b > a {
		result = b
	}

	return result
}