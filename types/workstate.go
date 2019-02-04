package types

type WorkState int

const (
	QUEUED WorkState = iota
	PROCESSING
	COMPLETED
	FAILED
)

var states = [...] string {
	"Queued",
	"Processing",
	"Completed",
	"Failed",
}

var constStates = map[string]WorkState {
	"Queued"     :  QUEUED,
	"Processing" :  PROCESSING,
	"Completed"  :  COMPLETED,
	"Failed"     :  FAILED,
}

func (state WorkState) String() string {
	return states[state]
}

func (state *WorkState) FromString(stateStr string) {
	*state = constStates[stateStr];
}