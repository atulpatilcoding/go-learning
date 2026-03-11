package main

import "fmt"

func main() {
	var name string
	var age int
	var percentage float64
	var isPassed bool
	//string
	fmt.Println("Atul Suhas Patil")
	//int
	fmt.Printf("age= %d, %d ,%d \n", 23, 20, 40)
	//float
	fmt.Printf("perentage=%f \n", 62.66)
	//boolean
	fmt.Printf("is persion passed=%t \n", true)

	// taking input from user
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Printf("My name is %s \n", name)

	fmt.Print("Age: ")
	fmt.Scan(&age)
	fmt.Printf("My Age is %d \n", age)

	fmt.Print("Percentage: ")
	fmt.Scan(&percentage)
	fmt.Printf("My percentage is %f \n", percentage)

	if percentage >= 35 {
		isPassed = true
	} else {
		isPassed = false
	}
	fmt.Printf(" Person is pass = %t \n", isPassed)

}
