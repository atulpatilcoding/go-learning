package main

import "fmt"

func main() {
	// Arrays are numbered sequences that store multiple values of the same type ,
	// Allowing you to group related data together. In Go, arrays have a fixed size defined at the time of declaration.
	// Arrays have fixed size means once an array is created, its size cannot be changed.
	//memory optimization

	// --------------------1------------------------
	// var arr[5]int
	// arr[1]=10
	// arr[2]=20
	// arr[3]=30
	// fmt.Println(arr)
	// fmt.Println(arr[2])
	// fmt.Print(len(arr))

	// --------------------2------------------------
	// zeroed values of different types
	// var vals[2]bool
	// var vals2[2]int
	// var vals3[2]float32
	// var vals4[2]string
	// var vals5[2]uint8
	// fmt.Println(vals)
	// fmt.Println(vals2)
	// fmt.Println(vals3)
	// fmt.Println(vals4)
	// fmt.Println(vals5)

	// --------------------3------------------------
	//  to declare arr in single line with short hand mathod
	// num:=[3]int{10,20,30}
	// fmt.Println(num)

	//2d array
	// nums:= [2][3]int{{10,20,30},{40,50,60}}
	// fmt.Println(nums)

	//--------------------4------------------------
	// here we are declaring arr with out giving size
	// here size will be decided by no of elements in the array
	// this is also fiex not dinamic sized array
	arr2 := [...]int{9, 7, 6, 4, 5, 3, 2, 4}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

}
