package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"encoding/json"
)

// MongoDB Config
var mongodb_server = "mongodb"
var mongodb_database = "cmpe281"
var mongodb_collection = "gumballgo"

var username = "admin"
var password = "megz3189"
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
	}
}

// API Signup Handler
func signupHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		maxWait := time.Duration(5 * time.Second)
		session, err := mgo.DialWithTimeout("10.0.1.145:27017", maxWait)
        if err != nil {
				fmt.Println("Signup:", err )
                panic(err)
		}
		admindb := session.DB("admin")
		err = admindb.Login(username, password)
		if err !=nil {
			panic(err)
		}
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
		uuid, err := uuid.NewV4()
		user.Id = uuid.String()
        err = c.Insert(user)
        if err != nil {
                log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, "User Created Successfully: " + user.Id)
	}
}

// API Login Handler
func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		maxWait := time.Duration(5 * time.Second)
		session, err := mgo.DialWithTimeout("10.0.1.145:27017", maxWait)
        if err != nil {
				fmt.Println("Signup:", err )
                panic(err)
		}
		admindb := session.DB("admin")
		err = admindb.Login(username, password)
		if err !=nil {
			panic(err)
		}
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		
		result := User{}
		err = c.Find(bson.M{"email": user.Email}).One(&result)

		if (user.Email == result.Email && user.Password == result.Password) {
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

db.createUser( {
	user: "admin",
	pwd: "megz3189",
	roles: [{ role: "root", db: "admin" }]
});