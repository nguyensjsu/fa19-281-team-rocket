/*
	Gumball API in Go (Version 1)
	Basic Version with no Backend Services
*/

package main

type payment struct {
	Id         string
	CardNumber string
	Cvv        string
	Expiry     string
	Zipcode    string
	Status     bool
	CartItems  []Item
	UserEmail  string
}

type Item struct {
	InventoryId string
	Quantity    string
	Item        string
	Price       int
	UserEmail   string
}

type OrderResponse struct {
	IsPaymentSuccess bool
	CartItems        []Item
	UserEmail        string
}

// var payment payment  = payment{

// }

var payments map[string]payment
