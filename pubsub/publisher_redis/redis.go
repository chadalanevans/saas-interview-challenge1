package publisher_redis

import "log"
import "errors"
import "github.com/gomodule/redigo/redis"
import . "github.com/chadalanevans/saas-interview-challenge1/types"

type RedisPublisher struct {
	Connection redis.Conn 
}

func (publisher *RedisPublisher) Initialize(url string) error {
	
	log.Printf("Connecting to Redis Server at %s", url)

	var connection, err = redis.Dial("tcp", url)

	if(err != nil) {
		log.Printf("Error trying to connect to redis: %s", err.Error())
		return err
	}

	publisher.Connection = connection

	log.Print("Connected to Redis Server")

	return nil
}

func (publisher RedisPublisher) Publish(event Event) error {

	if(publisher.Connection == nil) {
		log.Printf("Redis publisher not initialized")
		return errors.New("Redis publisher not initialized")
	}

	log.Printf("Sending event to redis %s", event.String())
	publisher.Connection.Do("PUBLISH", event.Type, event.String())

	return nil
}

func (publisher *RedisPublisher) Finish() error {

	if(publisher.Connection != nil) {
		publisher.Connection.Close()
		publisher.Connection = nil
	}

	return nil
}