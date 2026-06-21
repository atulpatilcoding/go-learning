package main

import "fmt"

// --------------------1------------------------
//here are moto is creating razorpay GW and using it by main func

// type payment struct{}

// func (p payment) makePayment(amount float32) {// we make this method for sellecting the payment gatway's
// 	razorpayPaymentGw := razorpay{}
// 	razorpayPaymentGw.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {  //we make this method for executing the payment by selected gatway
// 	// logic to make payment
// 	fmt.Println("making payment using razorpay", amount)
// }

// func main() {
// 	newPayment := payment{}
// 	newPayment.makePayment(100)
// }

// --------------------2------------------------
// //here are moto is to ceat multiple GW'sand use with the help of main

// type payment struct{}

// func (p payment) makePayment(amount float32) { // we make this method for sellecting the payment gatway's
// 	//razorpayPaymentGw := razorpay{}           // here we need to modify method when we want to change Payment GW
// 	//razorpayPaymentGw.pay(amount)             // here we need to change every time when we want diffrent payment GW (explicitly)
// 	stripePaymentGw := stripe{}
// 	stripePaymentGw.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) { //we make this method for executing the payment by selected gatway
// 	// logic to make payment
// 	fmt.Println("making payment using razorpay", amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	//logic to make payment
// 	fmt.Println("making payment using stripe", amount)
// }

// func main() {
// 	newPayment := payment{}
// 	newPayment.makePayment(100)
// }

// --------------------3------------------------
//Here we are improving the code
// what we are trying we have to change minimal code outside the "main" function

// type payment struct {
// 	gateway fakepayment //razorpay //stripe       // here we just need to change the gateway stripe or razorpay or fakepayment
// }

// func (p payment) makePayment(amount float32) { // we make this method for sellecting the payment gatway's
// 	//razorpayPaymentGw := razorpay{}           // here we need to modify method when we want to change Payment GW
// 	//razorpayPaymentGw.pay(amount)             // here we need to change every time when we want diffrent payment GW (explicitly)
// 	//stripePaymentGw := stripe{}
// 	p.gateway.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) { //we make this method for executing the payment by selected gatway
// 	// logic to make payment
// 	fmt.Println("making payment using razorpay", amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32) { //we make this method for executing the payment by selected gatway
// 	//logic to make payment
// 	fmt.Println("making payment using stripe", amount)
// }

// type fakepayment struct{}

// func (f fakepayment) pay(amount float32) {
// 	fmt.Println("making payment using fake gateway for testing pupose", amount)
// }

// func main() {
// 	//razorpayPaymentGw := razorpay{}
// 	//stripePaymentGw := stripe{}\
// 	fakeGw := fakepayment{}

// 	newPayment := payment{
// 		// gateway: stripePaymentGw,
// 		// gateway: razorpayPaymentGw,
// 		gateway: fakeGw,
// 	}
// 	newPayment.makePayment(100)
// }

// --------------------4------------------------
//Here we are implementing interface "paymenter"
// why we are using interface-> we not need to change gatway explicitly, insted go shoud change it implicitly
// what we are doing over here is

type paymenter interface { // the convention of naming the interface is using "er" as suffix to the name like "tnventer, storer,loger "
	pay(amount float32) // we are making contract ->  any struct having "pay" method and reciving value of "(amount float32)-----" use in interface
}

type payment struct {
	gateway paymenter //this implementation is flexible implementation
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) { //this implementation is concrete implementation
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

func main() {
	razorpayPaymentGw := razorpay{}

	newPayment := payment{
		gateway: razorpayPaymentGw,
	}

	newPayment.makePayment(100)
}
