package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"log"
	"os"
	"net/http"
	//"github.com/streadway/amqp"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"flag"
)

// MongoDB Config
var mongodb_server = "mongodb"
var mongodb_database = "cmpe281"
var mongodb_collection = "gumballgo"

var username = "admin"
var password = "meg@697"

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
	mx.HandleFunc("/login", loginHandler(formatter)).Methods("POST", "OPTIONS")
	mx.HandleFunc("/signup", signupHandler(formatter)).Methods("POST", "OPTIONS")
	//	mx.HandleFunc("/user/{id}", gumballOrderStatusHandler(formatter)).Methods("GET")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})


		emailPtr := flag.String("e", "deepika.yannamani@sjsu.edu", "The email address of the user subscribing to the topic")
		topicPtr := flag.String("t", "arn:aws:sns:us-east-2:166329604693:payments", "The ARN of the topic to which the user subscribes")
		flag.Parse()
		email := *emailPtr
		topicArn := *topicPtr

		if email == "" || topicArn == "" {
			fmt.Println("You must supply an email address and topic ARN")
			fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
			os.Exit(1)
		}

		// Initialize a session that the SDK will use to load
		// credentials from the shared credentials file. (~/.aws/credentials).
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

		svc := sns.New(sess)

		result, err := svc.Subscribe(&sns.SubscribeInput{
			Endpoint:              emailPtr,
			Protocol:              aws.String("email"),
			ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
			TopicArn:              topicPtr,
		})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(*result.SubscriptionArn)
	}
}

// API Signup Handler
func signupHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		maxWait := time.Duration(5 * time.Second)
		mongo_session, err := mgo.DialWithTimeout("10.0.1.36:27017", maxWait)
		if err != nil {
			fmt.Println("Signup:", err)
			panic(err)
		}
		admindb := mongo_session.DB("admin")
		err = admindb.Login(username, password)
		if err != nil {
			panic(err)
		}
		defer mongo_session.Close()
		mongo_session.SetMode(mgo.Monotonic, true)
		c := mongo_session.DB(mongodb_database).C(mongodb_collection)
		uuid, err := uuid.NewV4()
		user.Id = uuid.String()
		err = c.Insert(user)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, "User Created Successfully: "+user.Id)

		// emailPtr := flag.String("e", user.Email, "The email address of the user subscribing to the topic")
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
	}
}

// API Login Handler
func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		maxWait := time.Duration(5 * time.Second)
		session, err := mgo.DialWithTimeout("10.0.1.36:27017", maxWait)
		if err != nil {
			fmt.Println("Signup:", err)
			panic(err)
		}
		admindb := session.DB("admin")
		err = admindb.Login(username, password)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		result := User{}
		err = c.Find(bson.M{"email": user.Email}).One(&result)

		if user.Email == result.Email && user.Password == result.Password {
			formatter.JSON(w, http.StatusOK, "Login successfully")
		} else {
			formatter.JSON(w, http.StatusBadRequest, "Login error")
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
