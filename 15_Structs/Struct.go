package main

import "fmt"

//"structs"
//"time"

// A struct is used to create a collection of members of different data types, into a single variable..
// Structs are used to store multiple values of different data types into a single variable.
// if we, want's to group a set of veriables and constents
// --------------------1------------------------
// type Person struct {
// 	name   string
// 	age    int
// 	job    string
// 	salary int
// }

// func PersonPrint(pers Person) {
// 	fmt.Println("Name: ", pers.name)
// 	fmt.Println("Age: ", pers.age)
// 	fmt.Println("Job: ", pers.job)
// 	fmt.Println("Salary: ", pers.salary)
// }

// func main() {
// 	var pers1 Person
// 	var pers2 Person

// 	// Pers1 specification
// 	pers1.name = "Hege"
// 	pers1.age = 45
// 	pers1.job = "Teacher"
// 	pers1.salary = 6000

// 	// Pers2 specification
// 	pers2.name = "Cecilie"
// 	pers2.age = 24
// 	pers2.job = "Marketing"
// 	pers2.salary = 4500

// 	// here we are calling the function PersonPrint of Pers1 info
// 	PersonPrint(pers1)
// 	fmt.Println(pers1) // here we are printing the whole struct of Pers1 info

//		// here we are calling the function PersonPrint of Pers2 info
//		PersonPrint(pers2)
//		fmt.Println(pers2) // here we are printing the whole struct of Pers2 info
//	}
//
//	**** A function is used to perform a specific task, while a struct is used to store related data together ****
//
// --------------------2------------------------
// type order struct {
// 	id        string
// 	amount    float32 //use "Ctrl + Space " for auto complition
// 	status    string
// 	createdAt time.Time // nano sec precision //in this line The "createdAt" is veriable And we are using "time" key word of ".time" package like "Println" keyword from "fmt" package
// } // this sturuct is just a blue-print of instence using this structure we can make multiple instences

// func main() {
// 	myOrder := order{ // "myOrder" is the instence
// 		id:     "1",
// 		amount: 50,
// 		status: "Received",
// 	}
// 	myOrder.createdAt = time.Now()
// 	fmt.Println("Order Struct", myOrder)

// 	fmt.Println("Order ID:", myOrder.id)
// 	fmt.Println("Order Amount:", myOrder.amount)
// 	fmt.Println("Order Status:", myOrder.status)
// 	fmt.Println("Order Time:", myOrder.createdAt)

// 	var myOrder2 = order{
// 		id:        "2",
// 		amount:    100,
// 		status:    "deleverd",
// 		createdAt: time.Now(),
// 	}

// 	myOrder.status = "paid"
// 	fmt.Println("Order Struct", myOrder)
// 	fmt.Println("Order Struct", myOrder2)
// }
// --------------------3------------------------

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time
// 	//in Object Oriyented Programming we atach the methods to the class (behavior / dose part ) like this we can also do in go using "Struct"
// }

// // receiver type
// func (o *order) changeStatus(status string) { //"(o order)"here we are changing in the copy of sruct thats why we using pointer "(o *order)"
// 	//to connect struct "order" to this function we use "(o order)"
// 	//here we can use any other letter instd of "o"  but acording to the convention that followed by go dev is use first letter of struct we are using "o"
// 	o.status = status // here we need to de reference the poiner but we don't need to write "*o.status=status" go automatically retuns the dereference value
// }

// //like this we use function to get amount

// func (o order) getamount() float32 { //here we are not getting any thing in the function but returning float32 value from the function
// 	return o.amount //here we are not using pointer becouse we are not modifing the struct's value we just getting the value from struct and returning
// }
// func main() {
// 	var myOrder = order{
// 		// id: "1",  //if we not give any value in struct then we recive the zero valu fo that data type
// 		//if dont set any field, default value is zero value (int=>0,string=>"",bool=>false)
// 		amount:    100,
// 		status:    "deleverd",
// 		createdAt: time.Now(),
// 	}

//		fmt.Println(myOrder)
//		myOrder.changeStatus("confirmed")
//		fmt.Println(myOrder)
//		fmt.Println(myOrder.getamount())
//	}
//
// --------------------4------------------------
// constructor in go
// in golang there is no constructor is avelable but usinf struct and func we can use constructor

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time
// } // here we are telling the function is returning poin of "order" structure
// func newOrder(id string, amount float32, status string) *order { //when we are using func as constructor the convention is use "new" word befor it in this case "order" is structure so "newOrder"
// 	//here you can setup your initial packages and all
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder //for returning pointer we use here "&"" sign
// }

// func main() {
// 	myOrder := newOrder("1", 30.50, "received")
// 	fmt.Println(myOrder)
// }
// --------------------5------------------------

// till here we use struct globally becose need that struct multiple instances  times
// here we are learning if we need struct one time and inside the function
//in line struct

// func main() {

// 	language := struct {
// 		name   string
// 		isGood bool
// 	}{"golang", true}

// 	fmt.Println(language)

// }
// --------------------5------------------------
//steruct imbeding /Nested struct / inheretence /Composition
type customer struct {
	name  string
	phone string
}
type order struct {
	id     string
	amount float32
	status string
	customer
}

func main() {
	myOrder := order{
		id:     "1",
		amount: 50,
		status: "received",
		customer: customer{name: "John",
			phone: "1234567890"},
	}

	fmt.Println(myOrder)
	myOrder.customer.name = "rohan"
	fmt.Println(myOrder.customer)
}
