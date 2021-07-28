package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
	1. A rune is an alias to the int32 data type. It represents a Unicode code point.
	2. A Unicode code point or code position is a numerical value that is usually used
	   to represent a Unicode character.
	3. The int32 is big enough to represent the current volume of 140,000 unicode characters.
	*/
	a1 := '🦍'
	var a2 = 'k'
	var a3 rune = '🦡'

	fmt.Printf("%c - %s\n", a1, reflect.TypeOf(a1))
	fmt.Printf("%c - %s\n", a2, reflect.TypeOf(a2))
	fmt.Printf("%c - %s\n", a3, reflect.TypeOf(a3))
}
