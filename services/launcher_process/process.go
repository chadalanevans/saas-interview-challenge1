package launcher_process

import "log"
import "errors"
import . "github.com/chadalanevans/saas-interview-challenge1/workers/worker_sleep"
import . "github.com/chadalanevans/saas-interview-challenge1/types"

type Process struct {
	ConfiguredWorkers map[string]string
}

func (process *Process) Initialize() {
	
}

func (process Process) Launch(work Work) error {
	if(work.Type == "Sleep") {
	
		var sleepWorker = Sleep{}
		err := sleepWorker.Run(work)

		if(err != nil) {
			log.Printf("Worker completed with error: %s", err.Error())
			work.State = FAILED
		}

		work.State = COMPLETED
	} else {
		return errors.New("Unexpected Worker Type")
	}

	return nil
}

