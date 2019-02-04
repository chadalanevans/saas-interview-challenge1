package launcher_process

import "testing"
import . "github.com/chadalanevans/saas-interview-challenge1/services"
import . "github.com/chadalanevans/saas-interview-challenge1/types"

func TestFailedLaunchUnknownType(t *testing.T) {
	var launcher Launcher
	var work Work

	launcher = new(Process)
	var err = launcher.Launch(work);
	if ( err == nil && err.Error() != "Unexpected Worker Type") {
		t.Error("Expected an error with the type of worker")
	}
}