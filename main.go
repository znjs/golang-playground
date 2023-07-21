package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello, 世界")
	// STRINGS
	// var name string = "21"
	// var nameInt int = 21
	// var fName = "John"
	// fmt.Println(nameInt, name, fName)
	// lastName := "Doe"
	// fmt.Println(lastName)
	// // bits and memory

	// var num8 int8 = 127;
	// fmt.Println(num8);
	// fmt.Print("Hello, ")
	// fmt.Print("World"+"abc")

	// age := 23
	// name := "znjs"
	// * Printf -> Format String %_ -> format specifier
	// * Printf -> %q ideally should be used on strings for numbers it will be replaced with #
	// * Printf -> %T -> type of the variable
	// * Printf -> %0.1f
	// fmt.Printf("my age is %d and my name is %v \n",age, name)
	// fmt.Printf("my age is %d and my name is %c \n",age, name)
	// fmt.Printf("my age is %q and my name is %q \n",age, name)
	// fmt.Printf("my age is %T and my name is %T \n",age, name)
	// * SPRINTF returns a string

	// holder := fmt.Sprintf("my age is %d and my name is %v \n", 23, "znjs")
	// fmt.Println(holder)
	// fmt.Printf("\n %6.2f", 20.08)
	// fmt.Printf("\n %.01f", 2.0856)
	// fmt.Sprintf("\n %.01f", 23.0856)
	// *ARRAYS
	// var ages [3]int = [3]int{20, 25, 30}
	// agesTwo  := [3]int{20, 25, 30}

	// names := []string{"a", "b", "c","d"}
	// fmt.Println(names)
	// names[3] = "e"
	// fmt.Println(names)
	// names = append(names, "f")
	// fmt.Println(names)
	// fmt.Println(ages,agesTwo, len(ages))
	// fmt.Println(ages[1])

	// * RANGE SLICES
	/*
	* names[1:3] -> 1 is inclusive and 3 is exclusive
	* names[1:] -> 1 is inclusive and till end
	* names[:3] -> 0 is inclusive and 3 is exclusive
	*/
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:len(names)-1]
	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	fmt.Println(strings.Contains("hello world", "world"))
}
