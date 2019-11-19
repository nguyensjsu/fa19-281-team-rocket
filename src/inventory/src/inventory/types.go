package main

type Item struct {
	InventoryId	int `json:,string`
	Quantity	int `json:,string`
	Price		int `json:,string`
	Name		string
	Image		string
	Description	string
	Category	string
}