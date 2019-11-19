package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// {
//     "InventoryID" : "2",
//     "Quantity" : "2",
//     "Item" : "Dum Biriyani",
//     "Price" :14,
//     "UserEmail" : "sam.mam@gmail.com"
// }

type Item struct {
	InventoryId string
	Quantity    string
	Item        string
	Price       int
	UserEmail   string
}

type Orders struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CartItems        []Item             `json:"cartItems,omitempty" bson:"cartItems,omitempty"`
	IsPaymentSuccess bool               `json:"isPaymentSuccess,omitempty" bson:"isPaymentSuccess,omitempty"`
	OrderStatus      int                `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	OrderPlacedTime  int64              `json:"orderPlacedTime,omitempty" bson:"orderPlacedTime,omitempty"`
	userEmail        string             `json:"userEmail,omitempty" bson:"userEmail,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	fmt.Println(person)
	collection := client.Database("mongo").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []Person
	collection := client.Database("mongo").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("mongo").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

//creates new order (POST: /newOrder)
func CreateNewOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var orders Orders
	_ = json.NewDecoder(request.Body).Decode(&orders)
	fmt.Println("here")
	fmt.Println(orders)
	orders.OrderStatus = 0
	orders.OrderPlacedTime = time.Now().Unix()
	collection := client.Database("mongo").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, orders)
	json.NewEncoder(response).Encode(result)
}

//get order by Id (GET: /order/{id})
func GetOrderById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var orders Orders
	collection := client.Database("mongo").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Orders{ID: id}).Decode(&orders)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(orders)
}

//get Order status

//update order status

//get all order

//get all orders of a particular user

//get all orders by status

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/newOrder", CreateNewOrder).Methods("POST")
	router.HandleFunc("/order/{id}", GetOrderById).Methods("GET")
	http.ListenAndServe(":12345", router)
}
