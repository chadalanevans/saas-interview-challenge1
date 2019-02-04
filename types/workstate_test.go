package types

import "testing"
import "fmt"

func TestConstToString(t *testing.T) {

	const expectFailed = "Failed"
	const expectQueued = "Queued"

	var state WorkState = QUEUED
	var stateStr = state.String()

	if ( "Queued" != stateStr) {
		t.Error(fmt.Errorf("Expected %s, got %s", expectQueued, stateStr))
	}

	state = FAILED
	stateStr = state.String()

	if ( expectFailed != stateStr) {
		t.Error(fmt.Errorf("Expected %s, got %s", expectFailed, stateStr))
	}
}

func TestStringToConst(t *testing.T) {

	const expectFailed = FAILED
	const expectQueued = QUEUED

	const queuedStr = "Queued"
	const failedStr = "Failed"
 
	var state WorkState
	state.FromString(queuedStr)
	
	if (expectQueued != state) {
		t.Error(fmt.Errorf("Expected %d, got %d", expectQueued, state))
	}

	var state2 WorkState
	state2.FromString(failedStr)

	if (expectFailed != state2) {
		t.Error(fmt.Errorf("Expected %d, got %d", expectFailed, state2))
	}
}