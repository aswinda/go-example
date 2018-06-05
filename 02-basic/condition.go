package main

import "fmt"

func main() {
	var (
		x = 5
		y = 10
	)

	if x < y {
		fmt.Printf("x is less than y.\n")
	}

	/**
	if accept an initialization statement
	**/
	if x = 15; x > y {
		fmt.Printf("x is greater than y.\n")
	}

	/**
	else statement
	**/
	if y = 20; x > y {
		fmt.Printf("x is greater than y.\n")
	} else {
		fmt.Printf("x is less than y.\n")
	}
}