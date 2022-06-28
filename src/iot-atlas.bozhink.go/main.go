package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"iot-atlas.bozhink.go/apihandlers"
)

const port int = 60135

var client *mongo.Client

func main() {
	fmt.Println("Starting the application...")

	if len(os.Args) < 2 {
		log.Fatal("Usage: atlas-iot <mongo-db-uri>")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	uri := os.Args[1]
	clientOptions := options.Client().ApplyURI(uri)

	var e error
	client, e = mongo.Connect(ctx, clientOptions)
	if e != nil {
		log.Fatal(e)
	}

	if cancel != nil {
		cancel()
	}

	insertEventHandler := apihandlers.GetInsertRecordEndpointHandler(client)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/events", insertEventHandler).Methods("POST")

	fmt.Printf("Start at port %d...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

	fmt.Println("After all...")
}
