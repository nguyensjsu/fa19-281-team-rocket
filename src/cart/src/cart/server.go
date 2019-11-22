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
		// "github.com/aws/aws-sdk-go/aws"
    	// "github.com/aws/aws-sdk-go/aws/session"
    	// "github.com/aws/aws-sdk-go/service/sns"

    	// "flag"

    	//"os"
)

// MongoDB Config
var mongodb_server = "10.0.1.243"
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
	mx.HandleFunc("/addToCart", addItemsToCart(formatter)).Methods("POST","OPTIONS")
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
		mongo_session, err := mgo.Dial(mongodb_server)
        if err != nil {
            panic(err)
		}
		defer mongo_session.Close()
		mongo_session.SetMode(mgo.Monotonic, true)
		c := mongo_session.DB(mongodb_database).C(mongodb_collection_userCart)
		c.Insert(cart)
		
		if err != nil {
            log.Fatal(err)
		}

		formatter.JSON(w, http.StatusOK, "Added to cart successfully")
		


		//Adding SNS 
		// msgPtr := flag.String("m", "Your order is on its way", "The message to send to the subscribed users of the topic")
		// topicPtr := flag.String("t", "arn:aws:sns:us-west-2:253930511681:Payment", "The ARN of the topic to which the user subscribes")
		// flag.Parse()
		// message := *msgPtr
		// topicArn := *topicPtr
	
		// if message == "" || topicArn == "" {
		// 	fmt.Println("You must supply a message and topic ARN")
		// 	fmt.Println("Usage: go run SnsPublish.go -m MESSAGE -t TOPIC-ARN")
		// 	os.Exit(1)
		// }

		// sess := session.Must(session.NewSessionWithOptions(session.Options{
		// 	Profile: "profile_name",

		// 	Config: aws.Config{
		// 		Region: aws.String("us-west-2")
		// 	}

		// }))

		// sess, err := session.NewSessionWithOptions(session.Options{
		// 	// Specify profile to load for the session's config
		// 	Profile: "profile_name",
		
		// 	// Provide SDK Config options, such as Region.
		// 	Config: aws.Config{
		// 		Region: aws.String("us-west-2"),
		// 	},
		
		// 	// Force enable Shared Config support
		// 	//SharedConfigState: session.SharedConfigEnable,
		// })
	
		// svc := sns.New(sess)
	
		// result, err := svc.Publish(&sns.PublishInput{
		// 	Message:  aws.String(message),
		// 	TopicArn: topicPtr,
		// })
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	os.Exit(1)
		// }
	
		// fmt.Println(*result.MessageId)

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
	(*w).Header().Set("Access-Control-Allow-Headers", "*") 

}