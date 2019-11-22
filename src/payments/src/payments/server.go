/*
	Payment API in Go (Version 2)
	Process Order with Go Channels and Mutex
*/

package main

import (
	"encoding/json"
	"fmt"
	//	"time"
	"bytes"
	"log"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
)

var mongodb_server = "10.0.1.18:27017"
var mongodb_database = "store"
var mongodb_collection = "payments"

var username = "admin"
var password = "admin"

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

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payments", getPaymentsHandler(formatter)).Methods("GET")
	//	mx.HandleFunc("/payment/{id}", paymentDeleteHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/payment", newPaymentHandler(formatter)).Methods("POST","OPTIONS")
	mx.HandleFunc("/payment/{id}", getPaymentsHandler(formatter)).Methods("GET")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 2.0 alive!"})
	}
}

// API Create New Payment Order
func newPaymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//	uuid, _ := uuid.NewV4()

		enableCors(&w)
		
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		} 
		//	decoder := json.NewDecoder(req.Body)
		var p payment
		//err := decoder.Decode(&p)
		// if err != nil {
		// 	panic(err)
		// }

		err := json.NewDecoder(req.Body).Decode(&p)

		if err != nil {
			panic(err)
		}

		//	json.Unmarshal([]byte(req.Body), &p)
		log.Println("card details : %+v", p)

		log.Println(p)
		log.Println(p.Id)
		log.Println(p.CardNumber)
		log.Println(p.Cvv)

		cardLen := len(string(p.CardNumber))
		cvvLen := len(string(p.Cvv))

		//	var myDate time.Time

		//	year, month, day := time.Now().Date()
		log.Println(cardLen)
		log.Println(cvvLen)

		if cardLen != 12 || cvvLen != 3 {
			formatter.JSON(w, http.StatusBadRequest, "Payment Failed... Invalid Card number")
		} else {

			pay := OrderResponse{
				UserEmail:        p.UserEmail,
				IsPaymentSuccess: true,
				CartItems:        p.CartItems,
			}

			//err := decoder.Decode(&pay)

			// if err != nil {
			// 	 panic(err)
			//  }

			log.Println(pay)

			session, err := mgo.Dial(mongodb_server)
			if err != nil {
				panic(err)
			}
			defer session.Close()
			admindb := session.DB("admin")
			err = admindb.Login(username, password)
			if err != nil {
				panic(err)
			}
			c := session.DB(mongodb_database).C(mongodb_collection)
			c.Insert(pay)
			if err != nil {
				log.Fatal(err)
			}

			bytesRepresentation, err := json.Marshal(pay)
			if err != nil {
				log.Fatalln(err)
			}

			resp, err := http.Post("https://xy0os460h9.execute-api.us-west-2.amazonaws.com/prod/addToCart", "application/json", bytes.NewBuffer(bytesRepresentation))

			if err != nil {
				log.Fatalln(err)
			}

			var result map[string]interface{}

			json.NewDecoder(resp.Body).Decode(&result)

			log.Println(result)
			//	log.Println("Data stored ", result["data"])

			fmt.Println("Payment: ", pay)
			formatter.JSON(w, http.StatusOK, pay)
		}

	}

}

// API Get Order Status
func getPaymentsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// if req.Method == "OPTIONS" {
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }

		enableCors(&w)
		params := mux.Vars(req)
		var uuid string = params["id"]
		fmt.Println("Payment ID: ", uuid)

		if uuid == "" {
			//	fmt.Println("Payments:", payments)

			// get payments from db

			var payments_array []payment
			for key, value := range payments_array {
				fmt.Println("Key:", key, "Value:", value)
				payments_array = append(payments_array, value)
			}
			formatter.JSON(w, http.StatusOK, payments_array)

		} else {

			// get payment from db

			var payment = payments[uuid]
			fmt.Println("Payment: ", payment)
			formatter.JSON(w, http.StatusOK, payment)
		}
	}
}

// func addCorsHeader(res http.ResponseWriter) {
//     headers := res.Header()
//     headers.Add("Access-Control-Allow-Origin", "*")
//     headers.Add("Vary", "Origin")
//     headers.Add("Vary", "Access-Control-Request-Method")
//     headers.Add("Vary", "Access-Control-Request-Headers")
//     headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
//     headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
// }

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "")
}
