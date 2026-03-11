package main

import "fmt"

// A struct is used to create a collection of members of different data types, into a single variable..
// Structs are used to store multiple values of different data types into a single variable.
// --------------------1------------------------
// type Person struct {
// 	name   string
// 	age    int
// 	job    string
// 	salary int
// }

// func PersonPrint(pers Person) {
// 	fmt.Println("Name: ", pers.name)
// 	fmt.Println("Age: ", pers.age)
// 	fmt.Println("Job: ", pers.job)
// 	fmt.Println("Salary: ", pers.salary)
// }

// func main() {
// 	var pers1 Person
// 	var pers2 Person

// 	// Pers1 specification
// 	pers1.name = "Hege"
// 	pers1.age = 45
// 	pers1.job = "Teacher"
// 	pers1.salary = 6000

// 	// Pers2 specification
// 	pers2.name = "Cecilie"
// 	pers2.age = 24
// 	pers2.job = "Marketing"
// 	pers2.salary = 4500

// 	// here we are calling the function PersonPrint of Pers1 info
// 	PersonPrint(pers1)
// 	fmt.Println(pers1) // here we are printing the whole struct of Pers1 info

// 	// here we are calling the function PersonPrint of Pers2 info
// 	PersonPrint(pers2)
// 	fmt.Println(pers2) // here we are printing the whole struct of Pers2 info
// }

// --------------------2------------------------
func main() {
	var Name string
	fmt.Scan("%s", &Name)
	fmt.Println("Hello", Name)
}
