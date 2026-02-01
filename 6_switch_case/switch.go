package main

import (
	"fmt"
//	"time"
)

func main() {

	// --------------------1------------------------
	// var cou int
	// if age := 13; age > 0 && age < 10 {
	// 	cou = 1
	// } else if age > 10 && age < 18 {
	// 	cou = 2
	// } else if age >= 18 {
	// 	cou = 3
	// } else {
	// 	cou = 0
	// }

	// switch cou {
	// case 1:
	// 	fmt.Println("kid")
	// case 2:
	// 	fmt.Println("teenager")
	// case 3:
	// 	fmt.Println("Adult")
	// default:
	// 	fmt.Println("Wrong gae")
	//no need to Write break keyword it atomatically breaks after each case
	// }

	// --------------------2------------------------
	//multiple conditiond switch in a case

	// switch time.Now().Weekday() {

	// case time.Sunday, time.Saturday:
	// 	fmt.Printf("%s Weekend", time.Now().Local().Weekday())
	// default:
	// 	fmt.Printf("%s Weekday",time.Now().Local().Weekday())
	// }


	// --------------------3------------------------
	Whoami :=func(i interface{}) {
		switch i.(type){
		case int:
			fmt.Printf("int")
		case string:
			fmt.Printf("string")
		case bool:
			fmt.Printf("bool")
		default:
			fmt.Printf("unknown")
		}

	}
	Whoami(true)
}
