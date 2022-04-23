package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/karanbirsingh7/backend-challenge/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionURI string
var DB *mongo.Client

func init() {
	connectionURI = os.Getenv("MONGO_URI")
	if connectionURI == "" {
		log.Fatalf("DB connection not provided %q. Exiting.", "MONGO_URI")
	}
}

func ConnectDB() *mongo.Client {
	// DB connection and verification
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatalln("Cannot connect to mongo database", err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalln("Ping to DB failed with error:", err)
	}
	log.Println("Connection to DB success")
	return client
}

func main() {
	DB = ConnectDB()

	db.PopulateDBWithSampleData(DB)

	// SERVER
	// Endpoint 1: GET /talents -> List all
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/talents")
	})
	http.HandleFunc("/talents", ListTalentHandler)

	// Endpoint 1: POST /talent -> create a new entry
	http.HandleFunc("/talent", CreateTalentHandler)
	log.Println("Server starting on https://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func ListTalentHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db.FetchAllTalents(DB)
	if err != nil {
		http.Error(w, "Unable to fetch records from DB", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(data)
}

func CreateTalentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "new talent item created")
}
