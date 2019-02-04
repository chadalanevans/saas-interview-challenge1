package worker_sleep

import "testing"
import "fmt"
import "time"

import . "github.com/chadalanevans/saas-interview-challenge1/types"

type TestPayLoad struct {
	data string
}

func (payLoad TestPayLoad) String() string {
	return payLoad.data
}

func TestEmptyPayLoad(t *testing.T) {

	var payLoad = ""

	var sleepWorker = Sleep{}
	var err = sleepWorker.parsePayLoad(payLoad)

	if(err == nil) {
		t.Error("ParsePayLoad did not return an error")
	}
	
	if(err.Error() != "Sleep worker config data not provided") {
		t.Error("ParsePayLoad did not return the expect error")
	}
}

func TestParsingPayLoadData(t *testing.T) {

	const expectedTime = 10
	const expectedMessage = "Hello"
	var jsonPayLoad = fmt.Sprintf("{\"SleepTimeInSeconds\": %d, \"Message\": \"%s\"}", expectedTime, expectedMessage)

	var sleepWorker = Sleep{}
	var err = sleepWorker.parsePayLoad(jsonPayLoad)

	if(err != nil) {
		t.Error(fmt.Errorf("Unexpected error parsing payLoad: %s", err.Error()))
	}

	if(sleepWorker.config.Message != expectedMessage) {
		t.Error(fmt.Errorf("Message not correct - expected: %s got: %s", expectedMessage,  sleepWorker.config.Message))
	}


}

func TestSleepWorkerPausesInSeconds(t *testing.T) {
	const expectedTimeInSeconds = 1
	const message = "Hello World"


	//Create a work object
	var work = Work{Id: 1, Type: "SleepWorker", State: QUEUED}
	work.Data = fmt.Sprintf("{\"SleepTimeInSeconds\": %d, \"Message\": \"%s\"}", expectedTimeInSeconds, message)

	//Run the worker and see if it ran for the specified time
	var sleepWorker = Sleep{}
	start := time.Now()
	err := sleepWorker.Run(work)
	elapsed := time.Since(start)

	if(err != nil) {
		t.Error(err.Error())
	}

	if(elapsed / time.Second != expectedTimeInSeconds) {
		t.Error(fmt.Errorf("Sleep worker ran for %d seconds, expected %d seconds", elapsed * time.Second, expectedTimeInSeconds))
	}
	

}