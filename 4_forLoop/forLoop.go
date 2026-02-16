package main

import "fmt"

func main() {
	// only forloop is exists in golang

	// While loop
	// i:= 0
	// for i<5{
	// 	if i==0{
	// 		fmt.Println("this is while type for loop")
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// inpinite loop
	// for{
	// 	fmt.Println(1)
	// }

	//classic / traditional for loop
	// for i:=0; i<5; i++{
	// 	if i==0{
	// 		fmt.Println("this is classic or traditional for loop ")
	// 	}
	// 	if i==1{
	// 		continue
	// 	}

	// 	if i==3{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	//for range concept
	for i := range 5 {
		fmt.Println(i)
	}
	for i := range []int{10, 20, 30, 40, 50} {
		fmt.Println(i) // i is index
	}

	for i := range "hello" {
		fmt.Println(i)
	}

	for _, v := range []int{10, 20, 30, 40, 50} { // _ is used to ignore the index and v is used to get the value
		fmt.Println(v) // v is value and _ is index
	}

	for _, v := range "hello" {
		fmt.Println(v) //here v is value and _ is index but in string it will return the ascii value of the character
	}
}
