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
	Rune is alias for int32. It is an UTF-8 encoded code point
	**/
	c := []rune(s)	    
	c[0] = 'h'	        
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}
