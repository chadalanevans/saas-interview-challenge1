package worker_sleep

import "time"
import "log"
import "errors"
import "encoding/json"
import . "github.com/chadalanevans/saas-interview-challenge1/types"

const WORKER_TYPE string = "SleepWorker"

type Sleep struct {
	config Data
	info Work  
}

func (sleep *Sleep) Run(work Work) error {

	sleep.info = work
	var err = sleep.parsePayLoad(work.Data)
	if(err != nil) {
		log.Printf("Unable to start Sleep worker: %s", err.Error())
		return err
	}

	log.Printf("Sleep worker %d: started", sleep.info.Id)

	sleep.info.State = PROCESSING

	time.Sleep(time.Duration(sleep.config.SleepTimeInSeconds) * time.Second)

	log.Printf("Message from sleep worker: %s", sleep.config.Message)

	log.Printf("Sleep worker: %d exiting", sleep.info.Id)

	sleep.info.State = COMPLETED

	return nil
}


func (sleep *Sleep) parsePayLoad(payLoad string) error {

	if(payLoad == "") {
		return errors.New("Sleep worker config data not provided")
	}

	var err = json.Unmarshal([]byte(payLoad), &(sleep.config))

	if(err != nil) {
		return err
	}

	return nil
}

func (sleep *Sleep) Initializeinfo(id uint64) {
	sleep.info.Id = id
	sleep.info.Type = WORKER_TYPE
}
