package main

import (
	"log"
	"time"
	"io/ioutil"
	"net/http"
	"math/rand"
	"github.com/gorilla/mux"
	
	. "github.com/chadalanevans/saas-interview-challenge1/config"
	. "github.com/chadalanevans/saas-interview-challenge1/types"
	. "github.com/chadalanevans/saas-interview-challenge1/pubsub"
	. "github.com/chadalanevans/saas-interview-challenge1/pubsub/publisher_redis"
)

var publisher Publisher

func main() {

	// Initialize connection to the publisher service
	publisher = new(RedisPublisher)
	var err = publisher.Initialize(GetRedisUrl())
	if(err != nil) {
		log.Printf("Failed to initialize publisher: %s", err.Error())
	}
	defer publisher.Finish()

	// Start http listener for job api requests
	router := mux.NewRouter()
	router.HandleFunc("/job", CreateJob).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func CreateJob(response http.ResponseWriter, request *http.Request) {

	// Ready the body of the create job request
	body, err := ioutil.ReadAll(request.Body)
	if(err != nil) {
		log.Printf("Error reading http body: %s", err.Error())
		return
	}

	// Parse the body and create new work item
	var work Work
	work.FromString(string(body))
	work.Id = rand.Uint64()
	log.Printf("Got Work: %s", work.String())

	// Publish new event with the new job
	var newJobEvent Event
	newJobEvent.Type = "Job"
	newJobEvent.Data = work.String()
	newJobEvent.Created = time.Now()
	publisher.Publish(newJobEvent)
}