package main

import (
        "fmt"
        "log"
        "net/http"
		"encoding/json"
        "github.com/codegangsta/negroni"
     //   "github.com/streadway/amqp"
        "github.com/gorilla/mux"
        "github.com/unrolled/render"
      //  "github.com/satori/go.uuid"
        "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

// MongoDB Config
var mongodb_server = "10.0.3.176:27017"
var mongodb_database = "inventory"
var mongodb_collection = "items"

var username = "admin"
var password = "Welcome_1"

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
        mx.HandleFunc("/inventory", inventoryHandler(formatter)).Methods("GET")
        mx.HandleFunc("/inventory", createItemHandler(formatter)).Methods("POST")
        mx.HandleFunc("/inventory/{id}", getItemHandler(formatter)).Methods("GET")
        mx.HandleFunc("/inventory/{id}", updateItemHandler(formatter)).Methods("PUT")
        mx.HandleFunc("/inventory/{id}", deleteItemHandler(formatter)).Methods("DELETE")
}

// Helper Functions
func failOnError(err error, msg string) {
        if err != nil {
                log.Fatalf("%s: %s", msg, err)
                panic(fmt.Sprintf("%s: %s", msg, err))
        }
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
				enableCors(&w)
                formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
        }
}

func inventoryHandler(formatter *render.Render) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			enableCors(&w)
        	fmt.Println("Returning full inventory..")
                session, err := mgo.Dial(mongodb_server)
        	if err != nil {
                panic(err)
			}
			defer session.Close()
			var results []bson.M
			admindb := session.DB("admin")
			err = admindb.Login(username, password)
			if err !=nil {
				panic(err)
			}
			c := session.DB(mongodb_database).C(mongodb_collection)
			err = c.Find(bson.M{}).All(&results)
			if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Inventory :", results )
                formatter.JSON(w, http.StatusOK, results)
        }
}

func createItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		decoder := json.NewDecoder(req.Body)
		var i Item
		err := decoder.Decode(&i)
		if err != nil {
			panic(err)
		}
		log.Println(i)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
            panic(err)
		}
		defer session.Close()
		admindb := session.DB("admin")
			err = admindb.Login(username, password)
			if err !=nil {
				panic(err)
			}
		c := session.DB(mongodb_database).C(mongodb_collection)
		c.Insert(i)
		if err != nil {
            log.Fatal(err)
        }

	}
}

func getItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		params := mux.Vars(req)
		var iid string = params["id"]
		log.Println("Get request for " + iid)
    	// string to int
		i, err := strconv.Atoi(iid)
		if err != nil {
			// handle error
			panic(err)
		}
		var result bson.M
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
            panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		admindb := session.DB("admin")
		err = admindb.Login(username, password)
		if err !=nil {
			panic(err)
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		c.Find(bson.M{"inventoryid" : i }).One(&result)
		if err != nil {
            log.Fatal(err)
		}
		fmt.Println("Inventory :", result )
                formatter.JSON(w, http.StatusOK, result)
        }
}

func deleteItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		params := mux.Vars(req)
		var iid string = params["id"]
		log.Println("Delete request for " + iid)
    	// string to int
		i, err := strconv.Atoi(iid)
		if err != nil {
			// handle error
			panic(err)
		}
		var result bson.M
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
            panic(err)
		}
		defer session.Close()
		admindb := session.DB("admin")
		err = admindb.Login(username, password)
		if err !=nil {
			panic(err)
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		c.Remove(bson.M{"inventoryid" : i })
		if err != nil {
            log.Fatal(err)
		}
		fmt.Println("Inventory :", result )
                formatter.JSON(w, http.StatusOK, result)
        }
}

func updateItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		enableCors(&w)
		params := mux.Vars(req)
		var iid string = params["id"]
		log.Println("Update request for " + iid)
		decoder := json.NewDecoder(req.Body)
		var item Item
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		log.Println(item)
		//change := bson.M{"test" : "one"}
    	// string to int
		i, err := strconv.Atoi(iid)
		if err != nil {
			// handle error
			panic(err)
		}
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
            panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		admindb := session.DB("admin")
		err = admindb.Login(username, password)
		if err !=nil {
			panic(err)
		}
		c := session.DB(mongodb_database).C(mongodb_collection)
		c.Update(bson.M{"inventoryid" : i }, item)
		if err != nil {
            log.Fatal(err)
		}
	}
}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}