package main

import (
	"fmt"
	"time"
)

// order struct
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time //nanosecond precision
// }

// // create construct
// func newOrder(id string, amount float32, status string) *order {
// 	//intial setup goes here...
// 	myorder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myorder
// }

// // reciever type and when you change the value of struct must use *
// func (o *order) changeStatus(status string) {
// 	o.status = status

// }

// func (o order) getamount() float32 {
// 	return o.amount

// }

/*func main() {

	//if you wnat to only ones using struct use like this
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// myorder := newOrder("1", 30.50, "recieved")
	// fmt.Println(myorder.amount)
	//if you dont set any filled ,default value is zeroed value
	// myorder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "recieved",
	// }

	// myorder.changeStatus("confirmed")
	// fmt.Println(myorder)

	// myorder.createdAt = time.Now()
	// fmt.Println(myorder.status)

	// myorder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delevierd",
	// 	createdAt: time.Now(),
	// }

	// myorder.status = "paid"
	// fmt.Println("order struct2", myorder2)
	// fmt.Println("order struct", myorder)
}*/

// struct Embedding
type customer struct {
	name  string
	phone string
}

// composion
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	customer
}

func main() {
	// newcuomer := customer{
	// 	name:  "Priti",
	// 	phone: "1234567890",
	// }
	neworder := order{
		id:     "1",
		amount: 30,
		status: "recieved",
		customer: customer{
			name:  "Priti",
			phone: "1234567890",
		},
	}
	neworder.customer.name = "Anikets"
	fmt.Println(neworder)

}
