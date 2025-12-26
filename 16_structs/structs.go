package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string) order {
	return order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

// attaching a method to order struct
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	order1 := order{
		id:     "1",
		amount: 99.99,
		status: "received",
	}
	order1.createdAt = time.Now()

	fmt.Println("Amount:", order1.amount)
	fmt.Println("Order object 1:", order1)

	order2 := order{
		id:        "2",
		amount:    50.00,
		status:    "paid",
		createdAt: time.Now(),
	}

	fmt.Println("Order object 2:", order2)

	fmt.Println("Before changing status:", order2.status)
	// changing status using attached method
	order2.changeStatus("delivered")
	fmt.Println("After changing status:", order2.status)

	fmt.Println("Getting amount:", order2.getAmount())

	order3 := newOrder("3", 199.99, "paid")
	fmt.Println(order3)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)

	newCustomer := customer{
		name:  "parth",
		phone: "1234567890",
	}
	order4 := order{
		id:       "4",
		amount:   9.99,
		status:   "delivered",
		customer: newCustomer,
	}
	fmt.Println(order4)

	order5 := order{
		id:     "5",
		amount: 1.00,
		status: "ordered",
		customer: customer{
			name:  "parth",
			phone: "1234567890",
		},
	}
	order5.customer.name = "abcd"
	fmt.Println(order5)
}
