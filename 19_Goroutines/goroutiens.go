package main

import (
 "fmt"
 "time"
)

//------------------------1------------------------
//here we are just creating a function "task" and using this in main function
//in this example the "task" func wait for complete the itration of for loop and after a loop execution it gose for next loop or itration
// func task(id int) {
//  fmt.Println("Doing Task ", id)
// }

// func main() {
//  for i := 0; i < 11; i++ {
//   task(i) //just calling the "task" func
//  }

// }

//------------------------2------------------------

// //here we are using go routines concept
// //in go routines the func is not wait for the loop itration it creates light wait threds that therds are run Parallely that means no waiting and faster results
// // the threds running parllely depends on your local machins  howmany cores are present in your processor and all
// //but it is faster processing

// func task(id int) {
//  fmt.Println("Doing Task ", id)
// }

// func main() {
//  for i := 0; i < 11; i++ {
//   go task(i) //here we are calling the "task" func and implementing go routiens concept
//  }
//  time.Sleep(time.Second*2) // we are holding the main function form shutting down , becose the main function is terminating before the output comes
//  //this is not most time efficient method or idial method but we are using for learning over here
//  // the output is not in order becose multiple porocesses are running concurrtly with threds they are not wating for each other for order
// }

//------------------------3------------------------
//go routiens we can also use for in line func we creat annonomus func and use it like this

func main() {
 for i := 0; i < 11; i++ {

  go func(i int) { //we created here annonomus func to use go routines
   fmt.Println("Doing Task ", i)
  }(i)
 }
 time.Sleep(time.Second * 2)
}
