package main

import (
	"fmt"
	//"slices"   // importing slices package for using its methods
)

func main() {
	//slices are dinamic arrays,where upend and delete the elementes while run time it can grow and shrink acordingly
	//slices are refrence types
	//slices are built on top of arrays
	//slices have three components: pointer,length and capacity
	//pointer: it points to the first element of the array that is accessible by the slice
	//length: it is the number of elements present in the slice
	//capacity: it is the number of elements in the underlying array,starting from the index which is pointed by the pointer

	//-----------------------1-----------------------
	//uninitialized slices are nil
	// var s[]int
	// fmt.Println(s==nil)
	// fmt.Println(len(s))

	//-----------------------2-----------------------
	//here we are initializing slices with initial value of 2 and capacity of 3
	//in this is not a fixed values it is flexible ,it is used for "we know that we need atleast 2 values but can grow upto 3 values"
	//if we try to add more than 3 values it will automaticlly increase the capacity

	// var s= make([]int,2,3) //here length is 2 and capacity 3
	// fmt.Println(s==nil)
	// fmt.Println(len(s)) //initial length is 2
	// fmt.Println(cap(s)) //initial capacity is 3
	// fmt.Println(s)	//initially there are elements with zero value
	// s=append(s,10) // adding element 10
	// s=append(s,20)	 // adding element 20
	// fmt.Println(len(s)) //length is now 4
	// fmt.Println(cap(s)) //capacity is now 6 as it automaticlly increased it increased the capacity to double the previous capacity
	// fmt.Println(s) 	//now elements are [0 0 10 20]
	// fmt.Println(len(s)) //length is now 4
	// fmt.Println(cap(s)) //capacity is now 6
	// s=append(s,30,40,50) // adding more elements
	// fmt.Println(len(s))

	//-----------------------3-----------------------

	// num:=[]int {} //we can also initialize slices like this
	// num = append(num,1)
	// num =append(num,2,3,4,5)
	// fmt.Println(num)

	//-----------------------4-----------------------

	// copying slices in another empty slice
	// var num=make([]int ,2,5 )
	// num [0]=10
	// num=append (num,11)
	// num=append (num,15)

	// var num2=make([]int,len(num),cap(num))
	// copy(num2,num)
	// fmt.Println(num,num2)

	//-----------------------5-----------------------
	//if we neede a part of the slice we can use like this
	// slice operator
	// var num=[]int {10,20,30,40,50,60,70,80,90}
	// fmt.Println(num[3:6]) // from index 3 to index 5
	// fmt.Println(num[:6]) // starting from index 0 to index 5
	// fmt.Println(num[3:]) // starting from index 3 to the end

	//-----------------------5-----------------------
	//slices package

	// var num=[]int {10,20,30,40,50,60,70,80,90}
	// var add=[]int {1,2,3}
	// var sum=[]int {1,2,3}
	// fmt.Printf("%vn\%v\n%v\n", num, add, sum)
	// // Mehods from slices package
	// fmt.Println(slices.Equal(add,sum))	// Equal method from slices package
	// fmt.Println(slices.Compare(num,sum)) // Compare method from slices package
	// fmt.Println(slices.Concat(add,sum))
	// fmt.Println(slices.Contains(num,50))
	// fmt.Println(slices.Index(num,70))

	//-----------------------6-----------------------

	//2D slices
	var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
}
		