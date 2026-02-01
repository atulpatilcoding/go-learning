package main

import "fmt"

func main() {

	// age := 1
	// if age>18 {
	// 	fmt.Println("Adult")
	// } else if  age>10 && age<20 {
	// 	fmt.Println("teen ager")
	// }else{
	// fmt.Println("kid")
	// }

	// var role = "Admin"
	// var haspermission =false

	// if role == "Admin" && haspermission{
	// 	fmt.Println("Access granted")
	// }else {
	// 	fmt.Println ("Access denied")
	// }

	//we can declare an veriable in side "if" condition
	if age := 12; age >= 18 {
		fmt.Println("Adult")
	} else if age > 10 && age < 20 {
		fmt.Println("teen ager")
	}

	// go does not have ternary operator in the case of we have to use if else

}
