package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	sentence:= "The quick brown fox jumps over"
	parts := strings.Split(sentence," ")//
	for index:=range parts{
		fmt.Println(parts[index])
		//println()
	}

	var array [5] int
	abc:= [] int {1,2,3,4,5}
	 array[0] = 1
//	b := array[1] == 2
	fmt.Println(array[1])
	 fmt.Println(abc)
	 fmt.Printf("%i", abc)
	 println("...................")
	 fmt.Println(len(abc))

	for i:= range abc{
		fmt.Println(i)
	}
	for i:=0; i<len(abc); i++ {
		fmt.Println(i)
	}

	array2D:= [][] int {{0,0},{3,4} ,{3, 9}}
	fmt.Println(array2D[2][1])

	//str:="toluwase"

//	fmt.Printf(strings.Split(str,"w"))

	stri:= "i love jesus"
	strings.Split(stri, " ")
	a:=strings.Contains(stri, "o")
	//strings.Replace(stri, " ")
	if a == true {
		 b:=strings.ToUpper(string(stri[1]))
		 fmt.Println(strings.ToUpper(b))
	} else {
		fmt.Println("does not contain O")
	}

	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	println()
	fmt.Println(strings.Trim("Toluwase", " "))
}

func fizzBuzz(n int32) {
	// Write your code here
	for i:=0; i<len(n); i++ {
		if i%3==0 && i%5==0 {
			println("FizzBuzz")
		} else if i%3==0 && i%5==1 {
			println("Fizz")
		} else if i%3==1 && i%5==0 {
			println("Buzz")
		} else {
			println(i)
		}
	}

}