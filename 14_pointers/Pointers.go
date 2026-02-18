package main

import "fmt"

//Pass by value
func changeNum(num int) { // num is a copy of the value passed in the main function , so changes to num do not affect the original variable in main.
	num = 5                           // we try to change the value of num to 10 but it does not affect the original variable in main because num is a copy of the value passed in the main function.
	fmt.Println("In chngeNum :", num) // but in ChangeNum function it will print 10 because we have changed the value of num to 10 in the ChangeNum function.
}

//Pass by reference
func changeNumByReference(num *int) { // here "*" sign is used to declare a pointer variable num which can hold the memory address of an integer variable, so changes to num will affect the original variable in main because num is a pointer to the original variable in main.
	*num = 10 // we change the value of num to 10 using the dereference operator * ,
	// so it will affect the original variable in main because num is a pointer to the original variable in main.
	fmt.Println("In changeNumByReference :", *num) // but here "*" sign is used to dereference the pointer variable num and access the value it points to, so it will print 10 because we have changed the value of num to 10 in the ChangeNumByReference function.
}

func main() {
	// A pointer is a variable that holds the memory address of another variable.
	// In Go, you can use the & operator to get the memory address of a variable and the * operator to dereference a pointer and access the value it points to.
	num := 1
	changeNum(num)
	fmt.Println("In side main :", num) //no changre in num because we are passing the value of num to the changeNum function, so it does not affect the original variable in main.
	changeNumByReference(&num)         //here "&" sign is used to get the memory address of num and pass it to the changeNumByReference function, so it will affect the original variable in main because we are passing the memory address of num to the changeNumByReference function.
	fmt.Println("In side main :", num) // it will print 10 because we are passing the memory address of num to the changeNumByReference function, so it will affect the original variable in main.

}
