package main

import "fmt"

func sum(x, y int) (int, int) {
	result := x + y
	return result, 10
}

func main()	{

	a, b := sum(1, 5)

	fmt.Println(a, b)
}