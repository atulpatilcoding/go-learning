package main

//in normal case the created functions are deleted after the execution of the function,
// but in closure the created functions are not deleted and they can be used even after the execution of the function.
//In go the closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// Closures are functions that can capture and access variables from their surrounding scope, even after that scope has finished executing.
	// In Go, closures are created when you define a function inside another function and the inner function references variables from the outer function.
	// Example of a closure in Go:

	// 	func outer() func() {
	// 		message := "Hello, World!"
	// 		return func() {
	// 			println(message) // This inner function is a closure that captures the 'message' variable from the outer function.
	// 		}
	// 	}
	// 	closureFunc := outer() // Call the outer function to get the closure function
	// 	closureFunc() // Call the closure function, which will print "Hello, World!"

	// // --------------------1------------------------

	increment := counter() // Create a closure that maintains its own state
	println(increment())
	println(increment())
	println(increment())
}
