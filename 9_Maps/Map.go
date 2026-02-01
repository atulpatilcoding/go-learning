package main

import (
	"fmt"
	//"maps"
)

func main() {
	//maps are used to store key value pairs, they are reference types
	//maps are built using hash tables
	//maps are dynamic in size
	//maps are created using make function

	//creating map

	// m:=make(map[string]int) //here string is key type and int is value type
	// m["a"]=1
	// m["b"]=2
	// m["c"]=3
	//-----------------------1-----------------------

	// m:= make(map[string]string)   //for int it is 0
	m := make(map[string]int) //for string it is "   "
	m["age"] = 30
	n := map[string]string{ //another way to create map
		"name": "John",
		"city": "New York",
	}

	delete(n, "city") //deleting key value pair from map
	clear(n)          //clearing all key value pairs from map
	fmt.Println(m["age"])
	fmt.Println(n["name"]) //if key not found returns zero value of value type
	fmt.Println(len(m), len(n))

	//------------------------2-----------------------
	//defalt method in go to check if key exists in map or not and get its value

	// m := map[string]int{
	// 	"a": 25,
	// 	"b": 30,
	// }
	// //ok is used to check if key exists in map
	// vla,ok := m["a"] // here we are chacking if key "a" exists in map m or not
	// //  if exists then return true in ok and get its value in val if not ok will be false and store 0 in val
	// fmt.Println(vla)
	// if ok {
	// 	fmt.Println("All ok")
	// } else {
	// 	fmt.Println("Not ok")
	// }

	//------------------------3-----------------------
	//

	// m1 := map[string]int{"phones": 5, "Prise": 10}
	// m2 := map[string]int{"phones": 5, "Prise": 10}
	// fmt.Println(maps.Equal(m1, m2)) //in maps package Equal function is used to compare two maps

	//New line added 
}
