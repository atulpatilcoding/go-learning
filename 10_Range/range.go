package main

import "fmt"

func main() {
	//range is used  for iterating over elements in arrays , slices ,maps ,strings and channels

	//nums := []int{1, 2, 3, 4, 5, 6, 7} //slice created

	//usung normal for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//using range in for loop
	// for i, num := range nums { //here "num" is variable that retuns value at tha index "i", range is key word that is itreting over the "nums" slice
	// 	fmt.Println(i, num)
	// }
	//-----------------------1-----------------------
	//range kyword using on maps

	// m := map[string]string{"Name": "Atul", "City": "Sangli", "Country": "India"}
	// for k, v := range m { // Here rangre is iterating over map "m" and key is "k" and "v" is value // this is commanly using way to iterate over map
	// 	fmt.Println(k, v)
	// }

	// for k := range m { // if we want only keys from map we can use this way
	// 	fmt.Println(k)
	// }
	//-----------------------2-----------------------

	//Range using on strings
	// here "i" is the byte index of the rune in the string, not the character index.
	for i, ch := range "golang" { //here we are using "i" is index and "ch" is character at that index
		fmt.Println(i, ch) // here ch is retuning unicode value of character

	}
	for i, ch := range "golang" { //
		fmt.Println(i, string(ch)) // here we are converting unicode value to string using string() function // string() is type conversion function
	}
}
