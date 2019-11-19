package main

import (
        "fmt"
        "log"
        "net/http"
		"encoding/json"
	    "github.com/codegangsta/negroni"
		"github.com/gorilla/mux" 
		"github.com/unrolled/render"
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
		// "strconv"
)

// MongoDB Config
var mongodb_server = "localhost"
var mongodb_database = "grubhub"
// var mongodb_collection_cart = "cart"
// var mongodb_collection_user = "user"
var mongodb_collection_userCart = "userCart"



// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
			IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

//API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/inventory", inventoryHandler(formatter)).Methods("GET")
	mx.HandleFunc("/addToCart", addItemsToCart(formatter)).Methods("POST")
	mx.HandleFunc("/cartItems/{emailId}", getCartItems(formatter)).Methods("GET")
	// mx.HandleFunc("/inventory/{id}", updateItemHandler(formatter)).Methods("PUT")
	// mx.HandleFunc("/inventory/{id}", deleteItemHandler(formatter)).Methods("DELETE")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

//API to add items to cart
func addItemsToCart(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		decoder := json.NewDecoder(req.Body)
		var cart Cart
		err := decoder.Decode(&cart)
		if err != nil {
			panic(err)
		}
		log.Println(cart)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
            panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection_userCart)
		c.Insert(cart)
		if err != nil {
            log.Fatal(err)
        }

	}
}

//API to get items from cart
func getCartItems(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		params := mux.Vars(req)
		var emailId string = params["emailId"]
		
		var result []bson.M
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
            panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection_userCart)
		
		//var result []Cart
		fmt.Println("emaildId",emailId)
		c.Find(bson.M{"useremail" : emailId}).All(&result)
		
		if err != nil {
            log.Fatal(err)
		}
		fmt.Println("Cart Items :", result )
        formatter.JSON(w, http.StatusOK, result)
        }
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}