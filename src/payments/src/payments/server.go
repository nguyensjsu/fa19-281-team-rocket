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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"

	"flag"

	"net/http"
	"os"
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
	mx.HandleFunc("/payment", newPaymentHandler(formatter)).Methods("POST", "OPTIONS")
	mx.HandleFunc("/payment/{id}", getPaymentsHandler(formatter)).Methods("GET")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 2.0 alive!"})
		log.Println("1")

		// emailPtr := flag.String("e", "megha3131@gmail.com", "The email address of the user subscribing to the topic")
		// topicPtr := flag.String("t", "arn:aws:sns:us-east-2:166329604693:payments", "The ARN of the topic to which the user subscribes")
		// flag.Parse()
		// email := *emailPtr
		// topicArn := *topicPtr

		// if email == "" || topicArn == "" {
		// 	fmt.Println("You must supply an email address and topic ARN")
		// 	fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
		// 	os.Exit(1)
		// }

		// // Initialize a session that the SDK will use to load
		// // credentials from the shared credentials file. (~/.aws/credentials).
		// sess, err := session.NewSessionWithOptions(session.Options{
		// 	// Specify profile to load for the session's config
		// 	//Profile: "profile_name",

		// 	// Provide SDK Config options, such as Region.
		// 	Config: aws.Config{
		// 		Region: aws.String("us-east-2"),
		// 	},

		// 	// Force enable Shared Config support
		// 	//SharedConfigState: session.SharedConfigEnable,
		// })

		// svc := sns.New(sess)

		// result, err := svc.Subscribe(&sns.SubscribeInput{
		// 	Endpoint:              emailPtr,
		// 	Protocol:              aws.String("email"),
		// 	ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		// 	TopicArn:              topicPtr,
		// })
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	os.Exit(1)
		// }

		// fmt.Println(*result.SubscriptionArn)



		//defer svc.Close()

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

			mongo_session, err := mgo.Dial(mongodb_server)
			if err != nil {
				panic(err)
			}
			defer mongo_session.Close()
			admindb := mongo_session.DB("admin")
			err = admindb.Login(username, password)
			if err != nil {
				panic(err)
			}
			c := mongo_session.DB(mongodb_database).C(mongodb_collection)
			c.Insert(pay)
			if err != nil {
				log.Fatal(err)
			}

			bytesRepresentation, err := json.Marshal(pay)
			if err != nil {
				log.Fatalln(err)
			}

			resp, err := http.Post("http://34.219.240.229:8080/newOrder", "application/json", bytes.NewBuffer(bytesRepresentation))

			if err != nil {
				log.Fatalln(err)
			}

			var result map[string]interface{}

			json.NewDecoder(resp.Body).Decode(&result)

			log.Println(result)
			//	log.Println("Data stored ", result["data"])

			fmt.Println("Payment: ", pay)
			formatter.JSON(w, http.StatusOK, pay)

			//Adding SNS
		
	//		defer svc.Close()


	msgPtr := flag.String("m", "Your payment is successful, order is on your way.. Enjoy :)", "The message to send to the subscribed users of the topic")
	topicPtr := flag.String("t", "arn:aws:sns:us-east-2:166329604693:payments", "The ARN of the topic to which the user subscribes")
	flag.Parse()
	message := *msgPtr
	topicArn := *topicPtr
	log.Println("2")
	if message == "" || topicArn == "" {
		fmt.Println("You must supply a message and topic ARN")
		fmt.Println("Usage: go run SnsPublish.go -m MESSAGE -t TOPIC-ARN")
		os.Exit(1)
	}
	log.Println("3")

	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	//	Profile: "profile_name",

	// 	Config: aws.Config{
	// 		Region: aws.String("us-east-2a"),
	// 	},
	// }))

	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		//Profile: "profile_name",

		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},

		// Force enable Shared Config support
		//SharedConfigState: session.SharedConfigEnable,
	})

	log.Println("4")

	svc := sns.New(sess)

	res, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: topicPtr,
		//TargetArn : aws.String("deepika.yannamani@sjsu.edu"),
	})

	log.Println("5")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*res.MessageId)
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
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
