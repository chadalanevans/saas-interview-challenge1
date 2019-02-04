package types

import "testing"
import "fmt"

func TestInitializationOfWork(t *testing.T) {

	const expectType = "MyWork"
	const expectState = "Queued"

	var work Work
	work.Type="MyWork"

	if ( expectType != work.Type) {
		t.Error(fmt.Errorf("Expected %s, got %s", expectType, work.State))
	}

	if (expectState != work.State.String()) {
		t.Error(fmt.Errorf("Expected %s, got %s", expectState, work.State.String()))
	}
}

func TestWorkToString(t *testing.T) {
	const workType = "Work"
	const workState = FAILED
	const workId = 1
	const workData = "SomeWorkData"

	var expectedStr = fmt.Sprintf("{\"Id\":%d,\"Type\":\"%s\",\"State\":%d,\"Data\":\"%s\"}", workId, workType, workState, workData);

	var work = Work{Id: workId, Type: workType, State: workState, Data: workData}

	var workStr = work.String()

	if(workStr != expectedStr) {
		t.Errorf("Expected %s, got %s", expectedStr, workStr)
	}
}

func TestStringToWork(t *testing.T) {
	const workType = "SomeWork"
	const workState = COMPLETED
	const workId = 1111
	const workData = "Data"

	var workStr = fmt.Sprintf("{\"Id\": %d, \"Type\":\"%s\", \"State\":%d,\"Data\":\"%s\"}", workId, workType, workState, workData);

	var work Work
	work.FromString(workStr)

	if(work.Type != workType) {
		t.Errorf("Type Expected %s, got %s", workType, work.Type)
	}

	if(work.State != workState) {
		t.Errorf("State Expected %s, got %s", workState, work.State)
	}

	if(work.Id != workId) {
		t.Errorf("Id Expected %d, got %d", workId, work.Id)
	}

	if(work.Data != workData) {
		t.Errorf("Data Expected %s, got %s", workData, work.Data)
	}


}