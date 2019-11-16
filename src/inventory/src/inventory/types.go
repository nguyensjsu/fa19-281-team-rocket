package main

type Item struct {
	InventoryId	int `json:,string`
	Quantity	int `json:,string`
	Name		string
	Description	string
	Category	string
}