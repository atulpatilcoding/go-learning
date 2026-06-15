package main

import "fmt"

//----------------------1----------------------
//we are here repeating the same code for different types, we can use generics to avoid this repetition

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlices(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func main() {
// 	//nums := []int{1,2,3,4,5}
// 	//printSlice(nums)
// 	names := []string{"Golan", "Java", "Typescript ", "Python"}
// 	printStringSlices(names)
// }

//----------------------2----------------------
//we can use generics to avoid this repetition in function  
//we can define a generic function that can work with any type of slice
func printSlice[T any](items []T) { //insted of "[T any]" to allow any type we can also use "[T string | int | bool]" to allow only specific types
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	//nums := []int{1,2,3,4,5}
	names := []string{"Golan", "Java", "Typescript ", "Python"}
	vals := []bool{true, false, true, true}
	//printSlice(nums)
	printSlice(names)
	printSlice(vals)
}

//----------------------3----------------------
//Like function, we can also use generic types to struct. This allows us to create data structures that can work with any type of data. 
// For example, we can define a generic stack data structure that can hold any type of data:


// type Stack struct {
// 	element []int  //we have to repeat the same code for different types,so we can use generics to avoid this repetition
// }

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
