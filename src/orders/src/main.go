package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Person struct {
// 	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
// 	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
// }

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
	UserEmail        string             `json:"userEmail,omitempty" bson:"userEmail,omitempty"`
}

var client *mongo.Client

// func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	var person Person
// 	_ = json.NewDecoder(request.Body).Decode(&person)
// 	fmt.Println(person)
// 	collection := client.Database("mongo").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	result, _ := collection.InsertOne(ctx, person)
// 	json.NewEncoder(response).Encode(result)
// }

// func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	var people []Person
// 	collection := client.Database("mongo").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var person Person
// 		cursor.Decode(&person)
// 		people = append(people, person)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(people)
// }

// func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	params := mux.Vars(request)
// 	id, _ := primitive.ObjectIDFromHex(params["id"])
// 	var person Person
// 	collection := client.Database("mongo").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(person)
// }

//creates new order (POST: /newOrder)
func CreateNewOrder(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	response.Header().Set("content-type", "application/json")
	var orders Orders
	_ = json.NewDecoder(request.Body).Decode(&orders)
	fmt.Println("here")
	fmt.Println(orders)
	orders.OrderStatus = 1
	orders.OrderPlacedTime = time.Now().Unix()
	collection := client.Database("mongo").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, orders)
	json.NewEncoder(response).Encode(result)
}

//get order by Id (GET: /order/{id})
func GetOrderById(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
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

//get Order status(GET: /order/{id})
func GetOrderStatus(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
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
	response.Write([]byte(`{ "orderStatus": "` + strconv.Itoa(orders.OrderStatus) + `" }`))
}

//get all order(GET: /orders)
func GetAllOrders(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	response.Header().Set("content-type", "application/json")
	var orders []Orders
	collection := client.Database("mongo").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var o Orders
		cursor.Decode(&o)
		orders = append(orders, o)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(orders)
}

//get all orders of a particular user (GET: /allOrdersByEmail/{uEmail})
func GetAllOrdersByUserEmail(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	response.Header().Set("content-type", "application/json")
	//response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	email := params["uEmail"]
	var orders []Orders
	collection := client.Database("mongo").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var o Orders
		cursor.Decode(&o)
		if o.UserEmail == email {
			orders = append(orders, o)
		}
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(orders)
}

//get all orders by status (GET: /allOrdersByStatus/{status})
func GetAllOrdersByStatus(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	response.Header().Set("content-type", "application/json")
	//response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	os := params["status"]
	i, err := strconv.Atoi(os)
	if err != nil {
		// handle error
		fmt.Println(err)
		//os.Exit(2)
	}
	var orders []Orders
	collection := client.Database("mongo").Collection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var o Orders
		cursor.Decode(&o)
		if o.OrderStatus == i {
			orders = append(orders, o)
		}
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(orders)
}

//delete order by id (DELETE: /deleteOrder/{id})
func DeleteById(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	//var orders Orders
	collection := client.Database("mongo").Collection("orders")
	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	deleteResult, _ := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if deleteResult.DeletedCount == 0 {
		log.Fatal("Error on deleting one Hero")
	}
	response.Write([]byte(`{ "message": "order deleted" }`))
}

//update order status (PUT: /updateOrderStatus/{id}/{status})
func UpdateOrdeStatus(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	response.Header().Set("content-type", "application/json")
	//response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	status := params["status"]
	intStatus, err := strconv.Atoi(status)
	if err != nil {
		// handle error
		fmt.Println(err)
		//os.Exit(2)
	}
	collection := client.Database("mongo").Collection("orders")
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"orderStatus": intStatus}}
	result, err := collection.UpdateOne(context.Background(),
		filter,
		update)
	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)
		os.Exit(1)
	}
	fmt.Println(result)
	//response.Write([]byte(`{ "message": "` + + `" }`))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	//router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	//router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	//router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/newOrder", CreateNewOrder).Methods("POST")
	router.HandleFunc("/order/{id}", GetOrderById).Methods("GET")
	router.HandleFunc("/orderStatus/{id}", GetOrderStatus).Methods("GET")
	router.HandleFunc("/orders", GetAllOrders).Methods("GET")
	router.HandleFunc("/allOrdersByStatus/{status}", GetAllOrdersByStatus).Methods("GET")
	router.HandleFunc("/allOrdersByEmail/{uEmail}", GetAllOrdersByUserEmail).Methods("GET")
	router.HandleFunc("/deleteOrder/{id}", DeleteById).Methods("DELETE")
	router.HandleFunc("/updateOrderStatus/{id}/{status}", UpdateOrdeStatus).Methods("PUT")
	http.ListenAndServe(":12345", router)
}
