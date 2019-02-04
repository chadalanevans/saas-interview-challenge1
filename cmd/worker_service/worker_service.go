package main

import (
    "log"
    "github.com/gomodule/redigo/redis"
    . "github.com/chadalanevans/saas-interview-challenge1/config"
    . "github.com/chadalanevans/saas-interview-challenge1/types"
    . "github.com/chadalanevans/saas-interview-challenge1/services"
    . "github.com/chadalanevans/saas-interview-challenge1/services/launcher_process"
)

var launcher Launcher

func main() {
    var redisUrl = GetRedisUrl()
    log.Printf("Connecting to Redis Server at - %s", redisUrl)

    var redisClient, err = redis.Dial("tcp", redisUrl)

    if(err != nil) {
        log.Printf("Failed to connect to server: %s", err.Error())
        return
    }
    defer redisClient.Close()

    var pubSubClient = redis.PubSubConn{Conn: redisClient}
    pubSubClient.Subscribe("Job")

    
    launcher = new(Process)
    
    for {
		switch res := pubSubClient.Receive().(type) {
        case redis.Message:
            log.Printf("Recieved new message: %s", res.Data)
            var work, parseErr = ParseMessage(res)
            if(parseErr != nil) {
                log.Printf("Trouble parsing message: %s, got error %s", res.Data, parseErr.Error())
                return 
            }
            ScheduleWork(work)
		case redis.Subscription:
			log.Printf("subscription message: %s: %s %d\n", res.Channel, res.Kind, res.Count)
		case error:
			log.Printf("Error receiving messages %s", res.Error())
			return
		}
	}
}

func ParseMessage(message redis.Message) (Work, error) {
    var event Event
    var err = event.FromString(string(message.Data))

    var work Work
    err = work.FromString(string(event.Data))
    return work, err
}

func ScheduleWork(work Work) {
    log.Printf("Launching work")
    var err = launcher.Launch(work)
    if(err != nil) {
        log.Printf("Error attempting to run work: %s", err.Error())
    }
    
}