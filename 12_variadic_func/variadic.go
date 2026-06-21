package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments.
// The args are received as a slice.
func sum(nums ...int) int {

	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	fmt.Println(1, 2, 3, 4, 5, "Hello") //this line is aloso represent a variadic function, where the fmt.Println function can accept any number of arguments of any type.
	result := sum()
	fmt.Println(result)
}

//the variadic function can be called with any number of arguments, including zero.
// In this example,the sum function takes a variable number of integer arguments and returns their total.
// The nums parameter is treated as a slice of integers within the function, allowing us to iterate over it and calculate the sum.
