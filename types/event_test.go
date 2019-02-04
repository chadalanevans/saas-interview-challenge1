package types

import "testing"
import "fmt"
import "time"

func TestEventToString(t *testing.T) {
	
	const expectType = "TestType"
	const expectData = "SomeData"
	var expectCreated = time.Now()
	var expectEventStr = fmt.Sprintf("{\"Type\":\"%s\",\"Data\":\"%s\",\"Created\":\"%s\"}", expectType, expectData, expectCreated.Format(time.RFC3339Nano))  
	
	var event = Event{Type: expectType, Data: expectData, Created: expectCreated}

	var eventStr = event.String()

	if ( eventStr != expectEventStr) {
		t.Error(fmt.Errorf("Expected %s, got %s", expectEventStr, eventStr))
	}
}

func TestStringToEvent(t *testing.T) {
	const eventType = "SomeEvent"
	const eventData = "SomeData"
	var eventCreated = time.Now()

	var eventStr = fmt.Sprintf("{\"Type\": \"%s\", \"Data\":\"%s\", \"Created\":\"%s\"}", eventType, eventData, eventCreated.Format(time.RFC3339Nano));

	var event Event
	event.FromString(eventStr)

	if(event.Type != eventType) {
		t.Errorf("Type Expected %s, got %s", eventType, event.Type)
	}

	if(event.Data != eventData) {
		t.Errorf("Data Expected %s, got %s", eventData, event.Data)
	}

	if(event.Created.Format(time.RFC3339Nano) != eventCreated.Format(time.RFC3339Nano)) {
		t.Errorf("Created Expected %s, got %s", eventCreated, event.Created)
	}


}
