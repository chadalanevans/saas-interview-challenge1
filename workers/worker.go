package workers

import . "github.com/chadalanevans/saas-interview-challenge1/types"

type Worker interface {
	Run(work Work) error
}