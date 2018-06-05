package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("i = %d\n", i)
	}

	/**
	iterate through array
	**/
	num := []int{1,2,3,4,5}
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	fmt.Printf("Sum = %d\n", sum)

	
}