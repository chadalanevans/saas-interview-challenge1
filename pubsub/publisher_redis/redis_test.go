package publisher_redis

import "testing"
import "time"
import "fmt"
import "log"
import "github.com/gomodule/redigo/redis"
import . "github.com/chadalanevans/saas-interview-challenge1/pubsub"
import . "github.com/chadalanevans/saas-interview-challenge1/types"

func TestPublishMessageSuccess(t *testing.T) {

	var publisher Publisher
	var redisPublisher = new(RedisPublisher)
	publisher = redisPublisher

	const eventType = "newEvent"
	const eventData = "SomeData"
	var eventCreated = time.Now()
	
	var event = Event{Type: eventType, Data: eventData, Created: eventCreated}
	var redisUrl = "localhost:6379"

	//Setup the publisher
	var err = publisher.Initialize(redisUrl)
	if(err != nil) {
		t.Error(fmt.Sprintf("Failed to initialize publisher: %s", err.Error()))
		return
	}

	//Setup a test subscriber
	log.Printf("Connecting to Redis Server at %s", redisUrl)

	var client, conErr = redis.Dial("tcp", redisUrl)
	if(conErr != nil) {
		t.Error("Failed to initialize client")
		return
	}

	var subscriber = redis.PubSubConn{Conn: client}
	subscriber.Subscribe(event.Type)

	go func() {
		//Try to recieve the message
		var done = false
		for done != true {
			log.Printf("Waiting for message")
			switch res := subscriber.Receive().(type) {
			case redis.Message:
				log.Printf("Got Message - %s",string(res.Data))
				done = true;
			case redis.Subscription:
				log.Printf("got subscription message for %s of %s", res.Channel, res.Kind)
			case error:
				log.Printf(res.Error())
				done = true
			}
		}
	}()

	//Pause for subscriber to get set
	time.Sleep(time.Duration(2) * time.Second)

	//Publish the test event
	err = publisher.Publish(event)
	if(err != nil) {
		t.Error(fmt.Sprintf("Failed to publish event: %s", err.Error()))
	}

	log.Printf("Sent Message")

	err = publisher.Finish()
	if(err != nil) {
		t.Error(fmt.Sprintf("Failed to finish publisher: %s", err.Error()))
	}
}