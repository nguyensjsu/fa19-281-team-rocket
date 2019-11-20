package main

type User struct {
	Id             	string 	
	Email 			string `json:"email"`
	Password		string	`json:"password"`
	Name 			string `json:"name"`
}