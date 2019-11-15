/*
	Payment API in Go (Version 2)
	Process Order with Go Channels and Mutex
*/

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"net/http"
)

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
	mx.HandleFunc("/payment/{id}", paymentDeleteHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/payment", newPaymentHandler(formatter)).Methods("POST")
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
		uuid, _ := uuid.NewV4()

		var p payment
        		_ = json.NewDecoder(req.Body).Decode(&p)

        cardLen := len(string(p.cardNumber))
        cvvLen := len(string(p.cvv))

        var myDate time.Time

        year, month, day := time.Now().Date()

        if cardLen !== 12 ||  cvvLen !== 3 || year > p.year || (year === p.year && month > p.month) {

            formatter.JSON(w, http.StatusOK, "Payment Failed... Invalid Card number")
        }
        else {
        // db update

       // formatter.JSON(w, http.StatusOK, "Payment Successful, Order Placed")

         var pay = payment{
         			Id:          uuid.String(),
         			PaymentStatus : "Payment Successful, Order Placed",
         		}


       	bytesRepresentation, err := json.Marshal(p)
       	if err != nil {
       		log.Fatalln(err)
       	}

       	resp, err := http.Post("https://grubhub/orders", "application/json", bytes.NewBuffer(bytesRepresentation))

       	if err != nil {
       		log.Fatalln(err)
       	}

       	var result map[string]interface{}

       	json.NewDecoder(resp.Body).Decode(&result)

       	log.Println(result)
       	log.Println("Data stored ", result["data"])


         fmt.Println("Payment: ", pay)
         formatter.JSON(w, http.StatusOK, pay)

     }
}

// API Get Order Status
func getPaymentsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		params := mux.Vars(req)
		var uuid string = params["id"]
		fmt.Println("Payment ID: ", uuid)

		if uuid == "" {
			fmt.Println("Payments:", payments)

			// get payments from db

			var payments_array []payment
			for key, value := range payments {
				fmt.Println("Key:", key, "Value:", value)
				payments_array = append(payments_array, value)
			}
			formatter.JSON(w, http.StatusOK, payments_array)

		} else {

		// get payment from db

			var payment = orders[uuid]
			fmt.Println("Payment: ", payment)
			formatter.JSON(w, http.StatusOK, payment)
		}
	}
}

