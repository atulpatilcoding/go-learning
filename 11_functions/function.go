package main

import (
	"fmt"
)

func add(a int, b int) int {
	//func add(a ,b int) int { // you can also write like this
	return a + b
}


//in multiple conditionswe can return multiple values from a function in Go.
//in go first value is the return value and second value is the error value. we can also return a boolean value to check if there is an error or not. for example, if we want to get the language of a person we can return the language as a string and the error as a string and the boolean value to check if there is an error or not.

func getLanguage() (string, string, bool) { // if you want to return multiple values
	return "Go", "Java", true
}
func main() {

	e := 50
	d := 70
	result := add(e, d)
	fmt.Println(result)

	// multiple return values
	lang1, lang2, lang3 := getLanguage()
	fmt.Println(lang1, lang2, lang3)
	fmt.Println(getLanguage()) // you can also print like this but it will print in the form of tuple
	//--------------------------------

	// Anonymous function
	func() {
		fmt.Println("This is an anonymous function")
	}() // you have to call the function immediately after defining it

}
