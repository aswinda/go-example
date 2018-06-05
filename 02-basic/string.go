package main

import "fmt"

func main() {
	var s string
	s = "Hello"
	fmt.Printf("%s\n", s)

	/**
	this will give you error
	string in Go are immutable
	**/
	// s[1] = "a"

	/**
	you change change string by converting in array 
	Rune is alias for int32
	**/
	c := []rune(s)	    1
	c[0] = 'c'	        2
	s2 := string(c)     3
	fmt.Printf("%s\n", s2) 4
}
